package middleware

import (
	"github.com/gin-gonic/gin"
)

// AdaptHandler adapts func(*gin.Context) error a func(*gin.Context)
// also aborts  gin.Context when the handler return an error
func AdaptHandler(handler func(*gin.Context) error) func(*gin.Context) {
	return func(c *gin.Context) {
		if err := handler(c); err != nil {
			_ = c.Error(err)
			c.Abort()
		}
	}
}
