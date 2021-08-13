package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/apex/gateway"
)

// ContentType contains the Content-Type header sent on all responses.
const ContentType = "application/json; charset=utf8"

// MessageResponse models a simple message responses.
type MessageResponse struct {
	Message string `json:"message"`
}

// WelcomeMessageResponse is the response returned by the / endpoint.
var WelcomeMessageResponse = MessageResponse{"Welcome to the example API!"}

type test_struct struct {
	Foo string
}

// RootHandler is a http.HandlerFunc for the / endpoint.
func RootHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t test_struct
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	log.Println(t.Foo)

	json.NewEncoder(w).Encode(WelcomeMessageResponse)
}

// RegisterRoutes registers the API's routes.
func RegisterRoutes() {
	http.Handle("/", h(RootHandler))
}

// h wraps a http.HandlerFunc and adds common headers.
func h(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("event: %v", r)
		w.Header().Set("Content-Type", ContentType)
		next.ServeHTTP(w, r)
	})
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	RegisterRoutes()

	log.Fatal(gateway.ListenAndServe(":3000", nil))
}
