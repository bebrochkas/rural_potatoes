package films_tags

import (
	"github.com/bebrochkas/rural_potatoes/core/internal/db"
	"github.com/bebrochkas/rural_potatoes/core/models"
)

func InsertFilmWTags(film *models.Film, tags []models.Tag) error {
	for i := range tags {
		err := db.DB.Where("name = ?", tags[i].Name).FirstOrCreate(&tags[i]).Error
		if err != nil {
			return err
		}
	}

	err := db.DB.Where("title = ?", film.Title).FirstOrCreate(film).Error
	if err != nil {
		return err
	}

	// Теперь tags содержит обновлённые объекты с ID, и Append будет работать
	err = db.DB.Model(film).Association("Tags").Replace(tags) // или Append(tags)
	if err != nil {
		return err
	}
	return nil
}
