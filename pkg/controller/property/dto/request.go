package dto

type CreatePropertyRequest struct {
	Title    string `json:"title" binding:"required"`
	Location struct {
		Longitude float64 `json:"longitude" binding:"required"`
		Latitude  float64 `json:"latitude"`
	} `json:"location" binding:"required"`
	Pricing struct {
		SalePrice int `json:"salePrice"`
	} `json:"pricing" binding:"required"`
	PropertyType string   `json:"propertyType" binding:"required"`
	Bedrooms     int      `json:"bedrooms" binding:"required"`
	Bathrooms    int      `json:"bathrooms" binding:"required"`
	ParkingSpots int      `json:"parkingSpots"`
	Area         int      `json:"area" binding:"required"`
	Photos       []string `json:"photos"`
}
