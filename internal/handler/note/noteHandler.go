package noteHandler

import (
	"github.com/definev/gonote/database"
	"github.com/definev/gonote/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetNotes(c *fiber.Ctx) error {
	db := database.DB

	var notes []model.Note

	// Find all note in database
	db.Find(&notes)

	// If no note found, return 404
	if len(notes) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"status":  "error",
			"message": "No note found",
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"code":    200,
		"status":  "success",
		"message": "Success",
		"data":    notes,
	})
}

func CreateNote(c *fiber.Ctx) error {
	db := database.DB
	note := new(model.Note)

	err := c.BodyParser(&note)
	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"code":    500,
				"status":  "error",
				"message": "Check your input",
				"data":    err,
			},
		)
	}

	note.ID = uuid.New()

	err = db.Create(note).Error
	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"code":    500,
				"status":  "error",
				"message": "Failed to create note",
				"data":    err,
			},
		)
	}

	return c.JSON(fiber.Map{
		"code":    200,
		"status":  "success",
		"message": "Success to create note",
		"data":    note,
	})
}

func GetNote(c *fiber.Ctx) error {
	db := database.DB
	var note model.Note

	id := c.Params("id")

	db.Find(&note, "id = ?", id)

	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"status":  "error",
			"message": "No note present",
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"code":    200,
		"status":  "success",
		"message": "Notes Found",
		"data":    note,
	})
}

func UpdateNote(c *fiber.Ctx) error {
	type updateNote struct {
		Title    string `json:"title"`
		SubTitle string `json:"sub_title"`
		Text     string `json:"Text"`
	}

	db := database.DB
	var note model.Note

	id := c.Params("id")

	db.Find(&note, "id = ?", id)

	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"status":  "error",
			"message": "No note founded",
			"data":    nil,
		})
	}

	updateNoteParams := updateNote{}
	err := c.BodyParser(&updateNoteParams)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"code":    500,
			"status":  "error",
			"message": "Check your input",
			"data":    err,
		})
	}

	note.Text = updateNoteParams.Text
	note.Title = updateNoteParams.Title
	note.Subtitle = updateNoteParams.SubTitle

	db.Save(&note)

	return c.JSON(fiber.Map{
		"code":    200,
		"status":  "success",
		"message": "Success to update note",
		"data":    note,
	})

}

func DeleteNote(c *fiber.Ctx) error {
	db := database.DB
	var note model.Note

	id := c.Params("id")

	db.Find(&note, "id = ?", id)

	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"status":  "error",
			"message": "No note founded",
			"data":    nil,
		})
	}

	err := db.Delete(&note).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"code":    500,
			"status":  "error",
			"message": "Failed to delete note",
			"data":    err,
		})
	}

	return c.JSON(fiber.Map{
		"code":    200,
		"status":  "success",
		"message": "Success to delete note",
		"data":    note,
	})
}
