package handlers

import (
	"net/http"
	"social-sys/internal/api/dto"
	"social-sys/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service service.AuthServiceInterface
}

func NewAuthHandler(service service.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var loginInput dto.LoginInput

	if err := ctx.ShouldBindJSON(&loginInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	accessToken := h.service.Login(&loginInput)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successfully",
		"AccessToken": accessToken,
	})
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var authInput dto.CreateUserInput

	if err := ctx.ShouldBindJSON(&authInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := h.service.Register(&authInput); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Registger account successfully",
	})
}
