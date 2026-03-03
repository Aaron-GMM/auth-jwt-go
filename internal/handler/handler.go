package handler

import (
	"github.com/Aaron-GMM/auth-jwt-go/internal/handler/authHandler"
	"github.com/Aaron-GMM/auth-jwt-go/internal/infra"
	"github.com/Aaron-GMM/auth-jwt-go/internal/infra/config"
	"github.com/Aaron-GMM/auth-jwt-go/internal/repository"
	"github.com/Aaron-GMM/auth-jwt-go/internal/service"
	"gorm.io/gorm"
)

var (
	Logger *infra.Logger
	Db     *gorm.DB
)

func InitializeHandler() *authHandler.AuthHandler {
	Logger = config.GetLogger("handler")
	Db = config.GetDb()

	repo := repository.NewAuthRepository(Db)
	authService := service.NewAuthService(repo)
	return authHandler.NewAuthHandler(authService)
}
