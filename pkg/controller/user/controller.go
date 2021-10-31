package user

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"net/http"

	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/controller/user/dto"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/domain"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/infrastructure/rest"
	"github.com/gin-gonic/gin"
)

type userService interface {
	Create(user *domain.User) error
	Login(email string, password string) bool
}

//jwt service
type jwtService interface {
	GenerateToken(email string, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
}

type Controller struct {
	userService userService
	jwtService  jwtService
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
		authToken := c.jwtService.GenerateToken(request.Email, true)
		ctx.JSON(http.StatusOK, gin.H{
			"token": authToken,
		})
		return nil
	}

	ctx.JSON(http.StatusUnauthorized, nil)
	return nil
}

func New(userService userService, jwtService jwtService) (*Controller, error) {
	c := &Controller{
		userService: userService,
		jwtService:  jwtService,
	}

	return c, c.validate()
}

func (c *Controller) validate() error {
	if c.userService == nil {
		return errors.New("service should not be nil")
	}
	if c.jwtService == nil {
		return errors.New("jwt service should not be nil")
	}

	return nil
}
