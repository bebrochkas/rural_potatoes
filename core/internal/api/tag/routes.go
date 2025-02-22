package tag

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(api fiber.Router) {
	auth := api.Group("/tags")

	// fetch
	auth.Get("/", fetchTags)

	//
}
