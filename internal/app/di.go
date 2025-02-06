package app

import (
	"lesson13/configs"
	db2 "lesson13/internal/app/db"
	authS "lesson13/internal/app/services/auth"
	userS "lesson13/internal/app/services/user"
	userR "lesson13/internal/app/services/user/repository"
	"lesson13/internal/app/usecases"
	"log"
)

type DI struct {
	UserUseCase usecases.UserUseCase
}

func NewDI(config *configs.Config) *DI {
	db, err := db2.NewPostgresConnection(config)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	var secret = []byte(config.Token.AccessSecret)
	var refreshSecret = []byte(config.Token.RefreshSecret)
	userRepo := userR.NewRepository(db)

	userService := userS.NewService(userRepo)
	authService := authS.NewService(secret, refreshSecret)

	userUseCase := usecases.NewUserUseCase(userService, authService)
	return &DI{
		UserUseCase: userUseCase,
	}
}
