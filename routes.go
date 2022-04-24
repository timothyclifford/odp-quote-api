package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func dashboard(c *fiber.Ctx) error {
	allQuotes := getAllQuotes()
	return c.JSON(allQuotes)
}

func createQuote(c *fiber.Ctx) error {
	q := new(Quote)

	if err := c.BodyParser(q); err != nil {
		return err
	}

	saveQuote(*q)

	return c.SendStatus(http.StatusCreated)
}

func getStubQuote() Quote {
	return Quote{"abc", "abc", 0, "first", "last", "email", "phone", "address", "suburb", "postcode"}
}
