package auth

import (
	"errors"
	"fmt"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/infrastructure/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"strings"
)

func ExtractTokenMetadata(ctx *gin.Context) (*string, error) {
	token, err := VerifyToken(ctx)
	if err != nil {
		return nil, err
	}
	account, err := Extract(token)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func VerifyToken(ctx *gin.Context) (*jwt.Token, error) {
	tokenString := ExtractToken(ctx)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(utils.GetSecretKey()), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractToken(ctx *gin.Context) string {
	bearToken := ctx.GetHeader("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return strArr[0]
}

func Extract(token *jwt.Token) (string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {

		userName, userOk := claims["name"].(string)

		if ok == false || userOk == false {
			return "", errors.New("unauthorized")
		} else {
			return userName, nil
		}
	}
	return "", errors.New("something went wrong")
}
