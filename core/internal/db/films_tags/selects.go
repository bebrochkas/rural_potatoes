package films_tags

import (
	"github.com/bebrochkas/rural_potatoes/core/internal/db"
	"github.com/bebrochkas/rural_potatoes/core/models"
)

func SelectTagsFilms(offset int, limit int, tags []string) ([]models.Film, error) {

	var films []models.Film

	query := db.DB.Model(&models.Film{}).Preload("Tags")

	if tags[0] != "" {
		query = query.Joins("JOIN film_tags ON film_tags.film_id = films.id").
			Joins("JOIN tags ON tags.id = film_tags.tag_id").
			Where("tags.id IN (?)", tags)
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
