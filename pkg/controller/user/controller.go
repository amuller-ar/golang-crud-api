package user

import (
	"errors"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/auth"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/controller/user/dto"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/domain"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/infrastructure/rest"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userService interface {
	Create(user *domain.User) error
	Login(email string, password string) bool
	SetFavoriteProperty(propertyID uint, userEmail string) error
	GetUserFavorites(userEmail string) ([]domain.Favorite, error)
}

type authService interface {
	CreateToken(userName string) (*auth.TokenDetails, error)
	ExtractTokenMetadata(*http.Request) (*auth.AccessDetails, error)
}

type Controller struct {
	userService userService
	authService authService
}

func New(userService userService, authService authService) (*Controller, error) {
	c := &Controller{
		userService: userService,
		authService: authService,
	}

	return c, c.validate()
}

func (c *Controller) validate() error {
	if c.userService == nil {
		return errors.New("service should not be nil")
	}
	if c.authService == nil {
		return errors.New("authService should not be nil")
	}

	return nil
}

func (c *Controller) Create(ctx *gin.Context) error {
	request, err := dto.ParseCreateUserRequest(ctx)
	if err != nil {
		return rest.NewError(http.StatusBadRequest, err.Error(), err)
	}

	if err := c.userService.Create(request); err != nil {
		return rest.NewError(http.StatusInternalServerError, err.Error())
	}

	ctx.Status(http.StatusCreated)
	return nil
}

func (c *Controller) Login(ctx *gin.Context) error {
	var request dto.LoginRequest

	if err := ctx.ShouldBind(&request); err != nil {
		return errors.New("no data found")
	}

	authenticated := c.userService.Login(request.Email, request.Password)
	if authenticated {
		authToken, err := c.authService.CreateToken(request.Email)
		if err != nil {
			ctx.Status(http.StatusUnauthorized)
			return err
		}

		ctx.JSON(http.StatusOK, gin.H{
			"token": authToken,
		})

		return nil
	}

	ctx.JSON(http.StatusUnauthorized, nil)
	return nil
}

func (c *Controller) SetFavoriteProperty(ctx *gin.Context) error {
	request, err := dto.NewFavoriteRequest(ctx)
	if err != nil {
		return rest.NewError(http.StatusBadRequest, err.Error(), err)
	}

	metadata, err := auth.ExtractTokenMetadata(ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return nil
	}

	if err := c.userService.SetFavoriteProperty(request.PropertyId, metadata.UserName); err != nil {
		return rest.NewError(http.StatusInternalServerError, err.Error())
	}

	ctx.Status(http.StatusOK)
	return nil
}

func (c Controller) GetFavorites(ctx *gin.Context) error {

	metadata, err := auth.ExtractTokenMetadata(ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "unauthorized")
		return nil
	}

	response, err := c.userService.GetUserFavorites(metadata.UserName)
	if err != nil {
		return rest.NewError(http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, response)
	return nil
}
