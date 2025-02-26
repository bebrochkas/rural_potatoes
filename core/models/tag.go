package models

type Tag struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"uniqueIndex;not null" json:"name"`
	Hex   string `gorm:"not null" json:"hex"`
	Type  string `gorm:"not null" json:"type"`
	Films []Film `gorm:"many2many:film_tags;" json:"tags"`
}
