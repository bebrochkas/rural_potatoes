package film

import (
	"strings"

	"github.com/bebrochkas/rural_potatoes/core/internal/crypto"
	"github.com/bebrochkas/rural_potatoes/core/internal/db/films"
	"github.com/bebrochkas/rural_potatoes/core/internal/db/reviews"
	"github.com/bebrochkas/rural_potatoes/core/internal/db/tags"
	"github.com/bebrochkas/rural_potatoes/core/models"
	"github.com/gofiber/fiber/v2"
)

func fetchFilms(c *fiber.Ctx) error {

	prompt := c.Query("q")

	offset := c.QueryInt("offset", 0)

	limit := c.QueryInt("limit", 30)

	tags_query := c.Query("tags", "")

	user_id := crypto.GetUserID(c)

	var err error
	var tagIds []string
	var strict bool = true

	switch tags_query {
	case "feed":
		tagIds, err = tags.SelectFeedTags(user_id)
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

	var filmsWreviews []models.Film

	for _, film := range films {
		likes, dislikes, userPositive, err := reviews.ProcessReviewsForFilm(user_id, film.ID)
		if err != nil {
			return err
		}

		film.Likes = likes
		film.Dislikes = dislikes
		film.UserPositive = userPositive

		filmsWreviews = append(filmsWreviews, film)

	}

	return c.JSON(filmsWreviews)

}
