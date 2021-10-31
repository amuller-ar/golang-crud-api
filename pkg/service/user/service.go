package user

import (
	"errors"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/domain"
)

type userRepository interface {
	Create(user *domain.User) error
	GetByEmail(email string) (*domain.User, error)
}

type Service struct {
	repository userRepository
}

func New(repository userRepository) (*Service, error) {
	service := &Service{repository: repository}

	return service, service.validate()
}

func (s Service) validate() error {
	if s.repository == nil {
		return errors.New("repository should not be nil")
	}
	return nil
}

func (s Service) Create(user *domain.User) error {
	return s.repository.Create(user)
}

func (s Service) Login(email string, password string) bool {
	user, err := s.repository.GetByEmail(email)
	if err != nil {
		return false
	}

	return user.Email == email && user.Password == password
}

func (s Service) SetFavoriteProperty(propertyID uint, userID uint) error {
	return nil
}
