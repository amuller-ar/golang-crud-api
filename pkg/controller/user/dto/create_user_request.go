package dto

import (
	"fmt"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/domain"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/infrastructure/validator"
	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Email string `json:"email" validate:"required,email"`
}

func ParseCreateUserRequest(ctx *gin.Context) (*domain.User, error) {
	var request CreateUserRequest

	if err := ctx.BindJSON(&request); err != nil {
		return nil, fmt.Errorf("error binding json request. cause: %v", err)
	}

	if err := validator.Validate(request); err != nil {
		return nil, err
	}

	user := request.ToDomainUser()

	return &user, nil
}

func (r CreateUserRequest) ToDomainUser() domain.User {
	return domain.User{
		Email:    r.Email,
		Password: r.Email,
	}
}
