package user

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(api fiber.Router) {
	auth := api.Group("/auth")

	// auth
	auth.Post("/register", Register)
	auth.Post("/login", Login)

	//
}
