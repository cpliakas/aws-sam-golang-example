package main

// SO Answer:
// https://stackoverflow.com/a/15685432/1913888
// Test:
// $ curl -X POST -d "{\"foo\": \"that\"}" http://localhost:8082/test

// still getting "Missing Auth Token"
// $ curl -X POST -d "{\"foo\": \"that\"}" http://localhost:3000/

import (
	"encoding/json"
	"log"
	"net/http"
)

type test_struct struct {
	Foo string
}

func test(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t test_struct
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	log.Println(t.Foo)
}

func main() {
	http.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
