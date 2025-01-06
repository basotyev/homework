package app

import (
	"github.com/jackc/pgx/v4/pgxpool"
	authS "lesson13/internal/app/services/auth"
	userS "lesson13/internal/app/services/user"
	userR "lesson13/internal/app/services/user/repository"
	"lesson13/internal/app/usecases"
)

type DI struct {
	UserUseCase usecases.UserUseCase
}

func NewDI(db *pgxpool.Pool) *DI {
	var secret = []byte("Mwefkjnkjn234k1@mk&4Jnams")
	var refreshSecret = []byte("Nmsjdnfkjn234123#KJNKJN")
	userRepo := userR.NewRepository(db)

	userService := userS.NewService(userRepo)
	authService := authS.NewService(secret, refreshSecret)

	userUseCase := usecases.NewUserUseCase(userService, authService)
	return &DI{
		UserUseCase: userUseCase,
	}
}
