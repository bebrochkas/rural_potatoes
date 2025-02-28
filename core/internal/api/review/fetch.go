package review

import (
	"github.com/bebrochkas/rural_potatoes/core/internal/db/reviews"
	"github.com/gofiber/fiber/v2"
)

func fetchReviews(c *fiber.Ctx) error {
	reviews, err := reviews.SelectFilmReviews(c.QueryInt("id"))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}
	return c.JSON(reviews)
}
