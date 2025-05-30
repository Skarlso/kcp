/*
Copyright 2022 The KCP Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package webhook

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"sync"
	"testing"

	admissionv1 "k8s.io/api/admission/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type AdmissionWebhookServer struct {
	ResponseFn   func(obj runtime.Object, review *admissionv1.AdmissionReview) (*admissionv1.AdmissionResponse, error)
	ObjectGVK    schema.GroupVersionKind
	Deserializer runtime.Decoder

	t *testing.T

	host, port string
	lock       sync.Mutex
	calls      int
}

func (s *AdmissionWebhookServer) StartTLS(t *testing.T, certFile, keyFile, host, port string) {
	t.Helper()

	s.t = t
	// The host passed to StartTLS is the Host of the rest.Config, which
	// can be just host, host:port or a full URL.
	u, err := url.Parse(host)
	if err != nil {
		t.Fatalf("error parsing host %q: %v", host, err)
	}
	host, _, err = net.SplitHostPort(u.Host)
	if err != nil {
		t.Fatalf("error splitting host %q: %v", u.Host, err)
	}
	s.host = host
	s.port = port

	serv := &http.Server{Addr: net.JoinHostPort(s.host, s.port), Handler: s}
	t.Cleanup(func() {
		t.Log("Shutting down the HTTP server")
		err := serv.Shutdown(context.TODO())
		if err != nil {
			t.Logf("unable to shutdown server gracefully err: %v", err)
		}
	})

	go func() {
		err := serv.ListenAndServeTLS(certFile, keyFile)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			t.Logf("unable to shutdown server gracefully err: %v", err)
		}
	}()
}

func (s *AdmissionWebhookServer) GetURL() string {
	u := &url.URL{
		Scheme: "https",
		Host:   net.JoinHostPort(s.host, s.port),
		Path:   "/hello",
	}
	return u.String()
}

func (s *AdmissionWebhookServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	// Make sure that this is a request for the object that was set.
	s.t.Log("made it webhook")
	if req.Body == nil {
		msg := "Expected request body to be non-empty"
		s.t.Logf("%v", msg)
		http.Error(resp, msg, http.StatusBadRequest)
		return
	}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		msg := fmt.Sprintf("Request could not be decoded: %v", err)
		s.t.Logf("%v", msg)
		http.Error(resp, msg, http.StatusBadRequest)
		return
	}

	// verify the content type is accurate
	contentType := req.Header.Get("Content-Type")
	if contentType != "application/json" {
		msg := fmt.Sprintf("contentType=%s, expect application/json", contentType)
		s.t.Logf("%v", msg)
		http.Error(resp, msg, http.StatusBadRequest)
		return
	}

	obj, gvk, err := s.Deserializer.Decode(data, nil, nil)
	if err != nil {
		msg := fmt.Sprintf("Unable to decode object: %v", err)
		s.t.Logf("%v", msg)
		http.Error(resp, msg, http.StatusBadRequest)
		return
	}

	if *gvk != admissionv1.SchemeGroupVersion.WithKind("AdmissionReview") {
		msg := fmt.Sprintf("Expected AdmissionReview but got: %T", obj)
		s.t.Logf("%v", msg)
		http.Error(resp, msg, http.StatusBadRequest)
		return
	}
	requestedAdmissionReview, ok := obj.(*admissionv1.AdmissionReview)
	if !ok {
		// return an error
		msg := fmt.Sprintf("Expected AdmissionReview but got: %T", obj)
		s.t.Logf("%v", msg)
		http.Error(resp, msg, http.StatusBadRequest)
		return
	}
	obj, objGVK, err := s.Deserializer.Decode(requestedAdmissionReview.Request.Object.Raw, nil, nil)
	if err != nil {
		msg := fmt.Sprintf("Unable to decode admissions requested object: %v", err)
		s.t.Logf("%v", msg)
		http.Error(resp, msg, http.StatusBadRequest)
		return
	}

	if s.ObjectGVK != *objGVK {
		// return an error
		msg := fmt.Sprintf("Expected ObjectGVK: %v but got: %T", s.ObjectGVK, obj)
		s.t.Logf("%v", msg)
		http.Error(resp, msg, http.StatusBadRequest)
		return
	}

	responseAdmissionReview := &admissionv1.AdmissionReview{
		TypeMeta: requestedAdmissionReview.TypeMeta,
	}
	r, err := s.ResponseFn(obj, requestedAdmissionReview)
	if err != nil {
		s.t.Logf("%v", err)
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	responseAdmissionReview.Response = r
	responseAdmissionReview.Response.UID = requestedAdmissionReview.Request.UID

	respBytes, err := json.Marshal(responseAdmissionReview)
	if err != nil {
		s.t.Logf("%v", err)
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}

	s.lock.Lock()
	defer s.lock.Unlock()
	s.calls++

	resp.Header().Set("Content-Type", "application/json")
	if _, err := resp.Write(respBytes); err != nil {
		s.t.Logf("%v", err)
	}
}

func (s *AdmissionWebhookServer) Calls() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.calls
}
