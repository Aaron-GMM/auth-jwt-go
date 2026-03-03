package router

import (
	"github.com/Aaron-GMM/auth-jwt-go/internal/handler/authHandler"
	"github.com/Aaron-GMM/auth-jwt-go/internal/repository"
	"github.com/Aaron-GMM/auth-jwt-go/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initializeRouter(router *gin.Engine, db *gorm.DB) {
	repo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(repo)
	authH := authHandler.NewAuthHandler(authService)

	v1 := router.Group("/api/v1")
	{
		v1.POST("/register", authH.Register)

		v1.POST("/login", authH.Login)

	}

}
