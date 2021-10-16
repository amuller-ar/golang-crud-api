package middleware

import "github.com/gin-gonic/gin"

// AdaptHandler adapta func(*gin.Context) error  a func(*gin.Context)
// tambien aborta el gin.Context actual cuanto el handler retorna error
func AdaptHandler(handler func(*gin.Context) error) func(*gin.Context) {
	return func(c *gin.Context) {
		if err := handler(c); err != nil {
			_ = c.Error(err)
			c.Abort()

		}
	}
}
