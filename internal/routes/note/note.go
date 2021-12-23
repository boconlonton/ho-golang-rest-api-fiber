package noteRoutes

import (
	noteHandler "github.com/boconlonton/ho-golang-rest-api-fiber/internal/handlers/note"
	"github.com/gofiber/fiber/v2"
)

func SetupNoteRoutes(router fiber.Router) {
	note := router.Group("/notes")

	// Create a Note
	note.Post("/", noteHandler.CreateNotes)

	// List all Notes
	note.Get("/", noteHandler.ListNotes)

	// Read a Note
	note.Get("/:noteId", noteHandler.GetNote)

	// Update a Note
	note.Put("/:noteId", noteHandler.UpdateNote)

	// Delete a Note
	note.Delete("/:noteId", noteHandler.DeleteNote)
}
