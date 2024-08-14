package endpoints

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func PingGET(c *fiber.Ctx, server string) error {
	timeNow := time.Now().String()

	return c.JSON(map[string]any{
		"ping":   timeNow,
		"server": server,
	})
}
