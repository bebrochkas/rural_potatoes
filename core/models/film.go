package models

type Film struct {
	ID           uint    `gorm:"primaryKey" json:"id"`
	Title        string  `gorm:"not null" json:"title"`
	Description  string  `gorm:"not null" json:"description"`
	PosterPreUrl string  `gorm:"not null" json:"posterPreUrl"`
	PosterUrl    string  `gorm:"not null" json:"posterUrl"`
	BackdropUrl  string  `gorm:"not null" json:"backdropUrl"`
	Rate         float32 `gorm:"not null" json:"rate"`
	Likes        uint    `gorm:"not null;default:0" json:"likes"`
	Dislikes     uint    `gorm:"not null;default:0" json:"dislikes"`
	Tags         []Tag   `gorm:"many2many:film_tags;" json:"tags"`
}
