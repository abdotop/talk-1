package main

import (
	"fmt"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintln(w, "GET Request")
	case http.MethodPost:
		fmt.Fprintln(w, "POST Request")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/hello", handleRequest)
	http.ListenAndServe(":8080", nil)
}
