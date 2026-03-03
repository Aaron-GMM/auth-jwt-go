package config

import (
	"github.com/Aaron-GMM/auth-jwt-go/internal/handler"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	Logger := handler.Logger
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		Logger.ErrorF("bcrypt hash error: %v", err.Error())
		return "", err
	}
	Logger.Info("bcrypt hash check success")

	return string(bytes), nil
}

func CheckPasswordHash(password, hash string) bool {
	Logger := handler.Logger
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		Logger.ErrorF("bcrypt hash error: %v", err.Error())
		return false
	}
	Logger.Info("bcrypt hash check success")
	return true
}
