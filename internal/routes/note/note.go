package note

import "github.com/gofiber/fiber/v2"

func SetupNoteRoutes(router fiber.Router) {
	note := router.Group("/note")

	// Create a Note
	note.Post("/", func(c *fiber.Ctx) error {})

	// List all Notes
	note.Get("/", func(c *fiber.Ctx) error {})

	// Read a Note
	note.Get("/:noteId", func(c *fiber.Ctx) error {})

	// Update a Note
	note.Put("/:noteId", func(c *fiber.Ctx) error {})

	// Delete a Note
	note.Delete("/:noteId", func(c *fiber.Ctx) error {})
}
