package property

import (
	"errors"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/domain"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/repository"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/repository/utils"
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

func (r Repository) Create(property domain.Property) (*domain.Property, error) {
	if err := r.db.Create(&property).Error; err != nil {
		return nil, err
	}

	return &property, nil
}

func (r Repository) Update(property domain.Property) error {
	var result domain.Property

	err := r.db.
		Model(&domain.Property{}).
		Where("id = ?", property.Model.ID).
		First(&result).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.PropertyNotFoundError{ID: property.Model.ID}
		}

		return err
	}

	property.Model = result.Model

	if err := r.db.Save(&property).Error; err != nil {
		return err
	}

	return nil
}

func (r Repository) GetProperties() ([]domain.Property, error) {
	var props []domain.Property

	if err := r.db.Find(&props).Error; err != nil {
		return nil, err
	}

	return props, nil
}

func (r Repository) Search(params *domain.SearchParameters) (*domain.PaginatedResponse, error) {
	var rows []domain.Property
	pagination := domain.Pagination{
		Page:  params.Page,
		Limit: params.PageSize,
		Sort:  "updated_at",
	}

	query := r.db.Scopes(utils.Paginate(rows, &pagination, r.db))

	if params.Status != nil {
		query = query.Where(&domain.Property{Status: *params.Status})
	}

	if params.BoundingBox != nil {
		query = query.
			Where("longitude >= ?", params.BoundingBox.MinLongitude).
			Where("longitude <= ?", params.BoundingBox.MaxLongitude).
			Where("latitude >= ?", params.BoundingBox.MinLatitude).
			Where("latitude <= ?", params.BoundingBox.MaxLatitude)
	}

	if err := query.Find(&rows).Error; err != nil {
		return nil, err
	}

	return &domain.PaginatedResponse{
		Page:       pagination.Page,
		PageSize:   pagination.Limit,
		Total:      pagination.TotalRows,
		TotalPages: pagination.TotalPages,
		Data:       rows,
	}, nil
}

func (r Repository) GetByID(id uint) (*domain.Property, error) {
	var prop domain.Property

	q := r.db.First(&prop, id)

	if err := q.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.PropertyNotFoundError{ID: id}
		}
	}

	return &prop, nil
}
