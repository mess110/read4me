package endpoints

import (
	"github.com/gofiber/fiber/v2"
)

func UserGET(c *fiber.Ctx) error {
	userID := c.Params("id")

	for _, user := range Users {
		if user.ID == userID {
			return c.JSON(user)
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(map[string]any{
		"error":  "User not found",
		"status": 404,
	})
}
