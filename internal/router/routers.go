package router

import (
	"github.com/Aaron-GMM/auth-jwt-go/internal/handler"
	"github.com/gin-gonic/gin"
)

func initializeRouter(router *gin.Engine) {
	authH := handler.InitializeHandler()
	v1 := router.Group("/api/v1")
	{
		v1.POST("/register", authH.Register)

		v1.POST("/login", authH.Login)

	}

}
