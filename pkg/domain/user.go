package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string
	Password  string
	Favorites []Favorite `gorm:"foreignkey:UserID"`
}

type Favorite struct {
	UserID     uint     `gorm:"primaryKey"`
	PropertyID uint     `gorm:"primaryKey"`
	Property   Property `gorm:"foreignkey:PropertyID"`
}

func (u User) Validate(email string, password string) bool {
	return u.Email == email && u.Password == password
}
