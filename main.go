package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/apex/gateway"
)

// ContentType is the Content-Type header set in responses.
const ContentType = "application/json; charset=utf8"

// Message contains a simple message response.
type Message struct {
	Message string `json:"message"`
}

// Messages used by http.HandlerFunc functions.
var (
	WelcomeMessage = Message{"Welcome to the example API!"}
	HelloMessage   = Message{"Hello, world!"}
	GoodbyeMessage = Message{"Goodbye, world!"}
)

// RootHandler is a http.HandlerFunc for the / path.
func RootHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(WelcomeMessage)
}

// HelloHandler is a http.HandlerFunc for the /hello path.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(HelloMessage)
}

// GoodbyeHandler is a http.HandlerFunc for the /goodbye path.
func GoodbyeHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(GoodbyeMessage)
}

// RegisterRoutes registers the API's routes.
func RegisterRoutes() {
	http.Handle("/", h(RootHandler))
	http.Handle("/hello", h(HelloHandler))
	http.Handle("/goodbye", h(GoodbyeHandler))
}

// h wraps a http.HandlerFunc and adds common headers.
func h(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", ContentType)
		next.ServeHTTP(w, r)
	})
}

func main() {
	RegisterRoutes()
	log.Fatal(gateway.ListenAndServe(":3000", nil))
}
