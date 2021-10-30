package property

import (
	"errors"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/domain"
)

type Service struct {
	CreateMock        func(property domain.Property) (*domain.Property, error)
	UpdateMock        func(property domain.Property) error
	GetPropertiesMock func() ([]domain.Property, error)
	SearchMock        func(parameters *domain.SearchParameters) (*domain.PaginatedResponse, error)
}

func (s Service) Create(property domain.Property) (*domain.Property, error) {
	if s.CreateMock == nil {
		return nil, errors.New("CreateMock nil")
	}

	return s.CreateMock(property)
}

func (s Service) Update(property domain.Property) error {
	if s.UpdateMock == nil {
		return errors.New("UpdateMock nil")
	}

	return s.UpdateMock(property)
}

func (s Service) GetProperties() ([]domain.Property, error) {
	if s.GetPropertiesMock == nil {
		return nil, errors.New("GetPropertiesMock nil")
	}

	return s.GetPropertiesMock()
}

func (s Service) Search(parameters *domain.SearchParameters) (*domain.PaginatedResponse, error) {
	if s.SearchMock == nil {
		return nil, errors.New("SearchMock nil")
	}

	return s.SearchMock(parameters)
}
