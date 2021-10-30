package factory

import (
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/domain"
	"gorm.io/gorm"
	"time"
)

func GetProperty() *domain.Property {
	return &domain.Property{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Title:        "Test",
		Description:  "Test Mock",
		Location:     domain.Location{},
		Pricing:      domain.Pricing{},
		PropertyType: domain.House,
		BedRooms:     1,
		BathRooms:    2,
		ParkingSpots: nil,
		Area:         30,
		Photos:       nil,
		Status:       "ACTIVE",
	}
}
