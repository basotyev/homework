package usecases

import (
	"context"
	"errors"
	"lesson13/internal/app/models"
	"lesson13/internal/app/services/auth"
	"lesson13/internal/app/services/user"
	userR "lesson13/internal/app/services/user/repository"
	"testing"
)

type mocks struct {
	userRepoMock userR.RepositoryMock
}

func defaultMocks() *mocks {
	return &mocks{
		userRepoMock: userR.RepositoryMock{},
	}
}

func CreateUserUC(mocks *mocks) UserUseCase {
	userS := user.NewService(&mocks.userRepoMock)
	authS := auth.NewService([]byte("secret"), []byte("secret"))

	return &userUseCase{userService: userS, authService: authS}
}

var InvalidInput = errors.New("invalid input")
var InvalidResult = errors.New("invalid result")

func TestRegisterEmail(t *testing.T) {
	mocks := defaultMocks()
	mocks.userRepoMock.CreateUserFunc = func(ctx context.Context, user *models.User) error {
		if user.Email != "TEST" {
			return InvalidInput
		}
		return nil
	}
	uc := CreateUserUC(mocks)
	err := uc.Register(context.Background(), "test", "TEST", "TEST")
	if err != nil {
		t.Error(err)
		t.Error(InvalidResult)
	}
}

func TestRegisterName(t *testing.T) {
	mocks := defaultMocks()
	mocks.userRepoMock.CreateUserFunc = func(ctx context.Context, user *models.User) error {
		if user.Name != "test" {
			return InvalidInput
		}
		return nil
	}
	uc := CreateUserUC(mocks)
	err := uc.Register(context.Background(), "test", "TEST", "TEST")
	if err != nil {
		t.Error(err)
		t.Error(InvalidResult)
	}
}
