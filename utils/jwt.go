package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/kgermando/lobi/config"
)

var SECRET_KEY string = config.ConfigEnv("SECRET")

func GenerateJwt(issuer string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(12 * time.Hour)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    issuer,
		ExpiresAt: jwt.NewNumericDate(expireTime),
	})

	token, err := claims.SignedString([]byte(SECRET_KEY))

	return token, err
}

func VerifyJWT(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil || !token.Valid {
		return "", err
	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	return claims.Issuer, nil
}
