package dto

import (
	"fmt"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/infrastructure/validator"
	"github.com/gin-gonic/gin"
)

type FavoriteRequest struct {
	PropertyId uint `json:"propertyId" validate:"required"`
}

func NewFavoriteRequest(ctx *gin.Context) (*FavoriteRequest, error) {
	var request FavoriteRequest

	if err := ctx.BindJSON(&request); err != nil {
		return nil, fmt.Errorf("error binding json request. cause: %v", err)
	}

	if err := validator.Validate(request); err != nil {
		return nil, err
	}

	return &request, nil
}
