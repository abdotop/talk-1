package main

import (
	"net/http"
)

type app struct {
}

func New() *app {
	return new(app)
}

func (a *app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Zone01!"))
}

func main() {
	app := New()

	http.ListenAndServe(":8080", app)
}
