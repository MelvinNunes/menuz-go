package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/MelvinNunes/menuz-go/docs"
	"github.com/MelvinNunes/menuz-go/internal/app/lang"
	"github.com/MelvinNunes/menuz-go/internal/app/middleware"
	"github.com/MelvinNunes/menuz-go/internal/app/routes"
	"github.com/MelvinNunes/menuz-go/internal/app/validators"
	"github.com/MelvinNunes/menuz-go/internal/infrastructure/database"
	"github.com/MelvinNunes/menuz-go/internal/infrastructure/injection"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
)

var (
	DefaultPort string = "8000"
)

// @title MENUZ API
// @version 1.0
// @description This is documentation for Menuz in version 1.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email melvinfulana@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /api/v1
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file. Please add the .env file for the project!")
	}

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "MENUZ-API",
		AppName:       "Menuz v1.0.0",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(validators.GlobalErrorHandlerResp{
				Success: false,
				Message: "An internal error occurred!",
			})
		},
	})

	app.Use(logger.New(middleware.LoggerConfig()))
	app.Use(cors.New())
	app.Use(recover.New())

	db := database.InitDatabase()
	injection.InjectReposAndService(db)

	database.RunMigrations()
	database.SeedDatabase()

	app.Get("/docs/*", swagger.HandlerDefault)
	lang.ConfigureInternationalization(app)

	api := app.Group("/api")
	routes.Routes(api)
	routes.NotFoundHandler(api)

	validators.EnableValidation()

	port := os.Getenv("PORT")
	if port == "" {
		port = DefaultPort
	}

	err = app.Listen(fmt.Sprintf(":%v", port))

	if err != nil {
		log.Fatal(err)
	}
}
