package reviews

import (
	"github.com/bebrochkas/rural_potatoes/core/internal/db"
	"github.com/bebrochkas/rural_potatoes/core/models"
)

func InsertReview(userID uint, filmID uint, content string, positive bool) error {
	review := models.Review{}

	result := db.DB.
		Where(&models.Review{UserID: userID, FilmID: filmID}).
		FirstOrCreate(&review, models.Review{UserID: userID, FilmID: filmID})

	if result.Error != nil {
		return result.Error
	}

	updateResult := db.DB.
		Model(&review).
		Updates(map[string]interface{}{"Content": content, "Positive": positive})

	if updateResult.Error != nil {
		return updateResult.Error
	}

	return nil
}
