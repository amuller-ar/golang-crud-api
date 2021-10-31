package server

import (
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/cmd/api/middleware"
	propertyController "github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/controller/property"
	userController "github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/controller/user"
	"github.com/gin-gonic/gin"
)

type mapping struct {
	propertyController *propertyController.Controller
	userController     *userController.Controller
}

func NewMapping() *mapping {
	//add dependency below
	return &mapping{
		propertyController: resolvePropertyController(),
		userController:     resolveUserController(),
	}
}

func (m mapping) mapURLsToController(router *gin.Engine) {
	baseGroup := router.Group("/v1")
	{
		propertyGroup := baseGroup.Group("/properties")
		{
			propertyGroup.POST("", middleware.AdaptHandler(m.propertyController.Create))
			propertyGroup.GET("", middleware.AdaptHandler(m.propertyController.Search))
			propertyGroup.PUT("/:id", middleware.AdaptHandler(m.propertyController.Update))

		}

		userGroup := baseGroup.Group("/users")
		{
			userGroup.POST("", middleware.AdaptHandler(m.userController.Create))
			userGroup.POST("/login", middleware.AdaptHandler(m.userController.Login))

			authGroup := userGroup.Group("/me")
			{
				authGroup.Use(middleware.AuthorizeJWT())
				authGroup.POST("/favorites")
			}
		}
	}
}
