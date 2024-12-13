package usecases

import (
	"context"
	"lesson13/internal/app/models"
	"lesson13/internal/app/services/user"
)

type UserUseCase interface {
	Register(ctx context.Context, username string, password string) error
}

type userUseCase struct {
	userService user.Service
}

func NewUserUseCase(userService user.Service) UserUseCase {
	return &userUseCase{userService: userService}
}

func (u *userUseCase) Register(ctx context.Context, username string, email string) error {
	usr := &models.User{
		Name:  username,
		Email: email,
		Age:   0,
	}
	err := u.userService.CreateUser(ctx, usr)
	if err != nil {
		return err
	}
	return nil
}
