package tag

import (
	"github.com/bebrochkas/rural_potatoes/core/internal/db/tags"
	"github.com/bebrochkas/rural_potatoes/core/internal/pb"
	"github.com/gofiber/fiber/v2"
)

func suggestTags(c *fiber.Ctx) error {
	prompt := c.Query("q")

	possibleTagNames, err := pb.GetTags(prompt)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	tags, err := tags.SelectTags(append(possibleTagNames, prompt))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.JSON(tags)
}
