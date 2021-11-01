package dto

import (
	"fmt"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/domain"
	validation "github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/infrastructure/validator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PropertyRequest struct {
	Title        string   `json:"title" validate:"required"`
	Description  string   `json:"description"`
	Location     Location `json:"location" validate:"required"`
	Pricing      Pricing  `json:"pricing"`
	PropertyType string   `json:"propertyType" validate:"required"`
	Bedrooms     uint     `json:"bedrooms" validate:"required"`
	Bathrooms    uint     `json:"bathrooms" validate:"required"`
	ParkingSpots *uint    `json:"parkingSpots,omitempty" validate:"omitempty"`
	Area         uint     `json:"area" validate:"required"`
	Photos       []string `json:"photos"`
}

type Location struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}

type Pricing struct {
	SalePrice         float64 `json:"salePrice" validate:"required"`
	AdministrativeFee float64 `json:"administrative_fee"`
}

func (l Location) ToLocation() domain.Location {
	return domain.Location{
		Longitude: l.Longitude,
		Latitude:  l.Latitude,
	}
}

func (r PropertyRequest) ToProperty() domain.Property {
	model := domain.Property{
		Title:       r.Title,
		Description: r.Description,
		Location: domain.Location{
			Latitude:  r.Location.Latitude,
			Longitude: r.Location.Longitude,
		},
		Pricing: domain.Pricing{
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

	return model
}

func NewCreatePropertyRequest(ctx *gin.Context) (*PropertyRequest, error) {
	var request PropertyRequest

	if err := ctx.BindJSON(&request); err != nil {
		return nil, fmt.Errorf("error binding json request. cause: %v", err)
	}

	v := validator.New()
	v.RegisterStructValidation(requestValidator, PropertyRequest{})

	if err := validation.ValidateWithCustom(v, request); err != nil {
		return nil, err
	}

	return &request, nil
}

func requestValidator(sl validator.StructLevel) {
	request := sl.Current().Interface().(PropertyRequest)

	c := GetConstraint(request.PropertyType)

	if request.Bedrooms < c.MinBedrooms ||
		request.Bedrooms > c.MaxBedrooms {
		sl.ReportError(request.Bedrooms,
			"bedrooms",
			"Bedrooms",
			validation.OutOfRangeValidationTAg,
			fmt.Sprintf("range %d - %d", c.MinBedrooms, c.MaxBedrooms))
	}

	if request.Bathrooms < c.MinBathrooms || request.Bathrooms > c.MaxBathrooms {
		sl.ReportError(request.Bathrooms,
			"bathrooms",
			"Bathrooms",
			validation.OutOfRangeValidationTAg,
			fmt.Sprintf("range %d - %d", c.MinBathrooms, c.MaxBathrooms))
	}

	if request.Area < c.MinArea || request.Area > c.MaxArea {
		sl.ReportError(request.Area,
			"area",
			"Area",
			validation.OutOfRangeValidationTAg,
			fmt.Sprintf("range %d - %d", c.MinArea, c.MaxArea))
	}

	if request.ParkingSpots != nil {
		if *request.ParkingSpots < c.MinParkingSpots {
			sl.ReportError(request.ParkingSpots,
				"parkingSpots",
				"parkingSpots",
				validation.MinValueValidationTag,
				fmt.Sprintf("min value %d", c.MinParkingSpots))
		}
	}

	var min float64
	var max float64
	if request.Location.ToLocation().InBoundingBox(domain.MexicoBBox) {
		min = domain.MinMexSellPrice
		max = domain.MaxMexSellPrice

	} else {
		min = domain.DefaultMinPrice
		max = domain.DefaultMaxSellPrice
	}

	if request.Pricing.SalePrice < min {
		sl.ReportError(request.Pricing.SalePrice,
			"salePrice",
			"salePrice",
			validation.MinValueValidationTag,
			"")
	}

	if request.Pricing.SalePrice > max {
		sl.ReportError(request.Pricing.SalePrice,
			"salePrice",
			"salePrice",
			validation.MaxValueValidationTag,
			"")
	}
}
