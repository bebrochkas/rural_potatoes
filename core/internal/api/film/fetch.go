package film

import (
	"strings"

	"github.com/bebrochkas/rural_potatoes/core/internal/crypto"
	"github.com/bebrochkas/rural_potatoes/core/internal/db/films"
	"github.com/bebrochkas/rural_potatoes/core/internal/db/tags"
	"github.com/gofiber/fiber/v2"
)

func fetchFilms(c *fiber.Ctx) error {

	prompt := c.Query("q")

	offset := c.QueryInt("offset", 0)

	limit := c.QueryInt("limit", 30)

	tags_query := c.Query("tags", "")

	var err error
	var tagIds []string
	var strict bool = true

	switch tags_query {
	case "feed":
		tagIds, err = tags.SelectFeedTags(crypto.GetUserID(c))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}
		strict = !strict

	default:
		tagIds = strings.Split(tags_query, ",")

	}

	films, err := films.SelectTagsFilms(offset, limit, strict, tagIds, prompt)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.JSON(films)

}
