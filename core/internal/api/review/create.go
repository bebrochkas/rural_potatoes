package review

import (
	"github.com/bebrochkas/rural_potatoes/core/internal/crypto"
	"github.com/bebrochkas/rural_potatoes/core/internal/db/reviews"
	"github.com/bebrochkas/rural_potatoes/core/internal/db/tags"
	"github.com/bebrochkas/rural_potatoes/core/internal/db/users"
	"github.com/gofiber/fiber/v2"
)

func createReview(c *fiber.Ctx) error {

	var positive bool
	var coef float32
	var content_placeholder string
	switch c.Query("positive") {
	case "yes":
		coef = 1
		positive = true
		content_placeholder = "Лайк"
	case "no":
		coef = -1
		positive = false
		content_placeholder = "Дзлайк"
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "wrong positive query provided",
		})
	}

	filmID := c.QueryInt("id")
	content := c.Query("content", content_placeholder)

	tags, err := tags.SelectFilmTagsIds(filmID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	err = users.InsertTagScore(crypto.GetUserID(c), tags, coef)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	return reviews.InsertReview(crypto.GetUserID(c), uint(filmID), content, positive)

}
