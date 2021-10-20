package property

import (
	"errors"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/controller/property/dto"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/domain"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/infrastructure/rest"
	"github.com/gin-gonic/gin"
	"net/http"
)

type propertyService interface {
	Create(ctx *gin.Context, property domain.Property) (*domain.Property, error)
}

type Controller struct {
	propertyService propertyService
}

func (c Controller) Create(ctx *gin.Context) error {
	var request dto.CreatePropertyRequest

	if err := ctx.ShouldBind(&request); err != nil {
		return rest.NewError(http.StatusBadRequest, err.Error(), err)
	}

	prop, err := c.propertyService.Create(ctx, domain.Property{})
	if err != nil {
		return rest.NewError(http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, dto.NewCreatePropertyResponse(prop))
	return nil
}

func New(propertyService propertyService) (*Controller, error) {
	if propertyService == nil {
		return nil, errors.New("propertyService can't be nil")
	}

	return &Controller{propertyService: propertyService}, nil
}
