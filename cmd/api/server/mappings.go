package server

import (
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/cmd/api/middleware"
	propertyController "github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/controller/property"
	testController "github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/controller/testing"
	"github.com/gin-gonic/gin"
)

type mapping struct {
	testController     *testController.Controller
	propertyController *propertyController.Controller
}

func NewMapping() *mapping {
	//add dependency below
	return &mapping{
		testController:     resolveTestController(),
		propertyController: resolvePropertyController(),
	}
}

func (m mapping) mapURLsToController(router *gin.Engine) {
	//URL's can be grouped here

	router.GET("/test", middleware.AdaptHandler(m.testController.GET))

	baseGroup := router.Group("/v1")
	{
		propertyGroup := baseGroup.Group("/properties")
		{
			propertyGroup.POST("", middleware.AdaptHandler(m.propertyController.Create))
		}

	}

}
