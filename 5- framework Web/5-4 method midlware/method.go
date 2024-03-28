package main

import (
	"net/http"
)

type app struct {
	routes map[string]map[string]http.HandlerFunc // 1 Add a map of routes and
}

func New() *app {
	return &app{
		routes: make(map[string]map[string]http.HandlerFunc),
	}
}

func (a *app) Handle(path, method string, handler http.HandlerFunc) {
	_, ok := a.routes[path]
	if !ok {
		a.routes[path] = make(map[string]http.HandlerFunc)
	}
	a.routes[path][method] = handler
}

func (a *app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if route, ok := a.routes[r.URL.Path]; ok {
		if h, ok := route[r.Method]; ok {
			h.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("405 Method Not Allowed"))
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
	}
}

func main() {
	app := New()

	app.Handle("/hello", "GET", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("GET, Zone01!"))
	})

	app.Handle("/hello", "POST", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("POST, Zone01!"))
	})

	http.ListenAndServe(":8080", app)
}
