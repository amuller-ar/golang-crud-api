package property

import (
	"errors"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/domain"
)

type propertyRepository interface {
	Create(property domain.Property) (*domain.Property, error)
	Update(property domain.Property) error
	GetProperties() ([]domain.Property, error)
	Search(parameters *domain.SearchParameters) (*domain.PaginatedResponse, error)
}

type Service struct {
	repository propertyRepository
}

func (s Service) Create(property domain.Property) (*domain.Property, error) {

	property.Status = domain.InactiveStatus
	if property.IsInBoundingBox(domain.MexicoBBox) {
		property.Status = domain.ActiveStatus
	}

	prop, err := s.repository.Create(property)
	if err != nil {
		return nil, err
	}

	return prop, nil
}

func (s Service) Update(property domain.Property) error {
	return s.repository.Update(property)
}

func (s Service) GetProperties() ([]domain.Property, error) {
	return s.repository.GetProperties()
}

func (s Service) Search(parameters *domain.SearchParameters) (*domain.PaginatedResponse, error) {
	return s.repository.Search(parameters)

}

func New(repository propertyRepository) (*Service, error) {
	service := &Service{repository: repository}

	return service, service.validate()
}

func (s Service) validate() error {
	if s.repository == nil {
		return errors.New("repository should not be nil")
	}
	return nil
}
