package property

import (
	"errors"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/domain/models"
)

type propertyRepository interface {
	Create(property models.Property) (*models.Property, error)
	Update(property models.Property) error
	GetProperties() ([]models.Property, error)
}
type Service struct {
	repository propertyRepository
}

func (s Service) Create(property models.Property) (*models.Property, error) {

	if property.IsInBoundingBox(models.MexicoBBox) {
		property.Status = models.ActiveStatus
	}

	prop, err := s.repository.Create(property)
	if err != nil {
		return nil, err
	}

	return prop, nil
}

func (s Service) Update(property models.Property) error {
	return s.repository.Update(property)
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
