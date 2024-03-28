package talk

import (
	"net/http"
)

type Contex struct {
	R *http.Request
	W http.ResponseWriter
}

func (c *Contex) Param(key string) string {
	return c.R.URL.Query().Get(key)
}

func (c *Contex) WriteString(s string) (int, error) {
	return c.W.Write([]byte(s))
}

func (c *Contex) Status(statusCode int) *Contex {
	c.W.WriteHeader(statusCode)
	return c
}

// ...
