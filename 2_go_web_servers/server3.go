package main

import (
  "fmt"
  "net/http"
)

// Handler function needs to be converted to a HandlerFunc type
// Implicit conversion to HandlerFunc
func handler() http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "[HandlerFunc Implicit] Bonjour GoMN Gophers!")
  }
}

func main() {

  http.Handle("/", handler()) // handler needs to be invoked

  http.ListenAndServe(":8080", nil)
}