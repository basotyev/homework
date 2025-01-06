package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Service interface {
	CreateToken(id int, username string) (string, error)
	CreateRefreshToken(id int) (string, error)
	ValidateRefreshToken(refresh string) (int, error)
}

type service struct {
	secret        []byte
	refreshSecret []byte
}

func NewService(secret, refresh []byte) Service {
	return &service{
		secret:        secret,
		refreshSecret: refresh,
	}
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
	tokenString, err := claims.SignedString(s.refreshSecret)
	if err != nil {
		return "", err
	}
	fmt.Println(claims)
	fmt.Printf("Token: %s\n", tokenString)
	return tokenString, nil
}

func (s *service) ValidateRefreshToken(refresh string) (int, error) {
	token, err := jwt.Parse(refresh, func(token *jwt.Token) (interface{}, error) {
		return s.refreshSecret, nil
	})
	if err != nil {
		return 0, err
	}
	if !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}
	var id int
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		idFloat := claims["id"].(float64)
		id = int(idFloat)
		if err != nil {
			return 0, err
		}
		return id, nil
	}
	return 0, err
}
