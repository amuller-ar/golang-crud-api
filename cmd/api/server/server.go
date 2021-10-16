package server

import (
	"context"

	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/cmd/api/middleware"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.RedirectTrailingSlash = true
	router.RedirectFixedPath = true

	router.Use(func(c *gin.Context) {
		c.Request = c.Request.WithContext(context.Background())
	})

	router.Use(middleware.PanicRecovery())

	mapping := NewMapping()
	mapping.mapURLsToController(router)

	return router
}
