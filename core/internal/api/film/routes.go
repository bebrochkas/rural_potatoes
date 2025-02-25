package film

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(api fiber.Router) {
	auth := api.Group("/films")

	// fetch
	auth.Get("/", fetchFilms)

	//
	auth.Post("/rate", rateFilm)

}
