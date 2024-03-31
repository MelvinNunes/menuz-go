package security

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWTtoken(userID string) (*string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"id":    userID,
		"roles": true,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("Please set JWT_SECRET in your environment")
	}
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func GenerateApiKeyJWTtoken(userID string) (string, string) {
	// Create the Claims
	claims := jwt.MapClaims{
		"id":    userID,
		"roles": true,
		"exp":   time.Now().AddDate(1, 0, 0).Unix(),
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("Please set JWT_SECRET in your environment")
	}

	rawToken, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Panic("Error signing the JWT token")
	}

	hashedToken := hashToken(rawToken)
	if err != nil {
		log.Panic("Error generating the JWT token: ", err)
	}

	return rawToken, string(hashedToken)
}

func GetOnlineUserID(c *fiber.Ctx) string {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["id"].(string)
}
