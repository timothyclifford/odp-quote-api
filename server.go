package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := Setup()
	log.Fatal(app.Listen(":3000"))
}

func Setup() *fiber.App {
	app := fiber.New()

	app.Get("/", dashboard)
	app.Post("/quotes", createQuote)

	return app
}
