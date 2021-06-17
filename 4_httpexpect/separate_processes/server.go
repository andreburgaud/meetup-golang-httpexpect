// server.go

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/1", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "Hello Gophers [1]!")
	})

	http.HandleFunc("/2", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "Hello Gophers [2]!")
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
