package tags

import (
	"github.com/bebrochkas/rural_potatoes/core/internal/db"
	"github.com/bebrochkas/rural_potatoes/core/models"
)

func SelectTags(names []string) ([]models.Tag, error) {
	var tags []models.Tag
	query := db.DB.Model(&models.Tag{}).Order(`
			CASE
				WHEN type = 'thematic' THEN 1
				WHEN type = 'age rating' THEN 2
				WHEN type = 'rate' THEN 3
				WHEN type = 'release' THEN 4
				WHEN type = 'country' THEN 5
				ELSE 6
			END
		`)

	if len(names) > 0 {
		query = query.Where("name IN ?", names)
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
