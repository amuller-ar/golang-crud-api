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

	//SearchFavorites(parameters domain.FavoriteSearchParameters) (*domain.PaginatedResponse, error)
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

func (s Service) SetFavoriteProperty(propertyID uint, userEmail string) error {
	user, err := s.repository.GetByEmail(userEmail)
	if err != nil {
		return err
	}

	fav := domain.Favorite{
		UserID:     user.ID,
		PropertyID: propertyID,
	}

	return s.repository.SaveFavorite(&fav)
}

func (s Service) GetUserFavorites(userEmail string) ([]domain.Favorite, error) {
	favs, err := s.repository.GetUserFavorites(userEmail)
	if err != nil {
		return nil, err
	}

	return favs, nil
}
