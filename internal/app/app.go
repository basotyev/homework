package app

import (
	"github.com/gin-gonic/gin"
	"lesson13/internal/app/handlers"
	"net/http"
)

func Run() {
	router := gin.Default()
	handlers.InitRoutes(router)
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}
