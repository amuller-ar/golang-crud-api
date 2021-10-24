package property

import (
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/domain/models"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/repository"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func (r Repository) SaveProperty(property models.Property) (*models.Property, error) {
	if err := r.db.Create(&property).Error; err != nil {
		return nil, err
	}

	return &property, nil
}

func (r Repository) GetProperties() ([]models.Property, error) {
	var props []models.Property

	if err := r.db.Find(&props).Error; err != nil {
		return nil, err
	}

	return props, nil
}

func New(sqlClient *gorm.DB) (*Repository, error) {
	repository := &Repository{db: sqlClient}

	return repository, repository.validate()
}

func (r Repository) validate() error {
	if r.db == nil {
		return repository.ErrMissingDBClient
	}
	return nil
}
