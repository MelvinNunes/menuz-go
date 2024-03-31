package security

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func hashToken(token string) string {
	data := []byte(token)
	// Combine the salt and data
	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash)
}

func IsPasswordEqualToHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil // True if password matches hash
}

func GenerateAPIKey() (string, string) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	apiKey := hex.EncodeToString(bytes)
	key, err := bcrypt.GenerateFromPassword([]byte(apiKey), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return apiKey, string(key)
}

func IsApiKeyValid(rawKey string, hashedApiKey string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedApiKey), []byte(rawKey))
	return err == nil
}

func IsApiTokenKeyValid(rawKey string, hashedApiKey string) bool {
	return hashToken(rawKey) == hashedApiKey
}
