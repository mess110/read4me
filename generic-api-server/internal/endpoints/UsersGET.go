package endpoints

import (
	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var Users = []User{
	{ID: "1", Name: "Alice"},
	{ID: "2", Name: "Bob"},
}

func UsersGET(c *fiber.Ctx) error {
	return c.JSON(Users)
}
