package config

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(UserID uint) (string, error) {

	secreatKey := []byte(LoadConfig().JWTKey)
	claims := jwt.MapClaims{
		"user_id": UserID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iat":     time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secreatKey)
	if err != nil {
		Logger.ErrorF("Error generating token: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
