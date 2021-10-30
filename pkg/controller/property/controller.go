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
	Create(property domain.Property) (*domain.Property, error)
	Update(property domain.Property) error
	GetProperties() ([]domain.Property, error)
	Search(parameters *domain.SearchParameters) (*domain.PaginatedResponse, error)
}

type Controller struct {
	propertyService propertyService
}

func (c Controller) Create(ctx *gin.Context) error {
	request, err := dto.NewCreatePropertyRequest(ctx)
	if err != nil {
		return rest.NewError(http.StatusBadRequest, err.Error(), err)
	}

	prop, err := c.propertyService.Create(request.ToProperty())
	if err != nil {
		return rest.NewError(http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusCreated, prop)
	return nil
}

func (c Controller) Search(ctx *gin.Context) error {
	params, err := dto.NewSearchParameters(ctx)
	if err != nil {
		return rest.NewError(http.StatusBadRequest, err.Error())
	}

	response, err := c.propertyService.Search(params)
	if err != nil {
		return rest.NewError(http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, response)
	return nil
}

func (c Controller) Update(ctx *gin.Context) error {
	request, err := dto.NewUpdatePropertyRequest(ctx)
	if err != nil {
		return rest.NewError(http.StatusBadRequest, err.Error())
	}

	if err := c.propertyService.Update(request.ToProperty()); err != nil {
		switch err.(type) {
		case domain.PropertyNotFoundError:
			return rest.NewError(http.StatusNotFound, err.Error())
		default:
			return rest.NewError(http.StatusInternalServerError, err.Error())

		}
	}

	ctx.JSON(http.StatusOK, nil)
	return nil
}

func New(propertyService propertyService) (*Controller, error) {
	if propertyService == nil {
		return nil, errors.New("propertyService can't be nil")
	}

	return &Controller{propertyService: propertyService}, nil
}
