package repository

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type Repository interface {
	SaveAccessToken(ctx context.Context, id int, access string) error
	SaveRefreshToken(ctx context.Context, id int, refresh string) error
	GetAccessToken(ctx context.Context, id int) (string, error)
	GetRefreshToken(ctx context.Context, id int) (string, error)
}

type repository struct {
	redis             *redis.Client
	accessExpiration  time.Duration
	refreshExpiration time.Duration
}

func New(redis *redis.Client, accessExp, refreshExp int) Repository {
	return &repository{
		redis:             redis,
		accessExpiration:  time.Duration(accessExp) * time.Minute,
		refreshExpiration: time.Duration(refreshExp) * time.Minute,
	}
}

func (r *repository) SaveAccessToken(ctx context.Context, id int, access string) error {
	key := fmt.Sprintf("access:%d", id)
	return r.redis.Set(ctx, key, access, r.accessExpiration).Err()
}

func (r *repository) SaveRefreshToken(ctx context.Context, id int, refresh string) error {
	key := fmt.Sprintf("refresh:%d", id)
	return r.redis.Set(ctx, key, refresh, r.refreshExpiration).Err()
}

func (r *repository) GetAccessToken(ctx context.Context, id int) (string, error) {
	key := fmt.Sprintf("access:%d", id)
	return r.redis.Get(ctx, key).Result()
}

func (r *repository) GetRefreshToken(ctx context.Context, id int) (string, error) {
	key := fmt.Sprintf("refresh:%d", id)
	return r.redis.Get(ctx, key).Result()
}
