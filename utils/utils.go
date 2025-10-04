package utils

import (
	"crypto/rand"
	"encoding/base64"
	"log/slog"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err.Error()
	}
	return string(hash)
}

func CompareHashedPassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func GenerateToken(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		slog.Error("Error in generating session token")
	}
	return base64.URLEncoding.EncodeToString(bytes)
}
