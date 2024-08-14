package endpoints

import (
	"github.com/gofiber/fiber/v2"
)

type Book struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

var Books = []Book{
	{ID: "1", Name: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams"},
	{ID: "2", Name: "Two Years' Vacation", Author: "Jules Verne"},
}

func BooksGET(c *fiber.Ctx) error {
	return c.JSON(Books)
}
