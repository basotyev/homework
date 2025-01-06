package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostRefreshReq struct {
	Refresh string `json:"refresh"`
}

type PostRefreshResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (h *Handler) postRefresh(c *gin.Context) {
	var req PostRefreshReq
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    1,
			"message": "invalid refresh token",
		})
	}
	tok, err := h.DI.UserUseCase.Refresh(c, req.Refresh)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    2,
			"message": err.Error(),
		})
	}
	res := PostRefreshResponse{
		AccessToken:  tok.AccessToken,
		RefreshToken: tok.RefreshToken,
	}
	c.JSON(http.StatusOK, res)
}
