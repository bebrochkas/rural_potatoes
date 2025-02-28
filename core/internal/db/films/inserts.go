package films

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

	err = db.DB.Model(film).Association("Tags").Replace(tags)
	if err != nil {
		return err
	}
	return nil
}
