package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

const (
	House     = "HOUSE"
	Apartment = "APARTMENT"

	ActiveStatus   = "ACTIVE"
	InactiveStatus = "INACTIVE"
)

type Property struct {
	gorm.Model
	Title        string   `gorm:"text"`
	Description  string   `gorm:"text"`
	Location     Location `gorm:"embedded"`
	Pricing      Pricing  `gorm:"embedded"`
	PropertyType string   `gorm:"text"`
	BedRooms     uint
	BathRooms    uint
	ParkingSpots *uint
	Area         uint
	Photos       pq.StringArray `gorm:"type:text[]"`
	Status       string         `gorm:"type:text"`
}
