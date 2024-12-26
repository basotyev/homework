package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"lesson13/internal/app/utils"
	"net/http"
)

type PostLoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) postLogin(c *gin.Context) {
	var req PostLoginReq
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    1,
			"message": "invalid email/password",
		})
		return
	}

	token, err := h.DI.UserUseCase.Login(c, req.Email, req.Password)
	if err != nil {
		if errors.Is(err, utils.ErrInternalServerError) {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    2,
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken": token,
	})
}
