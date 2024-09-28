package errorchecker

import (
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
)

func ResolveError(c *fiber.Ctx, err string) error {
	var statusCode int
	localize, _ := fiberi18n.Localize(c, err)

	switch {
	case err == UserNotFound:
		statusCode = fiber.StatusNotFound
	case err == InternalErrorCreatingJWT:
		statusCode = fiber.StatusInternalServerError
	case err == InvalidPassword:
		statusCode = fiber.StatusUnauthorized
	}

	return c.Status(statusCode).JSON(fiber.Map{
		"message": localize,
	})
}
