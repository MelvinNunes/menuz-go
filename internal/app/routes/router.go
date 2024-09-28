package routes

import (
	"github.com/MelvinNunes/menuz-go/internal/app/middleware"
	"github.com/MelvinNunes/menuz-go/internal/domain/repository"
	"github.com/MelvinNunes/menuz-go/internal/domain/service"
	"github.com/MelvinNunes/menuz-go/internal/interfaces"
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Routes(db *gorm.DB, app fiber.Router) {
	// init repository and service
	userRepo := repository.NewUserRepository(db)
	userProfileRepo := repository.NewProfileRepository(db)
	userRoleRepo := repository.NewUserRoleRepository(db)

	userService := service.NewUserService(*userRepo, *userProfileRepo, *userRoleRepo)
	authService := service.NewAuthService(*userService)
	// end repository and service initialization

	// Handlers
	authHandler := interfaces.NewAuthHandler(*authService)
	accountHandler := interfaces.NewAccountHandler(*userService)

	v1 := app.Group("/v1")
	v1.Get("/health", interfaces.GetServerHealthStatusHandler)

	v1.Post("/login", authHandler.LoginHandler)
	v1.Post("/register", accountHandler.CreateAccountHandler)

	v1.Get("/accounts/me", middleware.TokenMiddleware(), accountHandler.MyAccountHandler)
}

func NotFoundHandler(app fiber.Router) {
	app.Use(func(c *fiber.Ctx) error {
		localize, _ := fiberi18n.Localize(c, "route_not_found")
		return c.Status(404).JSON(fiber.Map{
			"message": localize,
		})
	})
}
