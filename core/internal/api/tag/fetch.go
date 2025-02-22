package tag

import (
	"github.com/bebrochkas/rural_potatoes/core/internal/db/films_tags"
	"github.com/gofiber/fiber/v2"
)

func fetchTags(c *fiber.Ctx) error {
	tags, err := films_tags.SelectTags()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.JSON(tags)
}
