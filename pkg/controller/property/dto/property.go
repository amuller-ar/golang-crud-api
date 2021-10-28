package dto

import (
	"errors"
	"fmt"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/domain"
	validation "github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/infrastructure/validator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
)

type PropertyRequest struct {
	ID           *uint           `json:"id,omitempty" validate:"omitempty"`
	Title        string          `json:"title" validate:"required"`
	Description  string          `json:"description"`
	Status       string          `json:"status"`
	Location     domain.Location `json:"location" validate:"required"`
	Pricing      Pricing         `json:"pricing"`
	PropertyType string          `json:"propertyType" validate:"required"`
	Bedrooms     uint            `json:"bedrooms" validate:"required"`
	Bathrooms    uint            `json:"bathrooms" validate:"required"`
	ParkingSpots *uint           `json:"parkingSpots,omitempty" validate:"omitempty"`
	Area         uint            `json:"area" validate=:"required"`
	Photos       []string        `json:"photos"`
}

type Pricing struct {
	SalePrice         float64 `json:"salePrice" validate:"required"`
	AdministrativeFee float64 `json:"administrative_fee"`
}

func (r PropertyRequest) ToProperty() domain.Property {
	model := domain.Property{
		Title:       r.Title,
		Description: r.Description,
		Status:      r.Status,
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

	if r.ID != nil {
		model.Model.ID = *r.ID
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

func NewUpdatePropertyRequest(ctx *gin.Context) (*PropertyRequest, error) {
	str := ctx.Param("id")
	if str == "" {
		return nil, errors.New("id is required")
	}

	id, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return nil, err
	}

	var request PropertyRequest
	if err := ctx.BindJSON(&request); err != nil {
		return nil, fmt.Errorf("error binding json request. cause: %v", err)
	}

	v := validator.New()
	v.RegisterStructValidation(requestValidator, PropertyRequest{})

	if err := validation.ValidateWithCustom(v, request); err != nil {
		return nil, err
	}

	uintID := uint(id)
	request.ID = &uintID

	return &request, nil
}

func NewCreatePropertyResponse(p *domain.Property) PropertyRequest {
	return PropertyRequest{
		ID:          &p.ID,
		Title:       p.Title,
		Description: p.Description,
		Status:      p.Status,
		Location: domain.Location{
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

func requestValidator(sl validator.StructLevel) {
	request := sl.Current().Interface().(PropertyRequest)

	var maxBedrooms,
		minBedrooms,
		minBathrooms,
		maxBathrooms,
		minArea,
		maxArea,
		minParkingSpots uint

	switch request.PropertyType {
	case domain.House:
		minBedrooms = 1
		maxBedrooms = 14
		minBathrooms = 1
		maxBathrooms = 12
		minArea = 50
		maxArea = 3000
		minParkingSpots = 0
		break
	case domain.Apartment:
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
		sl.ReportError(request.Bedrooms,
			"bedrooms",
			"Bedrooms",
			validation.OutOfRangeValidationTAg,
			fmt.Sprintf("range %d - %d", minBedrooms, maxBedrooms))
	}

	if request.Bathrooms < minBathrooms || request.Bathrooms > maxBathrooms {
		sl.ReportError(request.Bathrooms,
			"bathrooms",
			"Bathrooms",
			validation.OutOfRangeValidationTAg,
			fmt.Sprintf("range %d - %d", minBathrooms, maxBathrooms))
	}

	if request.Area < minArea || request.Area > maxArea {
		sl.ReportError(request.Area,
			"area",
			"Area",
			validation.OutOfRangeValidationTAg,
			fmt.Sprintf("range %d - %d", minArea, maxArea))
	}

	if request.ParkingSpots != nil {
		if *request.ParkingSpots < minParkingSpots {
			sl.ReportError(request.ParkingSpots,
				"parkingSpots",
				"parkingSpots",
				validation.MinValueValidationTag,
				fmt.Sprintf("min value %d", minParkingSpots))
		}
	}

	var min float64
	var max float64
	if request.Location.InBoundingBox(domain.MexicoBBox) {
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
