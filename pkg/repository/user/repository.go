package user

import (
	"errors"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/domain"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/repository"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(sqlClient *gorm.DB) (*Repository, error) {
	r := &Repository{db: sqlClient}

	return r, r.validate()
}

func (r Repository) validate() error {
	if r.db == nil {
		return repository.ErrMissingDBClient
	}
	return nil
}

func (r Repository) Create(user *domain.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (r Repository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User

	err := r.db.First(&user, "email = ?", email).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.UserNotFoundError{Email: email}
		}

		return nil, err
	}

	return &user, nil
}

func (r Repository) GetUserFavorites(email string) ([]domain.Favorite, error) {
	user := &domain.User{}

	err := r.db.Model(&domain.User{}).Where("email = ?", email).
		Preload("Favorites").
		Preload("Favorites.Property", "status = ?", domain.ActiveStatus).
		Find(user).Error
	if err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, errors.New("user not found")
	}

	return user.Favorites, nil
}

func (r Repository) SaveFavorite(favorite *domain.Favorite) error {
	if err := r.db.Create(favorite).Error; err != nil {
		return err
	}

	return nil
}
