package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Service interface {
	CreateToken(id int, username string) (string, error)
	CreateRefreshToken(id int) (string, error)
}

type service struct {
	secret []byte
}

func NewService(secret []byte) Service {
	return &service{secret: secret}
}

func (s *service) CreateToken(id int, username string) (string, error) {
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
	return tokenString, nil
}

func (s *service) CreateRefreshToken(id int) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour).Unix(),
	})
	tokenString, err := claims.SignedString(s.secret)
	if err != nil {
		return "", err
	}
	fmt.Println(claims)
	fmt.Printf("Token: %s\n", tokenString)
	return tokenString, nil
}
