package tags

import (
	"github.com/bebrochkas/rural_potatoes/core/internal/db"
	"github.com/bebrochkas/rural_potatoes/core/models"
	"strings"
)

func SelectTags(names []string) ([]models.Tag, error) {
	var tags []models.Tag
	query := db.DB.Model(&models.Tag{})

	if len(names) > 0 {
		conditions := make([]string, len(names))
		args := make([]interface{}, len(names))

		for i, name := range names {
			conditions[i] = "name % ?"
			args[i] = name
		}

		whereClause := strings.Join(conditions, " OR ")
		query = query.Where(whereClause, args...)
	}

	if err := query.Find(&tags).Error; err != nil {
		return tags, err
	}
	return tags, nil
}
func SelectFilmTagsIds(filmId int) ([]models.Tag, error) {
	var film models.Film
	err := db.DB.Preload("Tags").First(&film, filmId).Error

	return film.Tags, err
}

func SelectFeedTags(userID uint) ([]string, error) {
	var tagIds []string
	err := db.DB.Model(models.UserTagScore{}).Where("user_id = ? AND score >= 2", userID).Pluck("tag_id", &tagIds).Error
	return tagIds, err
}
