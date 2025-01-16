package internal

import (
	"github.com/gin-gonic/gin"
	"lesson13/internal/app"
	db2 "lesson13/internal/app/db"
	"lesson13/internal/app/handlers"
	"log"
	"net/http"
)

func Run() {
	db, err := db2.NewPostgresConnection("postgres://postgres:pass@api-db:5439/test-db?sslmode=disable")
	if err != nil {
		log.Fatalln(err)
		return
	}
	di := app.NewDI(db)

	router := gin.Default()
	handlers.InitRoutes(router, di)
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}
