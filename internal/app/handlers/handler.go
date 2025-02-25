package handlers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "lesson13/docs"
	"lesson13/internal/app"
	"lesson13/internal/app/middleware"
	"lesson13/internal/app/models"
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
	if di.Config.App.Env != models.EnvProduction {
		inner.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	users := inner.Group("/users")
	{
		users.POST("/", h.Register)
		users.Use(middleware.AuthMiddleware)
		users.GET("/healthcheck", h.test)
	}
	inner.POST("/login", h.postLogin) // /api/v1/login
	inner.POST("/refresh", h.postRefresh)

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
