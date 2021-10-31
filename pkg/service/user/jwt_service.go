package user

import (
	"fmt"
	"github.com/alan-muller-ar/alan-muller-ar-lahaus-backend/pkg/infrastructure/utils"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type authCustomClaims struct {
	Name string `json:"name"`
	User bool   `json:"user"`
	jwt.RegisteredClaims
}

type JwtService struct {
	secretKey string
	issure    string
}

func NewJwtService() (*JwtService, error) {
	secretKey := utils.GetSecretKey()

	return &JwtService{
		secretKey: secretKey,
		issure:    "Bikash",
	}, nil
}

func (s *JwtService) GenerateToken(email string, isUser bool) string {
	expiresAt := &jwt.NumericDate{
		Time: time.Now().Add(time.Hour * 48),
	}

	issuedAt := &jwt.NumericDate{
		Time: time.Now(),
	}

	claims := &authCustomClaims{
		email,
		isUser,
		jwt.RegisteredClaims{
			ExpiresAt: expiresAt,
			Issuer:    s.issure,
			IssuedAt:  issuedAt,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (s *JwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token : %v", token.Header["alg"])

		}
		return []byte(s.secretKey), nil
	})

}
