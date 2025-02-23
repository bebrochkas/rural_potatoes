package film

import (
	"strings"

	"github.com/bebrochkas/rural_potatoes/core/internal/db/films_tags"
	"github.com/gofiber/fiber/v2"
)

func fetchFilms(c *fiber.Ctx) error {
	offset := c.QueryInt("offset", 0)

	limit := c.QueryInt("limit", 30)

	tags := strings.Split(c.Query("tags", ""), ",")

	films, err := films_tags.SelectTagsFilms(offset, limit, tags)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.JSON(films)

}
