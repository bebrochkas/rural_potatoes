package models

type Tag struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"uniqueIndex;not null" json:"name"`
	Hex  string `gorm:"not null" json:"Hex"`
	Type string `gorm:"not null" json:"Type"`
}
