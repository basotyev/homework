package internal

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"lesson13/configs"
	"lesson13/internal/app"
	"lesson13/internal/app/handlers"
	"log"
	"net/http"
)

func Run() {
	var path string
	flag.StringVar(&path, "config", "./configs/config.yml", "для конфигурации")
	flag.Parse()
	config, err := configs.NewConfig(path)
	if err != nil {
		log.Fatalln(err)
	}
	di := app.NewDI(config)

	router := gin.Default()
	handlers.InitRoutes(router, di)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.App.Port),
		Handler: router,
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
