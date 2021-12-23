package router

import (
	noteRoutes "github.com/boconlonton/ho-golang-rest-api-fiber/internal/routes/note"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// Setup the Note routes
	noteRoutes.SetupNoteRoutes(api)
}
