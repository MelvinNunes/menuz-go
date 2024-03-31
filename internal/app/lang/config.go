package lang

import (
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/text/language"
)

func ConfigureInternationalization(app *fiber.App) {
	// Middleware that handles translation
	app.Use(
		fiberi18n.New(&fiberi18n.Config{
			RootPath:        "./internal/app/lang/locales",
			AcceptLanguages: []language.Tag{language.Portuguese, language.English},
			DefaultLanguage: language.Portuguese,
		}),
	)
}
