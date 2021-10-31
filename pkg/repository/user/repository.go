package user

import (
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

	err := r.db.Model(&domain.User{}).Where("email = ?", email).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
