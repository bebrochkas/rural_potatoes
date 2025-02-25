package users

import (
	"github.com/bebrochkas/rural_potatoes/core/internal/db"
	"github.com/bebrochkas/rural_potatoes/core/models"
	"gorm.io/gorm"
)

func InsertUser(user models.User) (bool, error) {
	var existingUser models.User
	result := db.DB.Where("username = ?", user.Username).FirstOrCreate(&existingUser, user)

	if result.RowsAffected == 0 {
		return true, nil
	}
	return false, nil
}

func InsertTagScore(userID uint, tags []models.Tag, coef float32) error {
	for _, tag := range tags {

		var coef_weight float32
		switch tag.Type {
		case "thematic":
			coef_weight = 2
		case "counrty":
			coef_weight = 0.5
		case "realese":
			coef_weight = 0.4
		default:
			coef_weight = 0.2
		}

		tagScore := models.UserTagScore{}

		res := db.DB.Where(&models.UserTagScore{UserID: userID, TagID: tag.ID}).
			Attrs(models.UserTagScore{Score: coef * coef_weight}).
			FirstOrCreate(&tagScore, &models.UserTagScore{UserID: userID, TagID: tag.ID})

		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected > 0 {
			continue
		}
		if err := db.DB.Model(&models.UserTagScore{}).Where(&models.UserTagScore{UserID: userID, TagID: tag.ID}).Update("score", gorm.Expr("score + ?", coef*coef_weight)).Error; err != nil {
			return err
		}
	}

	return nil
}
