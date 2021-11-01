package user

import (
	"errors"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/domain"
)

type userRepository interface {
	Create(user *domain.User) error
	GetByEmail(email string) (*domain.User, error)
	SaveFavorite(favorite *domain.Favorite) error
	GetUserFavorites(email string) ([]domain.Favorite, error)
}

type propertyRepository interface {
	GetByID(id uint) (*domain.Property, error)
}

type Service struct {
	userRepository     userRepository
	propertyRepository propertyRepository
}

func New(userRepository userRepository,
	propertyRepository propertyRepository) (*Service, error) {

	service := &Service{
		userRepository:     userRepository,
		propertyRepository: propertyRepository,
	}

	return service, service.validate()
}

func (s Service) validate() error {
	if s.userRepository == nil {
		return errors.New("userRepository should not be nil")
	}
	if s.propertyRepository == nil {
		return errors.New("propertyRepository should not be nil")
	}
	return nil
}

func (s Service) Create(user *domain.User) error {
	return s.userRepository.Create(user)
}

func (s Service) Login(email string, password string) (bool, error) {
	user, err := s.userRepository.GetByEmail(email)
	if err != nil {
		return false, err
	}

	return user.Validate(email, password), nil
}

func (s Service) SetFavoriteProperty(propertyID uint, userEmail string) error {
	user, err := s.userRepository.GetByEmail(userEmail)
	if err != nil {
		return err
	}

	if _, err := s.propertyRepository.GetByID(propertyID); err != nil {
		return err
	}

	fav := domain.Favorite{
		UserID:     user.ID,
		PropertyID: propertyID,
	}

	return s.userRepository.SaveFavorite(&fav)
}

func (s Service) GetUserFavorites(userEmail string) ([]domain.Favorite, error) {
	f, err := s.userRepository.GetUserFavorites(userEmail)
	if err != nil {
		return nil, err
	}

	return f, nil
}
