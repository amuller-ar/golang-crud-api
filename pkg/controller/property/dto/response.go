package dto

import "github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/domain"

type CreatePropertyResponse struct {
	ID       uint          `json:"id"`
	Title    string        `json:"title"`
	Status   domain.Status `json:"status"`
	Location struct {
		Longitude float64 `json:"longitude"`
		Latitude  float64 `json:"latitude"`
	} `json:"location"`
	Pricing struct {
		SalePrice int `json:"salePrice"`
	} `json:"pricing"`
	PropertyType string   `json:"propertyType"`
	Bedrooms     int      `json:"bedrooms"`
	Bathrooms    int      `json:"bathrooms"`
	ParkingSpots int      `json:"parkingSpots"`
	Area         int      `json:"area"`
	Photos       []string `json:"photos"`
}

func NewCreatePropertyResponse(property *domain.Property) CreatePropertyResponse {
	return CreatePropertyResponse{
		ID:     property.ID,
		Title:  property.Title,
		Status: property.Status,
	}
}
