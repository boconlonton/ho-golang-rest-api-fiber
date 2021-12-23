package main

import (
	"encoding/json"
	"github.com/boconlonton/ho-golang-rest-api-fiber/database"
	"github.com/boconlonton/ho-golang-rest-api-fiber/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Start a new fiber
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// Connect to the Database
	database.ConnectDB()

	// Set up the router
	router.SetupRoutes(app)

	// Listen on PORT:3000
	app.Listen(":3000")
}
