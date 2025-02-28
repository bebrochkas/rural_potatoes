package review

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(api fiber.Router) {
	auth := api.Group("/reviews")

	auth.Get("/", fetchReviews)
	auth.Post("/create", createReview)

}
