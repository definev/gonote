package noteRouters

import (
	noteHandler "github.com/definev/gonote/internal/handler/note"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
	note := router.Group("/note")

	note.Get("/", noteHandler.GetNotes)
	note.Post("/", noteHandler.CreateNote)
	note.Get("/:id", noteHandler.GetNote)
	note.Put("/:id", noteHandler.UpdateNote)
	note.Delete("/:id", noteHandler.DeleteNote)
}
