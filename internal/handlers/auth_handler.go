package handlers

import (
	"github.com/MelvinNunes/menuz-go/internal/app/validators"
	"github.com/MelvinNunes/menuz-go/internal/domain/service"
	"github.com/MelvinNunes/menuz-go/internal/infrastructure/dtos"
	errorchecker "github.com/MelvinNunes/menuz-go/internal/infrastructure/error"
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authenticationService *service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authenticationService: &authService}
}

func (h *AuthHandler) LoginHandler(c *fiber.Ctx) error {
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

	res, err := h.authenticationService.Login(data.Email, data.Password)
	if err != nil {
		solvedError := errorchecker.ResolveError(c, err.Error())
		return solvedError
	}

	localize, _ := fiberi18n.Localize(c, "auth.success")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    *res,
		"message": localize,
	})
}
