package user

import (
	"github.com/bebrochkas/rural_potatoes/core/internal/crypto"
	"github.com/bebrochkas/rural_potatoes/core/internal/db/users"
	"github.com/bebrochkas/rural_potatoes/core/models"
	"github.com/gofiber/fiber/v2"
)

type authRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c *fiber.Ctx) error {
	var req authRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	user := models.User{
		Username: req.Username,
		Password: crypto.GeneratePassword(req.Password),
	}

	exists, err := users.InsertUser(user)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	switch exists {
	case true:
		return c.Status(200).JSON(fiber.Map{
			"message": "user already exists",
		})

	default:
		return c.Status(201).JSON(fiber.Map{
			"message": "user created",
		})
	}

}

func Login(c *fiber.Ctx) error {
	var req authRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	user, err := users.SelectUser(req.Username)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "user not found",
		})
	}
	if !crypto.ComparePassword(user.Password, req.Password) {
		return c.Status(400).JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	token, err := crypto.GenerateToken(user.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"token": token,
	})
}
