package property

import (
	"errors"

	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/domain/models"
	"github.com/gin-gonic/gin"
)

type propertyRepository interface {
	SaveProperty(property models.Property) (*models.Property, error)
	GetProperties() ([]models.Property, error)
}
type Service struct {
	repository propertyRepository
}

func (s Service) Create(ctx *gin.Context, property models.Property) (*models.Property, error) {
	prop, err := s.repository.SaveProperty(property)
	if err != nil {
		return nil, err
	}

	return prop, nil
}

func (s Service) GetProperties() ([]models.Property, error) {
	return s.repository.GetProperties()
}

func New(repository propertyRepository) (*Service, error) {
	return &Service{repository: repository}, nil
}

func (s Service) validate() error {
	if s.repository == nil {
		return errors.New("repository should not be nil")
	}
	return nil
}
