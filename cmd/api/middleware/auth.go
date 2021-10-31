package middleware

import (
	"fmt"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/service/user"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		service, _ := user.NewJwtService()
		token, err := service.ValidateToken(authHeader)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
