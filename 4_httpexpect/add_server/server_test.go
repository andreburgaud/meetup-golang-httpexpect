package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

var tt = []struct {
	name     string
	a        string
	b        string
	expected string
}{
	{"Add2and3", "2", "3", "5"},
	{"Add0and3", "0", "3", "3"},
	{"Add-100and2", "-100", "2", "-98"},
}

func TestAddServer(t *testing.T) {
	s := httptest.NewServer(handler())
	defer s.Close()

	// Create a new expect object
	e := httpexpect.New(t, s.URL)

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			e.GET("/add").
				WithQuery("a", tc.a).
				WithQuery("b", tc.b).
				Expect().
				Status(http.StatusOK).
				Text().
				Equal(tc.expected)
		})
	}
}
