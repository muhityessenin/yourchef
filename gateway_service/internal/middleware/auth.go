package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Claims struct {
	UserID string `json:"id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
		fmt.Println(authHeader)
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		if time.Unix(claims.ExpiresAt, 0).Before(time.Now()) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
			return
		}

		c.Set("id", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()
	}
}
