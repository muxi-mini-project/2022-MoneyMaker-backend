package token

import (
	"miniproject/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	ID string
	jwt.StandardClaims
}

func GenerateToken(id string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(7 * 24 * time.Hour)
	claims := Claims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    config.Issuer,
			NotBefore: time.Now().Unix(), // 生效时间
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.Key))
	return token, err
}
