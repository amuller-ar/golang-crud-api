package domain

import "gorm.io/gorm"

type Property struct {
	gorm.Model
	Title        string
	Description  string
	Location     Location
	Pricing      Pricing
	PropertyType PropertyType
	BedRooms     uint
	BathRooms    uint
	ParkingSpots uint
	Area         uint
	Photos       []string
	Status       Status
}

type PropertyType string
type Status string
