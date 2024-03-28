package main

import (
	"net/http"
)

type app struct {
	routes map[string]http.HandlerFunc // 1 Add a map of routes
}

func New() *app {
	return &app{
		routes: make(map[string]http.HandlerFunc),
	}
}

func (a *app) Handle(path string, handler http.HandlerFunc) {
	a.routes[path] = handler
}

func (a *app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := a.routes[r.URL.Path]; ok {
		h(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
	}
}

func main() {
	app := New()

	app.Handle("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Zone01!"))
	})

	http.ListenAndServe(":8080", app)
}
