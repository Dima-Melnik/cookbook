package auth

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func GetENV() (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", errors.New("error loading env")
	}

	secret := os.Getenv("JWT_SECRET")
	if secretStr == "" {
		return "", errors.New("invalid value from .env")
	}

	return secret, nil
}
