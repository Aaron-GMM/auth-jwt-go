package authHandler

import (
	"net/http"

	"github.com/Aaron-GMM/auth-jwt-go/internal/handler/requestDTOs"
	"github.com/Aaron-GMM/auth-jwt-go/internal/handler/responseDTOs"
	"github.com/Aaron-GMM/auth-jwt-go/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (handler *AuthHandler) Login(c *gin.Context) {
	var request requestDTOs.LoginRequestDTO

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest,
			responseDTOs.ErrorResponse{Error: "Invalid Data:" + err.Error()})
		return
	}

	token, err := handler.service.Login(request)

	if err != nil {
		c.JSON(http.StatusUnauthorized, responseDTOs.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, responseDTOs.LoginResponse{Token: token})

}
