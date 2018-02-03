package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	main "github.com/cpliakas/aws-sam-golang-example"
)

func TestHandlers(t *testing.T) {

	// Implements the testing table pattern.
	tests := map[string]struct {
		handler http.HandlerFunc
		method  string
		code    int
		body    interface{}
	}{
		"RootHandler": {
			handler: main.RootHandler,
			method:  http.MethodGet,
			code:    http.StatusOK,
			body:    main.WelcomeMessage,
		},
		"HelloHandler": {
			handler: main.HelloHandler,
			method:  http.MethodGet,
			code:    http.StatusOK,
			body:    main.HelloMessage,
		},
		"GoodbyeHandler": {
			handler: main.GoodbyeHandler,
			method:  http.MethodGet,
			code:    http.StatusOK,
			body:    main.GoodbyeMessage,
		},
	}

	for name, test := range tests {
		t.Logf("Running test case for %s", name)

		req, err := http.NewRequest(test.method, "/", nil)
		if err != nil {
			t.Fatalf("error creating request: %s", err)
		}

		rr := httptest.NewRecorder()
		h := http.HandlerFunc(test.handler)
		h.ServeHTTP(rr, req)

		if code := rr.Code; code != test.code {
			t.Errorf("expected status code %v, got %v", test.code, code)
		}

		ex, err := json.Marshal(test.body)
		if err != nil {
			t.Fatalf("error marshalling expected body: %s", err)
		}

		got := bytes.TrimSpace(rr.Body.Bytes())
		if !bytes.Equal(got, ex) {
			t.Errorf("expected message '%s', got '%s'", ex, got)
		}
	}
}
