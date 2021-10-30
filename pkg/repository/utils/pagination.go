package utils

import (
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/domain"
	"gorm.io/gorm"
	"math"
)

func Paginate(value interface{}, pagination *domain.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())

	}
}
