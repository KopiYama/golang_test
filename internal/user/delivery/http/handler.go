package http

import (
	"github.com/gin-gonic/gin"
	userUsecase "golang_test/internal/user/usecase"
	"net/http"
)

type UserHandler struct {
	usecase userUsecase.UserUsecase
}

func NewUserHandler(u userUsecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: u}
}

func (h *UserHandler) GetAll(c *gin.Context) {
	users, err := h.usecase.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Successfully",
		"data": gin.H{
			"user": users,
		},
	})
}

func (h *UserHandler) Create(c *gin.Context) {
	var req struct {
		RoleID   int    `json:"role_id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "invalid request"})
		return
	}

	err := h.usecase.Create(c.Request.Context(), req.RoleID, req.Name, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Successfully"})
}

func (h *UserHandler) Update(c *gin.Context) {
	var req struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "invalid request"})
		return
	}

	// contoh: ambil user ID dari token atau session jika sudah ada
	userID := 1 // hardcoded sementara

	err := h.usecase.Update(c.Request.Context(), userID, req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Successfully"})
}

func (h *UserHandler) Delete(c *gin.Context) {
	userID := c.Param("user_id")
	err := h.usecase.Delete(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Successfully"})
}
