package main

import (
	"fmt"
	"net/http"
)

// handler function that returns a string
// expected signature (http.ResponseWriter, *http.Request)
func handler1(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "[handleFunc 1] Bonjour GoMN Gophers!")
}

func handler2(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "[handleFunc 2] Bonjour GoMN Gophers!")
}

func main() {
	// http.HandleFunc registers handlers as handler functions
	// for patterns"/1" and "1/2"
	// in the DefaultServeMux
	// https://golang.org/src/net/http/server.go?s=78220:78291#L2504
	http.HandleFunc("/1", handler1)
	http.HandleFunc("/2", handler2)
	http.ListenAndServe(":8080", nil)
}
