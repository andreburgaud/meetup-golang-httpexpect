package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestHello(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, _ *http.Request) {
			fmt.Fprintf(w, "Hello Gophers!")
		}))
	defer server.Close()

	// Create a new expect object
	e := httpexpect.New(t, server.URL)

	e.GET("/").Expect().Status(http.StatusOK).
		Text().Equal("Hello Gophers!")
}
