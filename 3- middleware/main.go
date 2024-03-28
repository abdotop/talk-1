package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Zone01!")
}

func midleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("405 Method not allowed"))
		}
	}
}

func main() {
	next := http.HandlerFunc(helloHandler)

	h := midleware(next)

	http.HandleFunc("/hello", h)

	http.ListenAndServe(":8080", nil)
}
