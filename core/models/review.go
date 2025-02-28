package models

import (
	"time"
)

type Review struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	FilmID    uint      `gorm:"not null" json:"film_id"`
	Positive  bool      `gorm:"not null;default:false" json:"positive"`
	Content   string    `gorm:"not null" json:"content"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	User User `gorm:"foreignKey:UserID" json:"user"`
	Film Film `gorm:"foreignKey:FilmID"`
}
