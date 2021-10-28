package domain

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

const (
	House     = "HOUSE"
	Apartment = "APARTMENT"

	ActiveStatus   = "ACTIVE"
	InactiveStatus = "INACTIVE"

	MinMexSellPrice     = 1
	MaxMexSellPrice     = 15000000
	DefaultMinPrice     = 50000000
	DefaultMaxSellPrice = 3500000000
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

func (p Property) IsInBoundingBox(box BoundingBox) bool {
	return box.InBoundingBox(p.Location)
}
