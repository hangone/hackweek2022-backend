package model

import (
	"log"
	"nothing/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

type UserJwt struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var (
	//jwtSecret           = viper.GetString("jwt.signingKey")
	tokenExpireDuration = time.Hour * 24 * 30 // 一个月
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
	return token.SignedString([]byte(config.Config.Jwt.SigningKey))
}

func ParseToken(tokenString string) (*UserJwt, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &UserJwt{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.Jwt.SigningKey), nil
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*UserJwt); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, errors.New("未知的 Token")
}
