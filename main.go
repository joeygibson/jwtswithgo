package main

import (
	"github.com/dgrijalva/jwt-go"
	logger "github.com/sirupsen/logrus"
	"time"
)

var secret = "testing"

func createToken(message string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"msg": message,
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func main() {
	message := "this is a test of JWT in Go"

	logger.Info("Creating token")

	token, err := createToken(message)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Printf("Token: %s\n", token)
}
