package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"lesson13/internal/app"
	"net/http"
	"strings"
)

func AuthMiddleware(c *gin.Context, di *app.DI) {
	authHeader := c.GetHeader("Authorization") // "Bearer ertgnoenrg.fdgdfgdf.gdfgdfgdfg"
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, map[string]string{"error": "token was not found"})
		c.Abort()
		return
	}
	bearer := strings.Split(authHeader, " ")
	if len(bearer) != 2 || bearer[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
		c.Abort()
		return
	}

	token, err := di.Services.Auth.VerifyTokenAccessToken(c.Request.Context(), bearer[1])
	if err != nil {
		fmt.Printf("Token verification failed: %v\\n", err)
		c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
		c.Abort()
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		c.Set("user_id", claims["id"])
		c.Set("username", claims["sub"])
	} else {
		fmt.Println(err)
		c.Abort()
	}
	c.Next()
}
