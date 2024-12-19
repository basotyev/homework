package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	db2 "lesson13/internal/app/db"
	"lesson13/internal/app/handlers"
	userS "lesson13/internal/app/services/user"
	userR "lesson13/internal/app/services/user/repository"
	"lesson13/internal/app/usecases"
	"log"
	"net/http"
)

func Run() {
	db, err := db2.NewPostgresConnection("postgres://postgres:pass@localhost:5432/test-db?sslmode=disabled")
	if err != nil {
		log.Fatalln(err)
		return
	}
	di := NewDI(db)

	router := gin.Default()
	handlers.InitRoutes(router, di)
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}

type DI struct {
	UserUseCase usecases.UserUseCase
}

func NewDI(db *pgxpool.Pool) *DI {
	userRepo := userR.NewRepository(db)
	userService := userS.NewService(userRepo)
	userUseCase := usecases.NewUserUseCase(userService)

	return &DI{
		UserUseCase: userUseCase,
	}
}
