package handlers

import (
	"github.com/gin-gonic/gin"
	"lesson13/internal/app"
	"lesson13/internal/app/middleware"
	"net/http"
	"strconv"
)

type Handler struct {
	DI *app.DI
}

func InitRoutes(router *gin.Engine, di *app.DI) {
	h := Handler{
		DI: di,
	}
	inner := router.Group("/api/v1")
	users := inner.Group("/users")
	{
		users.Handle(http.MethodPost, "/", h.Register)
		users.Use(middleware.AuthMiddleware)
		users.GET("/healthcheck", h.test)
	}
	inner.POST("/login", h.postLogin) // /api/v1/login
	inner.POST("/refresh", h.)

}

func (h *Handler) test(c *gin.Context) {
	c.JSON(200, nil)
}

func getUserId(c *gin.Context) int {
	val, ok := c.Get("user_id")
	if ok {
		id, err := strconv.Atoi(val.(string))
		if err != nil {
			return 0
		}
		return id
	}
	return 0
}

func getUsername(c *gin.Context) string {
	val, ok := c.Get("username")
	if !ok {
		return ""
	}
	return val.(string)
}
