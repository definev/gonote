package router

import (
	noteRouters "github.com/definev/gonote/internal/routes/note"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	noteRouters.SetupRoutes(api)
}
