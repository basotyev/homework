package handlers

import (
	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Status string `json:"status"`
	Data   *Data  `json:"data,omitempty"`
}

type Data struct {
}

func (h *Handler) Register(c *gin.Context) {
	getUserId(c)
	var req RegisterRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, RegisterResponse{Status: "invalid input"})
		return
	}
	err = h.DI.UserUseCase.Register(c, req.Username, req.Email, req.Password)
	if err != nil {
		c.JSON(500, RegisterResponse{Status: "invalid input"})
		return
	}
	c.JSON(200, gin.H{})
}
