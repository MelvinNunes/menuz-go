package interfaces

import (
	"github.com/MelvinNunes/menuz-go/internal/app/validators"
	"github.com/MelvinNunes/menuz-go/internal/domain/service"
	"github.com/MelvinNunes/menuz-go/internal/infrastructure/dtos"
	"github.com/MelvinNunes/menuz-go/internal/infrastructure/security"
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
)

func LoginHandler(c *fiber.Ctx) error {
	data := dtos.LoginDTO{}
	if err := c.BodyParser(&data); err != nil {
		localize, _ := fiberi18n.Localize(c, "body_parse_error")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": localize,
		})
	}

	validationMessage, err := validators.ValidateData(data)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"code": fiber.StatusBadRequest,
			"data": *validationMessage,
		})
	}

	user := service.UserService.GetUserByEmail(data.Email)
	if user == nil {
		localize, _ := fiberi18n.Localize(c, "user.not_found")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": localize,
		})
	}

	if !security.IsPasswordEqualToHash(data.Password, user.Password) {
		localize, _ := fiberi18n.Localize(c, "auth.invalid_password")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": localize,
		})
	}

	token, err := security.CreateJWTtoken(user.ID.String())

	if err != nil {
		localize, _ := fiberi18n.Localize(c, "internal_error")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"token":   token,
			"message": localize,
		})
	}

	localize, _ := fiberi18n.Localize(c, "auth.success")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token":   token,
		"message": localize,
	})
}
