package handlers

import (
	"github.com/gin-gonic/gin"
	"lesson13/internal/app/usecases"
	"net/http"
)

type Handler struct {
	UserUseCase usecases.UserUseCase
}

func NewHandler() *Handler {
	return &Handler{}
}

func InitRoutes(router *gin.Engine) {
	h := NewHandler()
	inner := router.Group("/users")
	{
		inner.Handle(http.MethodPost, "/", h.Register)
	}
}
