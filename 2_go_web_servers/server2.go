package main

import (
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "[HandlerFunc] Bonjour GoMN Gophers!")
}

func main() {

  // Handler function needs to be converted to a HandlerFunc type
  h := http.HandlerFunc(handler)

  // And add it to the DefaultServeMux
  http.Handle("/", h)

  http.ListenAndServe(":8080", nil)
}