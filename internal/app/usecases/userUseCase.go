package usecases

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"lesson13/internal/app/models"
	"lesson13/internal/app/services/auth"
	"lesson13/internal/app/services/user"
	"lesson13/internal/app/utils"
)

type UserUseCase interface {
	Register(ctx context.Context, username string, email string, password string) error
	Login(ctx context.Context, email, password string) (*models.Token, error)
	Refresh(ctx context.Context, refresh string) (*models.Token, error)
}

type userUseCase struct {
	userService user.Service
	authService auth.Service
}

func (u *userUseCase) Login(ctx context.Context, email, password string) (*models.Token, error) {
	usr, err := u.userService.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, utils.ErrUserNotFound
		}
		return nil, utils.ErrInternalServerError
	}
	if password != usr.Password {
		return nil, utils.ErrIncorrectPassword
	}
	token, err := u.authService.CreateToken(usr.Id, usr.Name)
	if err != nil {
		return nil, utils.ErrInternalServerError
	}
	refresh, err := u.authService.CreateRefreshToken(usr.Id)
	if err != nil {
		return nil, utils.ErrInternalServerError
	}
	return &models.Token{
		AccessToken:  token,
		RefreshToken: refresh,
	}, nil
}

func (u *userUseCase) Refresh(ctx context.Context, refresh string) (*models.Token, error) {

}

func NewUserUseCase(userService user.Service, authService auth.Service) UserUseCase {
	return &userUseCase{
		userService: userService,
		authService: authService,
	}
}

func (u *userUseCase) Register(ctx context.Context, username, email string, password string) error {
	usr := &models.User{
		Name:     username,
		Email:    email,
		Age:      0,
		Password: password,
	}
	err := u.userService.CreateUser(ctx, usr)
	if err != nil {
		return err
	}
	return nil
}
