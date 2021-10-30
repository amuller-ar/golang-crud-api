package dto

import (
	"fmt"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/domain"
	validation "github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/infrastructure/validator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UpdatePropertyRequest struct {
	ID           uint     `uri:"id" validate:"required"`
	Title        string   `json:"title" validate:"required"`
	Description  string   `json:"description"`
	Status       string   `json:"status" validation:"oneof=ACTIVE INACTIVE"`
	Location     Location `json:"location" validate:"required"`
	Pricing      Pricing  `json:"pricing"`
	PropertyType string   `json:"propertyType" validate:"required"`
	Bedrooms     uint     `json:"bedrooms" validate:"required"`
	Bathrooms    uint     `json:"bathrooms" validate:"required"`
	ParkingSpots *uint    `json:"parkingSpots,omitempty"`
	Area         uint     `json:"area" validate:"required"`
	Photos       []string `json:"photos"`
}

func (u UpdatePropertyRequest) ToProperty() domain.Property {
	return domain.Property{
		Model: gorm.Model{
			ID: u.ID,
		},
		Title:       u.Title,
		Description: u.Description,
		Location: domain.Location{
			Latitude:  u.Location.Latitude,
			Longitude: u.Location.Longitude,
		},
		Pricing: domain.Pricing{
			SalePrice:         u.Pricing.SalePrice,
			AdministrativeFee: u.Pricing.AdministrativeFee,
		},
		PropertyType: u.PropertyType,
		BedRooms:     u.Bedrooms,
		BathRooms:    u.Bathrooms,
		ParkingSpots: u.ParkingSpots,
		Area:         u.Area,
		Photos:       u.Photos,
		Status:       u.Status,
	}
}

func NewUpdatePropertyRequest(ctx *gin.Context) (*UpdatePropertyRequest, error) {
	var request UpdatePropertyRequest
	if err := ctx.ShouldBindUri(&request); err != nil {
		return nil, fmt.Errorf("error binding URI request. cause: %v", err)
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		return nil, fmt.Errorf("error binding json request. cause: %v", err)
	}

	v := validator.New()
	v.RegisterStructValidation(updateRequestValidator, UpdatePropertyRequest{})

	if err := validation.ValidateWithCustom(v, request); err != nil {
		return nil, err
	}

	return &request, nil
}

func updateRequestValidator(sl validator.StructLevel) {
	request := sl.Current().Interface().(UpdatePropertyRequest)

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

	if request.Status != domain.ActiveStatus && request.Status != domain.InactiveStatus {
		sl.ReportError(request.Pricing.SalePrice,
			"Status",
			"Status",
			validation.InvalidStatusValidationTag,
			"")
	}
}
