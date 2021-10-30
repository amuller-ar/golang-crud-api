package service

import (
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/domain"
	"github.com/stretchr/testify/mock"
)

type PropertyService struct {
	mock.Mock
}

func (s *PropertyService) Create(property domain.Property) (*domain.Property, error) {
	args := s.Called(property)

	var p *domain.Property
	if args.Get(0) != nil {
		p = args.Get(0).(*domain.Property)
	}

	return p, args.Error(1)
}

func (s *PropertyService) Update(property domain.Property) error {
	args := s.Called(property)

	return args.Error(0)
}

func (s *PropertyService) GetProperties() ([]domain.Property, error) {
	args := s.Called(nil)

	return args.Get(0).([]domain.Property), args.Error(1)
}

func (s *PropertyService) Search(parameters *domain.SearchParameters) (*domain.PaginatedResponse, error) {
	args := s.Called(parameters)

	return args.Get(0).(*domain.PaginatedResponse), args.Error(1)
}
