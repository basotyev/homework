package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostRefreshReq struct {
	Refresh string `json:"refresh"`
}

type PostRefreshResponse struct {
	AccessToken string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func(h *Handler) postRefresh(c *gin.Context) {
	var req PostLoginReq
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{} {
			"code":1,
			"message":"invalid refresh token",
		})
	}
	h.DI.UserUseCase.
}
