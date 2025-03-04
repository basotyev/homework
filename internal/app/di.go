package app

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"lesson13/configs"
	db2 "lesson13/internal/app/db"
	authS "lesson13/internal/app/services/auth"
	authR "lesson13/internal/app/services/auth/repository"
	userS "lesson13/internal/app/services/user"
	userR "lesson13/internal/app/services/user/repository"
	"lesson13/internal/app/usecases"
	"log"
)

type Services struct {
	Auth authS.Service
}

type DI struct {
	Config      *configs.Config
	UserUseCase usecases.UserUseCase
	Services    *Services
}

func NewDI(config *configs.Config) *DI {
	db, err := db2.NewPostgresConnection(config)
	if err != nil {
		log.Fatalln(err)
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		DB:       0,
		Password: "",
	})
	if _, err := redisClient.Ping(context.Background()).Result(); err != nil {
		log.Fatalln(err)
	}

	var secret = []byte(config.Token.AccessSecret)
	var refreshSecret = []byte(config.Token.RefreshSecret)
	userRepo := userR.NewRepository(db)
	authRepo := authR.New(redisClient, config.Token.AccessExpire, config.Token.RefreshExpire)

	userService := userS.NewService(userRepo)
	authService := authS.NewService(secret, refreshSecret, authRepo)

	userUseCase := usecases.NewUserUseCase(userService, authService)
	return &DI{
		UserUseCase: userUseCase,
		Config:      config,
		Services: &Services{
			Auth: authService,
		},
	}
}
