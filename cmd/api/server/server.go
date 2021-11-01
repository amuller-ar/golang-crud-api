package server

import (
	"context"

	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/cmd/api/middleware"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	router := gin.New()

	router.RedirectTrailingSlash = true
	router.RedirectFixedPath = true

	router.Use(func(c *gin.Context) {
		c.Request = c.Request.WithContext(context.Background())
	})

	router.Use(gin.Logger())
	router.Use(middleware.PanicRecovery())
	router.Use(middleware.ErrorHandler)

	mapping := NewMapping()
	mapping.mapURLsToController(router)

	return router
}
