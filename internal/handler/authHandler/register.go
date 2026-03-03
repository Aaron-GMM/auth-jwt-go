package authHandler

import (
	"net/http"

	"github.com/Aaron-GMM/auth-jwt-go/internal/handler/requestDTOs"
	"github.com/Aaron-GMM/auth-jwt-go/internal/handler/responseDTOs"
	"github.com/gin-gonic/gin"
)

func (h *AuthHandler) Register(c *gin.Context) {
	var req requestDTOs.RegisterRequestDTO
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := h.service.Register(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responseDTOs.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, responseDTOs.DefaultResponse{Message: "success", Data: token})
}
