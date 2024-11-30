package user

import (
	"context"
	"lesson13/models"
)

type Repository interface {
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUserById(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id int) (*models.User, error)
	RemoveUserById(ctx context.Context, id int) error
}
