package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

var secret = []byte("Mwefkjnkjn234k1@mk&4Jnams")

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization") // "Bearer ertgnoenrg.fdgdfgdf.gdfgdfgdfg"
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, "token was not found")
		c.Abort()
		return
	}
	bearer := strings.Split(authHeader, " ")
	if len(bearer) != 2 || bearer[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, "invalid token")
		c.Abort()
		return
	}

	token, err := verifyToken(bearer[1])
	if err != nil {
		fmt.Printf("Token verification failed: %v\\n", err)
		c.JSON(http.StatusUnauthorized, "invalid token")
		c.Abort()
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		c.Set("user_id", claims["id"])
		c.Set("username", claims["sub"])
	} else {
		fmt.Println(err)
	}
	c.Next()
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}
