package models

import ()

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Password string `gorm:"not null" json:"-"`
}

type UserTagScore struct {
	ID     uint    `gorm:"primaryKey" json:"id"`
	UserID uint    `gorm:"index;foreignKey:UserID;references:ID" json:"user_id"`
	TagID  uint    `gorm:"index;foreignKey:TagID;references:ID" json:"tag_id"`
	Score  float32 `json:"score"`

	User User `gorm:"foreignKey:UserID" json:"user"`
	Tag  Tag  `gorm:"foreignKey:TagID" json:"tag"`
}
