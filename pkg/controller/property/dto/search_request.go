package dto

import (
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/domain"
	"github.com/gin-gonic/gin"
)

type SearchRequest struct {
	Status      *string `form:"status,omitempty" validation:"omitempty" binding:"oneof= ALL ACTIVE INACTIVE INVALID"`
	Page        *int    `form:"pageNumber,omitempty" binding:"min=1"`
	PageSize    *int    `form:"pageSize,omitempty" binding:"min=1,max=20"`
	BoundingBox *BoundingBox
}

type BoundingBox struct {
	MinLongitude float64 `form:"minLongitude,omitempty" binding:"required"`
	MinLatitude  float64 `form:"minLatitude,omitempty" binding:"required"`
	MaxLongitude float64 `form:"maxLongitude,omitempty" binding:"required"`
	MaxLatitude  float64 `form:"maxLatitude,omitempty" binding:"required"`
}

func (r SearchRequest) GetStatus() *string {
	if r.Status == nil {
		return nil
	}

	if *r.Status == domain.AllStatus {
		return nil
	}

	return r.Status
}

func (r SearchRequest) GetPage() int {
	if r.Page == nil {
		return domain.DefaultPage
	}

	return *r.Page
}

func (r SearchRequest) GetPageSize() int {
	if r.PageSize == nil {
		return domain.DefaultPageSize
	}

	return *r.PageSize
}

func (r SearchRequest) GetBoundingBox() *domain.BoundingBox {
	if r.BoundingBox == nil {
		return nil
	}

	return &domain.BoundingBox{
		MinLongitude: r.BoundingBox.MinLongitude,
		MinLatitude:  r.BoundingBox.MinLatitude,
		MaxLongitude: r.BoundingBox.MaxLongitude,
		MaxLatitude:  r.BoundingBox.MaxLatitude,
	}
}

func NewSearchParameters(ctx *gin.Context) (*domain.SearchParameters, error) {
	var request SearchRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		return nil, err
	}

	return &domain.SearchParameters{
		Status:      request.GetStatus(),
		BoundingBox: request.GetBoundingBox(),
		Page:        request.GetPage(),
		PageSize:    request.GetPageSize(),
	}, nil
}
