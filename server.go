package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	"github.com/gofiber/template/ace"
)

func main() {
	// App and configuration
	app := fiber.New()

	// Logger
	app.Use(logger.New())

	engine := ace.New("./views", ".ace")
	err := engine.Load()
	if err != nil {
		panic(err)
	}
	app.Settings.Views = engine

	app.Get("/", func(c *fiber.Ctx) {
		err := c.Render("index", nil, "layouts/main")
		if err != nil {
			panic(err)
		}
	})

	// Server UP
	app.Listen("127.0.0.1:3000")
}
