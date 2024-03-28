package main

import (
	"net/http"
)

type Contex struct {
	R *http.Request
	W http.ResponseWriter
}

type HandlerFunc func(*Contex)

type app struct {
	routes map[string]map[string]HandlerFunc // 1 Add a map of routes and
}

func New() *app {
	return &app{
		routes: make(map[string]map[string]HandlerFunc),
	}
}

func (a *app) handle(path, method string, handler HandlerFunc) {
	_, ok := a.routes[path]
	if !ok {
		a.routes[path] = make(map[string]HandlerFunc)
	}
	a.routes[path][method] = handler
}

func (a *app) Get(path string, h HandlerFunc) { // get method
	a.handle(path, "GET", h)
}

func (a *app) Post(path string, h HandlerFunc) { // post method
	a.handle(path, "POST", h)
}

func (a *app) Put(path string, h HandlerFunc) { // put method
	a.handle(path, "PUT", h)
}

func (a *app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if route, ok := a.routes[r.URL.Path]; ok {
		if h, ok := route[r.Method]; ok {
			h(&Contex{r, w})
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("405 Method Not Allowed"))
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
	}
}

func (a *app) Run(adr string) error {
	return http.ListenAndServe(adr, a)
}

func main() {
	app := New()

	app.Get("/hello", func(c *Contex) {
		c.W.Write([]byte("GET, Zone01!"))
	})

	app.Post("/hello", func(c *Contex) {
		c.W.Write([]byte("POST, Zone01!"))
	})

	app.Run(":8080")
}
