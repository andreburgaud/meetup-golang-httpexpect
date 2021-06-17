package main

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestHello1(t *testing.T) {
	e := httpexpect.New(t, "http://localhost:8000")
	e.GET("/1").Expect().Status(http.StatusOK).Text().Equal("Hello Gophers [1]!")
}

func TestHello2(t *testing.T) {
	e := httpexpect.New(t, "http://localhost:8000")
	e.GET("/2").Expect().Status(http.StatusOK).Text().Equal("Hello Gophers [2]!")
}
