package users

import (
	"github.com/bebrochkas/rural_potatoes/core/internal/db"
	"github.com/bebrochkas/rural_potatoes/core/models"
)

func InsertUser(user models.User) (bool, error) {
	var existingUser models.User
	result := db.DB.Where("username = ?", user.Username).FirstOrCreate(&existingUser, user)

	if result.RowsAffected == 0 {
		return true, nil
	}
	return false, nil
}
