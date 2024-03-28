package main

import "github.com/abdotop/octopus"

func main() {
	app := octopus.New()

	app.Get("/", func(c *octopus.Ctx) {
		c.WriteString("Hello, Zone01!")
	})

	app.Run(":8080")
}
