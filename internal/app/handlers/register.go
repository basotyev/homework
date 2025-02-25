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

// Register godoc
//
//	@Summary		Create new user
//	@Description	Proceed registration on the service using email, username, password
//	@Tags			users
//	@Security		Bearer
//	@Accept			json
//	@Produce		json
//	@Param			data	body		RegisterRequest	true	"Request input struct"
//	@Success		200		{object}	RegisterResponse
//	@Failure		400		{object}	RegisterResponse
//	@Failure		404		{object}	RegisterResponse
//	@Failure		500		{object}	RegisterResponse
//	@Router			/users/ [post]
func (h *Handler) Register(c *gin.Context) {
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
	c.JSON(200, RegisterResponse{Status: "success", Data: &Data{}})
}
