package auth

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"lesson13/internal/app/services/auth/repository"
	"time"
)

type Service interface {
	CreateToken(ctx context.Context, id int, username string) (string, error)
	CreateRefreshToken(ctx context.Context, id int) (string, error)
	ValidateRefreshToken(ctx context.Context, refresh string) (int, error)
	VerifyTokenAccessToken(ctx context.Context, tokenString string) (*jwt.Token, error)
}

type service struct {
	secret        []byte
	refreshSecret []byte
	repo          repository.Repository
}

func NewService(secret, refresh []byte, repo repository.Repository) Service {
	return &service{
		secret:        secret,
		refreshSecret: refresh,
		repo:          repo,
	}
}

func (s *service) CreateToken(ctx context.Context, id int, username string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"sub": username,
		"iss": "task-manager-app",
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})
	tokenString, err := claims.SignedString(s.secret)
	if err != nil {
		return "", err
	}
	fmt.Println(claims)
	fmt.Printf("Token: %s\n", tokenString)
	err = s.repo.SaveAccessToken(ctx, id, tokenString)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *service) CreateRefreshToken(ctx context.Context, id int) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour).Unix(),
	})
	tokenString, err := claims.SignedString(s.refreshSecret)
	if err != nil {
		return "", err
	}
	fmt.Println(claims)
	fmt.Printf("Token: %s\n", tokenString)
	err = s.repo.SaveRefreshToken(ctx, id, tokenString)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *service) ValidateRefreshToken(ctx context.Context, refresh string) (int, error) {
	token, err := jwt.Parse(refresh, func(token *jwt.Token) (interface{}, error) {
		return s.refreshSecret, nil
	})
	if err != nil {
		return 0, err
	}
	if !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		id := claims["id"].(int)
		oldRefresh, err := s.repo.GetRefreshToken(ctx, id)
		if err != nil {
			return 0, fmt.Errorf("invalid token")
		}
		if oldRefresh != refresh {
			return 0, fmt.Errorf("invalid token")
		}
		return id, nil
	}

	return 0, fmt.Errorf("invalid token")
}

func (s *service) VerifyTokenAccessToken(ctx context.Context, tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return s.secret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		oldAccess, err := s.repo.GetAccessToken(ctx, claims["id"].(int))
		if err != nil {
			return nil, fmt.Errorf("invalid token")
		}
		if oldAccess != tokenString {
			return nil, fmt.Errorf("invalid token")
		}
	} else {
		fmt.Println(err)
	}
	return token, nil
}
