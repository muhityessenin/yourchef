package user

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("ieqriebqFNEIvbv9ewvnon3u543v34248jnveibviewpvb")

func GenerateToken(user *Entity) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.Id,
		"role": user.Role,
		"exp":  time.Now().Add(72 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
