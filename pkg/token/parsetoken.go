package token

import (
	"miniproject/config"

	"github.com/dgrijalva/jwt-go"
)

func ParseToken(tokenstring string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenstring, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Key), nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
