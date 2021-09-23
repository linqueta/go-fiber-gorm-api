package book

import (
	"github.com/gofiber/fiber"
)

func GetBooks(c *fiber.Ctx) {
	c.Send("All books")
}

func GetBook(c *fiber.Ctx) {
	c.Send("A single book")
}

func NewBook(c *fiber.Ctx) {
	c.Send("Adds a new books")
}

func DeleteBook(c *fiber.Ctx) {
	c.Send("Deletes a books")
}
