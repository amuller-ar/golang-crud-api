package dto

import (
	"fmt"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/domain/models"
	validation "github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/infrastructure/validator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreatePropertyRequest struct {
	ID           *uint    `json:"id,omitempty" validate:"omitempty"`
	Title        string   `json:"title" validate:"required"`
	Status       string   `json:"status"`
	Location     Location `json:"location" validate:"required"`
	Pricing      Pricing  `json:"pricing"`
	PropertyType string   `json:"propertyType" validate:"required"`
	Bedrooms     uint     `json:"bedrooms" validate:"required"`
	Bathrooms    uint     `json:"bathrooms" validate:"required"`
	ParkingSpots *uint    `json:"parkingSpots,omitempty" validate:"omitempty"`
	Area         uint     `json:"area" validate=:"required"`
	Photos       []string `json:"photos"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Pricing struct {
	SalePrice         float64 `json:"salePrice" validate:"required"`
	AdministrativeFee float64 `json:"administrative_fee"`
}

func NewCreatePropertyRequest(ctx *gin.Context) (*CreatePropertyRequest, error) {
	var request CreatePropertyRequest

	if err := ctx.BindJSON(&request); err != nil {
		return nil, fmt.Errorf("error binding json request. cause: %v", err)
	}

	validator := validator.New()
	validator.RegisterStructValidation(createPropertyRequestValidator, CreatePropertyRequest{})

	if err := validation.ValidateWithCustom(validator, request); err != nil {
		return nil, err
	}

	return &request, nil
}

func (r CreatePropertyRequest) ToProperty() models.Property {
	return models.Property{
		Title: r.Title,
		Location: models.Location{
			Latitude:  r.Location.Latitude,
			Longitude: r.Location.Longitude,
		},
		Pricing: models.Pricing{
			SalePrice:         r.Pricing.SalePrice,
			AdministrativeFee: r.Pricing.AdministrativeFee,
		},
		PropertyType: r.PropertyType,
		BedRooms:     r.Bedrooms,
		BathRooms:    r.Bathrooms,
		ParkingSpots: r.ParkingSpots,
		Area:         r.Area,
		Photos:       r.Photos,
	}
}

func NewCreatePropertyResponse(p *models.Property) CreatePropertyRequest {
	return CreatePropertyRequest{
		ID:     &p.ID,
		Title:  p.Title,
		Status: p.Status,
		Location: Location{
			Latitude:  p.Location.Latitude,
			Longitude: p.Location.Longitude,
		},
		Pricing: Pricing{
			SalePrice:         p.Pricing.SalePrice,
			AdministrativeFee: p.Pricing.AdministrativeFee,
		},
		PropertyType: p.PropertyType,
		Bedrooms:     p.BedRooms,
		Bathrooms:    p.BathRooms,
		ParkingSpots: p.ParkingSpots,
		Area:         p.Area,
		Photos:       p.Photos,
	}
}

func createPropertyRequestValidator(sl validator.StructLevel) {
	request := sl.Current().Interface().(CreatePropertyRequest)

	var maxBedrooms,
		minBedrooms,
		minBathrooms,
		maxBathrooms,
		minArea,
		maxArea,
		minParkingSpots uint

	switch request.PropertyType {
	case models.House:
		minBedrooms = 1
		maxBedrooms = 14
		minBathrooms = 1
		maxBathrooms = 12
		minArea = 50
		maxArea = 3000
		minParkingSpots = 0
		break
	case models.Apartment:
		minBedrooms = 1
		maxBedrooms = 6
		minBathrooms = 1
		maxBathrooms = 4
		minArea = 40
		maxArea = 400
		minParkingSpots = 1
		break
	default:
		break
	}

	if request.Bedrooms < minBedrooms || request.Bedrooms > maxBedrooms {
		sl.ReportError(request.Bedrooms, "bedrooms", "Bedrooms", "", "")
	}

	if request.Bathrooms < minBathrooms || request.Bathrooms > maxBathrooms {
		sl.ReportError(request.Bathrooms, "bathrooms", "Bathrooms", "", "")
	}

	if request.Area < minArea || request.Area > maxArea {
		sl.ReportError(request.Area, "area", "Area", "", "")
	}

	if request.ParkingSpots != nil {
		if *request.ParkingSpots < minParkingSpots {
			sl.ReportError(request.ParkingSpots, "parkingspots", "parkingSpots", "", "")
		}
	}

}
