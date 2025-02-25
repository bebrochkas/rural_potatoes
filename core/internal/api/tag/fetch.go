package tag

import (
	"github.com/bebrochkas/rural_potatoes/core/internal/db/tags"
	"github.com/gofiber/fiber/v2"
)

func fetchTags(c *fiber.Ctx) error {
	tags, err := tags.SelectTags(nil)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.JSON(tags)
}
