package main

import (
	"github.com/definev/gonote/database"
	"github.com/definev/gonote/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Connect to the Database
	database.ConnectDB()

	router.SetupRoutes(app)

	app.Listen(":3000")
}
