package handlers

import (
	"github.com/gin-gonic/gin"
	"lesson13/internal/app"
	"net/http"
)

type Handler struct {
	DI *app.DI
}

func InitRoutes(router *gin.Engine, di *app.DI) {
	h := Handler{
		DI: di,
	}

	inner := router.Group("/users")
	{
		inner.Handle(http.MethodPost, "/", h.Register)
	}
}
