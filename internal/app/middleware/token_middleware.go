package middleware

import (
	"log"
	"os"

	"github.com/gofiber/contrib/fiberi18n/v2"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func TokenMiddleware() func(*fiber.Ctx) error {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("Please set JWT_SECRET in your environment")
	}
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(secret)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			localize, _ := fiberi18n.Localize(c, "auth.not_authenticated")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": localize,
			})
		},
	})
}
