package main

// SO Answer:
// https://stackoverflow.com/a/15685432/1913888
// Test:
// $ curl -X POST -d "{\"foo\": \"that\"}" http://localhost:8082/test

import (
	"encoding/json"
	"log"
	"net/http"
)

type Payload struct {
	Foo string
}

func test(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t Payload
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	log.Println(t.Foo)
}

func main() {
	http.HandleFunc("/test", test)
	http.HandleFunc("/", test)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
