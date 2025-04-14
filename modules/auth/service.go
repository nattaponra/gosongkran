package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-secret-key") // Should load from env in real use

func AuthenticateUser(username, password string) (string, error) {
	// Replace with DB logic
	if username == "admin" && password == "pass123" {
		return GenerateJWT(username)
	}
	return "", errors.New("invalid credentials")
}

func CreateUser(req RegisterRequest) error {
	// You'd usually insert into a DB here
	return nil
}

func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
