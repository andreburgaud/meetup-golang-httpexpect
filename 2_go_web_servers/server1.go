package main

import (
	"fmt"
	"net/http"
)

// Handler should satisfy the http.Handler interface:
// type Handler interface {
//    ServeHTTP(ResponseWriter, *Request)
//}
// https://golang.org/src/net/http/server.go?s=78220:78291#L62

type handler struct{}

func (handler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "[ServeHTTP] Bonjour GoMN Gophers!")
}

func main() {
	// var h handler
	// http.ListenAndServe(":8080", h)
	http.ListenAndServe(":8080", handler{})
}
