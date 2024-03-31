package interfaces

import (
	"log"

	mailer "github.com/MelvinNunes/menuz-go/internal/app/integrations/email"
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
)

// Health godoc
// @Summary      Show server health
// @Description  Shows if server is up
// @Tags         Health
// @Accept       json
// @Produce      json
// @Success      200  {object}  string (string)
// @Failure      500  {object}  string (string)
// @Router       /health [get]
func GetServerHealthStatusHandler(c *fiber.Ctx) error {
	mailer.SendMail()

	localize, err := fiberi18n.Localize(c, "health")
	if err != nil {
		log.Panic("Error in fiberi18n.Localize: ", err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": localize,
	})
}
