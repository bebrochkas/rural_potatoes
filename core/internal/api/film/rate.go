package film

import (
	"github.com/bebrochkas/rural_potatoes/core/internal/crypto"
	"github.com/bebrochkas/rural_potatoes/core/internal/db/films"
	"github.com/bebrochkas/rural_potatoes/core/internal/db/tags"
	"github.com/bebrochkas/rural_potatoes/core/internal/db/users"
	"github.com/gofiber/fiber/v2"
)

func rateFilm(c *fiber.Ctx) error {

	filmId := c.QueryInt("id")

	tags, err := tags.SelectFilmTagsIds(filmId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	switch c.Query("action") {
	case "like":
		users.InsertTagScore(crypto.GetUserID(c), tags, +1)
		films.UpdateFilmRate(filmId, true)
	case "dislike":
		users.InsertTagScore(crypto.GetUserID(c), tags, -1)
		films.UpdateFilmRate(filmId, false)
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "uknown action provided",
		})
	}
	return nil
}
