package films_tags

import (
	"github.com/bebrochkas/rural_potatoes/core/internal/db"
	"github.com/bebrochkas/rural_potatoes/core/models"
	"gorm.io/gorm"
)

func SelectTagsFilms(offset int, limit int, tags []string) ([]models.Film, error) {
	var films []models.Film

	query := db.DB.Model(&models.Film{}).
		Preload("Tags", func(db *gorm.DB) *gorm.DB {
			return db.Order(`
			CASE
				WHEN type = 'rate' THEN 1
				WHEN type = 'age rating' THEN 2
				WHEN type = 'release' THEN 3
				WHEN type = 'country' THEN 4
				WHEN type = 'thematic' THEN 5
				ELSE 6
			END
		`)
		})

	if len(tags) > 0 && tags[0] != "" {
		query = query.Joins("JOIN film_tags ON film_tags.film_id = films.id").
			Joins("JOIN tags ON tags.id = film_tags.tag_id").
			Where("tags.id IN (?)", tags).
			Group("films.id").
			Having("COUNT(DISTINCT tags.id) = ?", len(tags))
	}

	query = query.Limit(limit).Offset(offset)

	err := query.Find(&films).Error
	if err != nil {
		return films, err
	}

	return films, nil
}

func SelectTags() ([]models.Tag, error) {
	var tags []models.Tag
	if err := db.DB.Find(&tags).Error; err != nil {
		return tags, err
	}
	return tags, nil
}
