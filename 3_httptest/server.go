package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// addHandler process the querystring
// /add?a=<value_a>&b=<value_b>
func addHandler(w http.ResponseWriter, r *http.Request) {
	qsa := r.FormValue("a")
	if qsa == "" {
		http.Error(w, "missing value a", http.StatusBadRequest)
		return
	}

	qsb := r.FormValue("b")
	if qsb == "" {
		http.Error(w, "missing value b", http.StatusBadRequest)
		return
	}

	a, err := strconv.Atoi(qsa)
	if err != nil {
		http.Error(w, "a is not a number: "+qsa, http.StatusBadRequest)
		return
	}

	b, err := strconv.Atoi(qsb)
	if err != nil {
		http.Error(w, "b is not a number: "+qsb, http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%d", a+b)
}

// handler wraps the router for testing
func handler() http.Handler {
	http.HandleFunc("/add", addHandler)
	return http.DefaultServeMux
}

func main() {
	http.ListenAndServe(":8080", handler())
}
