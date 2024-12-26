package user

import (
	"context"
	"lesson13/internal/app/models"
	"lesson13/internal/app/services/user/repository"
)

type Service interface {
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUserById(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id int) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	RemoveUserById(ctx context.Context, id int) error
}

type service struct {
	userRepo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{
		userRepo: repo,
	}
}

func (s *service) CreateUser(ctx context.Context, user *models.User) error {
	return s.userRepo.CreateUser(ctx, user)
}

func (s *service) UpdateUserById(ctx context.Context, user *models.User) error {
	return s.userRepo.UpdateUserById(ctx, user)
}

func (s *service) GetUserById(ctx context.Context, id int) (*models.User, error) {
	return s.userRepo.GetUserById(ctx, id)
}

func (s *service) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.userRepo.GetUserByEmail(ctx, email)
}

func (s *service) RemoveUserById(ctx context.Context, id int) error {
	return s.userRepo.RemoveUserById(ctx, id)
}
