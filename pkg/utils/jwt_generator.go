package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// Creates JWT Token with user's credentials
func JWTGenerator(username string) (string, error) {
	exp := time.Now().Add(time.Minute * 30).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = exp

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return t, nil
}
