package server

import (
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/cmd/api/middleware"
	testController "github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/controller/testing"
	"github.com/gin-gonic/gin"
)

type mapping struct {
	//controllers aca
	testController *testController.Controller
}

func NewMapping() *mapping {
	//agregar resolucion de dependencias debajo
	return &mapping{
		testController: resolveTestController(),
	}
}

func (m mapping) mapURLsToController(router *gin.Engine) {
	//se puede agrupar urls aqui

	router.GET("/test", middleware.AdaptHandler(m.testController.GET))
}
