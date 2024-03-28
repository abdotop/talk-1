package main

import "net/http"

type app struct {
}

func New() *app {
	return new(app)
}

func main() {
	app := New()

	http.ListenAndServe(":8080", app)
}
