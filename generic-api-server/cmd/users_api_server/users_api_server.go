package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.uber.org/zap"

	"github.com/mess110/read4me/generic-api-server/internal/endpoints"
	"github.com/mess110/read4me/generic-api-server/internal/utils"
)

func main() {
	port := 3000
	server := "users_api_server"

	logger, err := utils.InitLogger()
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize logger %v", err))
	}
	defer logger.Sync()

	app := fiber.New()

	app.Use(utils.LoggerMiddleware())
	app.Use(requestid.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return endpoints.PingGET(c, server)
	})
	app.Get("/users", endpoints.UsersGET)
	app.Get("/users/:id", endpoints.UserGET)

	zap.L().Info("starting", zap.String("server", server), zap.Int("port", port))
	err = app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		zap.L().Fatal("failed to start", zap.String("server", server), zap.Error(err))
	}
}
