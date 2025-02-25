package films

import (
	"github.com/bebrochkas/rural_potatoes/core/internal/db"
	"github.com/bebrochkas/rural_potatoes/core/models"
	"gorm.io/gorm"
)

func UpdateFilmRate(id int, like bool) error {
	switch like {
	case true:
		return db.DB.Model(models.Film{}).Where("id=?", id).Update("likes", gorm.Expr("likes + 1")).Error
	default:
		return db.DB.Model(models.Film{}).Where("id=?", id).Update("dislikes", gorm.Expr("dislikes + 1")).Error
	}

}
