package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/bebrochkas/rural_potatoes/core/config"
	"github.com/bebrochkas/rural_potatoes/core/internal/api/film"
	"github.com/bebrochkas/rural_potatoes/core/internal/api/tag"
	"github.com/bebrochkas/rural_potatoes/core/internal/api/user"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: config.Cfg.JWT_TOKEN},
	}))

	film.SetupRoutes(api)

	tag.SetupRoutes(api)

	app.Listen(":3000")

}
