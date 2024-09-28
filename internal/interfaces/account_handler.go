package interfaces

import (
	"github.com/MelvinNunes/menuz-go/internal/app/validators"
	"github.com/MelvinNunes/menuz-go/internal/domain/service"
	"github.com/MelvinNunes/menuz-go/internal/infrastructure/dtos"
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
)

type AccountHandler struct {
	userService    *service.UserService
	roleService    *service.RoleService
	accountService *service.AccountService
}

func NewAccountHandler(userService service.UserService) *AccountHandler {
	return &AccountHandler{userService: &userService}
}

func (h *AccountHandler) CreateAccountHandler(c *fiber.Ctx) error {
	data := dtos.CreateAccount{}
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

	if h.userService.UserExistsByEmail(data.Email) {
		localize, _ := fiberi18n.Localize(c, "user.email_already_exists")
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": localize,
		})
	}

	if h.userService.UserExistsByPhoneNumber(data.PhoneNumberCode, data.PhoneNumber) {
		localize, _ := fiberi18n.Localize(c, "user.phone_already_exists")
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": localize,
		})
	}

	role := h.roleService.GetRoleByName(data.Role)
	if role == nil {
		localize, _ := fiberi18n.Localize(c, "role.not_found")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": localize,
		})
	}

	err = h.accountService.CreateAccount(&data, *role)
	if err != nil {
		localize, _ := fiberi18n.Localize(c, "account.error_creating")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": localize,
		})
	}

	localize, _ := fiberi18n.Localize(c, "account.created")
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": localize,
	})
}

func (h *AccountHandler) MyAccountHandler(c *fiber.Ctx) error {
	user, _ := h.userService.GetOnlineUser(c)
	if user == nil {
		localize, _ := fiberi18n.Localize(c, "user.not_found")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": localize,
		})
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"data": user,
	})
}
