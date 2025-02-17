package crypto

import (
	"github.com/bebrochkas/rural_potatoes/core/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
	})

	t, err := token.SignedString(config.Cfg.JWT_TOKEN)
	if err != nil {
		return "", err
	}

	return t, nil
}

func VerifyToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Cfg.JWT_TOKEN), nil
	})
	if err != nil {
		return false, err
	}

	return token.Valid, nil
}
