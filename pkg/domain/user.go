package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email         string
	Password      string
	PropertiesFav []Favorite `gorm:"foreignKey:UserID"`
}

type Favorite struct {
	gorm.Model
	UserID   uint
	Property Property `gorm:"foreignKey:ID"`
}
