package routes

import (
	"github.com/MelvinNunes/menuz-go/internal/app/middleware"
	"github.com/MelvinNunes/menuz-go/internal/handlers"
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
)

func Routes(app fiber.Router, h *handlers.Handlers, m *middleware.Middlewares) {
	v1 := app.Group("/v1")
	v1.Get("/health", handlers.GetServerHealthStatusHandler)

	v1.Post("/login", h.AuthHandler.LoginHandler)
	v1.Post("/register", h.AccountHandler.CreateAccountHandler)

	v1.Get("/accounts/me", middleware.TokenMiddleware(), h.AccountHandler.MyAccountHandler)
}

func NotFoundHandler(app fiber.Router) {
	app.Use(func(c *fiber.Ctx) error {
		localize, _ := fiberi18n.Localize(c, "route_not_found")
		return c.Status(404).JSON(fiber.Map{
			"message": localize,
		})
	})
}
