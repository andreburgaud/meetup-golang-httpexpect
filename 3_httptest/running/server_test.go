package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
)

func TestAddRunningServer(t *testing.T) {
	tt := []struct {
		name     string
		a        string
		b        string
		expected int
	}{
		{"Add2and3", "2", "3", 5},
		{"Add0and3", "0", "3", 3},
		{"Add-100and2", "-100", "2", -98},
	}

	const url = "http://localhost:8080"

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			res, err := http.Get(fmt.Sprintf("%s/add?a=%s&b=%s", url, tc.a, tc.b))
			if err != nil {
				t.Fatalf("could not send GET request: %v", err)
			}
			defer res.Body.Close()

			if res.StatusCode != http.StatusOK {
				t.Errorf("expected status OK 200; got %v", res.Status)
			}

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read response: %v", err)
			}

			add, err := strconv.Atoi(string(bytes.TrimSpace(body)))
			if err != nil {
				t.Fatalf("expected an integer; got %q", add)
			}

			if add != tc.expected {
				t.Fatalf("expected double to be 5; got %v", add)
			}
		})
	}
}
