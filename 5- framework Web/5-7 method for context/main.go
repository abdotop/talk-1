package main

import (
	talk "func/app"
	"net/http"
)

func main() {
	app := talk.New()
	// talk

	app.Get("/hello", func(c *talk.Contex) {
		c.WriteString("GET Zone01")
	})

	app.Post("/hello", func(c *talk.Contex) {
		c.WriteString(c.Param("name"))
	})

	app.Put("/hello", func(c *talk.Contex) {
		c.Status(http.StatusBadRequest).WriteString("PUT Zone01")
	})

	app.Run(":8080")
}
