package property

import (
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/domain"
	"github.com/gin-gonic/gin"
)

type Service struct {
}

func (s Service) Create(ctx *gin.Context, property domain.Property) (*domain.Property, error) {
	return &property, nil
}

func New() (*Service, error) {
	return &Service{}, nil
}
