package dto

import "github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/domain"

// PropertyConstraints represents data validation values for property requests
type PropertyConstraints struct {
	MaxBedrooms     uint
	MinBedrooms     uint
	MinBathrooms    uint
	MaxBathrooms    uint
	MinArea         uint
	MaxArea         uint
	MinParkingSpots uint
}

// constraints contains PropertyConstraints for each propertyType
var constraints = map[string]PropertyConstraints{
	domain.Apartment: {
		MaxBedrooms:     6,
		MinBedrooms:     1,
		MinBathrooms:    1,
		MaxBathrooms:    4,
		MinArea:         40,
		MaxArea:         400,
		MinParkingSpots: 1,
	},
	domain.House: {
		MaxBedrooms:     14,
		MinBedrooms:     1,
		MinBathrooms:    1,
		MaxBathrooms:    12,
		MinArea:         50,
		MaxArea:         3000,
		MinParkingSpots: 0,
	},
}

// GetConstraint gets propertyType constraint
func GetConstraint(propertyType string) PropertyConstraints {
	return constraints[propertyType]
}
