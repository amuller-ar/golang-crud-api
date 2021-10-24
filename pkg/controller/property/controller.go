package property

import (
	"errors"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/controller/property/dto"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/domain/models"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/infrastructure/rest"
	"github.com/gin-gonic/gin"
	"net/http"
)

type propertyService interface {
	Create(ctx *gin.Context, property models.Property) (*models.Property, error)
	GetProperties() ([]models.Property, error)
}

type Controller struct {
	propertyService propertyService
}

func (c Controller) Create(ctx *gin.Context) error {
	request, err := dto.NewCreatePropertyRequest(ctx)
	if err != nil {
		return rest.NewError(http.StatusBadRequest, err.Error(), err)
	}

	prop, err := c.propertyService.Create(ctx, request.ToProperty())
	if err != nil {
		return rest.NewError(http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusCreated, dto.NewCreatePropertyResponse(prop))
	return nil
}

func (c Controller) GetProperties(ctx *gin.Context) error {
	result, err := c.propertyService.GetProperties()
	if err != nil {
		return rest.NewError(http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, result)
	return nil
}

func New(propertyService propertyService) (*Controller, error) {
	if propertyService == nil {
		return nil, errors.New("propertyService can't be nil")
	}

	return &Controller{propertyService: propertyService}, nil
}
