package main_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	main "github.com/cpliakas/aws-sam-golang-example"
)

// tests implements the testing table pattern.
var tests = map[string]struct {
	url     string
	handler http.HandlerFunc
	method  string
	code    int
	body    interface{}
}{
	"RootHandler": {
		url:     "/",
		handler: main.RootHandler,
		method:  http.MethodGet,
		code:    http.StatusOK,
		body:    main.WelcomeMessage,
	},
	"HelloHandler": {
		url:     "/hello",
		handler: main.HelloHandler,
		method:  http.MethodGet,
		code:    http.StatusOK,
		body:    main.HelloMessage,
	},
	"GoodbyeHandler": {
		url:     "/goodbye",
		handler: main.GoodbyeHandler,
		method:  http.MethodGet,
		code:    http.StatusOK,
		body:    main.GoodbyeMessage,
	},
}

// TestHandlers tests the handlers in isolation.
func TestHandlers(t *testing.T) {
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

// TestRoutes starts a server and tests the routes by making calls to it.
func TestRoutes(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping API tests in short mode.")
	}

	main.RegisterRoutes()
	server := httptest.NewServer(nil)
	defer server.Close()

	for _, test := range tests {
		t.Logf("Running test case for route %s", test.url)

		resp, err := http.Get(server.URL + test.url)
		if err != nil {
			t.Fatal(err)
		}

		if code := resp.StatusCode; code != test.code {
			t.Errorf("expected status code %v, got %v", test.code, code)
		}

		if header := resp.Header.Get("Content-Type"); header != main.ContentType {
			t.Errorf("expected Content-Type header %s, got %s", main.ContentType, header)
		}

		ex, err := json.Marshal(test.body)
		if err != nil {
			t.Fatalf("error marshalling expected body: %s", err)
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("error reading body: %s", err)
		}

		got := bytes.TrimSpace(body)
		if !bytes.Equal(got, ex) {
			t.Errorf("expected message '%s', got '%s'", ex, got)
		}
	}
}
