package http

import (
	"github.com/gin-gonic/gin"
	"golang_test/internal/user/usecase"
	"net/http"

	authUsecase "golang_test/internal/auth/usecase"
)

type AuthHandler struct {
	usecase authUsecase.AuthUsecase
}

func NewAuthHandler(u authUsecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{usecase: u}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "invalid request"})
		return
	}

	token, err := h.usecase.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Successfully",
		"data": gin.H{
			"access_token": token,
		},
	})
}

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(u usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: u}
}

func (h *AuthHandler) Logout(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "unauthorized"})
		return
	}

	err := h.usecase.Logout(c.Request.Context(), userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "failed to logout"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "logged out successfully"})
}
