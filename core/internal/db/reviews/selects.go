package reviews

import (
	"github.com/bebrochkas/rural_potatoes/core/internal/db"
	"github.com/bebrochkas/rural_potatoes/core/models"
)

func SelectFilmReviews(filmId int) ([]models.Review, error) {
	var reviews []models.Review
	if err := db.DB.Model(&models.Review{}).
		Preload("User"). // No Select clause
		Where("film_id = ?", filmId).
		Find(&reviews).Error; err != nil {
		return nil, err
	}
	return reviews, nil
}

func ProcessReviewsForFilm(userID, filmID uint) (uint, uint, *bool, error) {
	type Result struct {
		PositiveCount uint
		NegativeCount uint
	}

	var result Result

	err := db.DB.Model(&models.Review{}).
		Select("SUM(CASE WHEN positive = true THEN 1 ELSE 0 END) AS positive_count, "+
			"SUM(CASE WHEN positive = false THEN 1 ELSE 0 END) AS negative_count").
		Where("film_id = ?", filmID).
		Scan(&result).Error

	if err != nil {
		return 0, 0, nil, err
	}

	var userPositive bool
	if err := db.DB.Model(&models.Review{}).
		Select("positive").
		Where("user_id = ? AND film_id = ?", userID, filmID).
		Limit(1).
		Pluck("positive", &userPositive).Error; err != nil {
		return result.PositiveCount, result.NegativeCount, nil, nil
	}

	return result.PositiveCount, result.NegativeCount, &userPositive, nil
}
