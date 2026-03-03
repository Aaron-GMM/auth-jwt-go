package config

import (
	"golang.org/x/crypto/bcrypt"
)

var Logger = GetLogger("Security")

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		Logger.ErrorF("bcrypt hash error: %v", err.Error())
		return "", err
	}
	Logger.Info("bcrypt hash check success")

	return string(bytes), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		Logger.ErrorF("bcrypt hash error: %v", err.Error())
		return false
	}
	Logger.Info("bcrypt hash check success")
	return true
}
