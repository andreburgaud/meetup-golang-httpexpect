package main

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

const (
	url = "http://localhost:9999"
)

// TestHelloGophers tests a GET with JSON Response
func TestHelloGophers(t *testing.T) {
	e := httpexpect.New(t, url)

	e.GET("/hi").
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		ContainsKey("message").
		ValueEqual("message", "Hello Gophers!")

}

// TestGetItemID tests parallel GET with JSON Response
func TestGetItemID(t *testing.T) {
	var tt = []struct {
		name     string
		itemID   string
		expected string
	}{
		{"Request_ID_1", "1", "1"},
		{"Request_ID_12", "12", "12"},
		{"Request_ID_123", "123", "123"},
		{"Request_ID_1234", "1234", "1234"},
		{"Request_ID_12345", "12345", "12345"},
		{"Request_ID_123456", "123456", "123456"},
		{"Request_ID_1234567", "1234567", "1234567"},
		{"Request_ID_12345678", "12345678", "12345678"},
		{"Request_ID_123456789", "123456789", "123456789"},
	}

	e := httpexpect.New(t, url)

	for _, tc := range tt {
		tc := tc // Capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel() // Execute tests in parallel
			e.GET("/ids/{item_id}", tc.itemID).
				Expect().
				Status(http.StatusOK).
				JSON().
				Object().
				ContainsKey("item_id").
				ValueEqual("item_id", tc.expected)
		})
	}
}

func TestPostGetItem(t *testing.T) {
	e := httpexpect.New(t, url)

	var tt = []struct {
		name   string
		author string
		price  float64
	}{
		{"Golang", "Rob Pike", 200},
		{"Python", "Guido Van Rossum", 10},
		{"Haskell", "Simon Peyton Jones", 600000},
		{"Perl", "Larry Wall", -30},
		{"Java", "James Gosling", 50000},
		{"C++", "Bjarne Stroustrup", 50001},
		{"C", "Dennis Ritchie", 50},
		{"JavaScript", "Brendan Eich", -60},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			item := map[string]interface{}{
				"name":   tc.name,
				"author": tc.author,
				"price":  tc.price,
			}
			e.POST("/items").
				WithJSON(item).
				Expect().
				Status(http.StatusOK).
				JSON().
				Object().
				ContainsKey("name").
				ValueEqual("name", tc.name)

			e.GET("/items/{name}", tc.name).
				Expect().
				Status(http.StatusOK).
				JSON().
				Object().
				ContainsKey("name").
				ValueEqual("name", tc.name)
		})
	}

	e.GET("/items").
		Expect().
		Status(http.StatusOK).
		JSON().
		Array().
		Length().
		Equal(8)

}

func TestNotFoundItem(t *testing.T) {
	e := httpexpect.New(t, url)

	e.GET("/items/forth").
		Expect().
		Status(http.StatusNotFound).
		JSON().
		Object().
		ContainsKey("detail").
		ValueEqual("detail", "Item not found")

}
