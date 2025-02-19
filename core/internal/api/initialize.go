package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/bebrochkas/rural_potatoes/core/internal/api/user"
	"github.com/gofiber/fiber/v2/middleware/cors"
	// jwtware "github.com/gofiber/contrib/jwt"
)

func Initialize() {

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173, http://localhost:5174",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Content-Type, Authorization",
	}))

	api := app.Group("/api")

	user.SetupRoutes(api)

	// auth

	// app.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: jwtware.SigningKey{Key: config.Cfg.JWT_TOKEN},
	// }))

	app.Listen(":3000")

}
