package model

import (
	"nothing/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

type UserJwt struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte(config.Config.Jwt.SigningKey)

const (
	tokenExpireDuration = time.Hour * 24 * 30
)

func GenerateToken(username string) (string, error) {
	claims := UserJwt{
		username,
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),                          // 生成时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpireDuration)), // 过期时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseToken(tokenString string) (*UserJwt, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &UserJwt{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn't handle this token")
			}
		}
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*UserJwt); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, errors.New("couldn't handle this token")
}
