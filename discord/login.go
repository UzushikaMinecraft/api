package discord

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uzushikaminecraft/api/config"
)

// Login with Discord
// @Summary Login with Discord
// @Description Login with Discord
// @Tags login
// @Accept  json
// @Produce  json
// @Router /login [get]
func Login(c *fiber.Ctx) error {
	url := oauthConf.AuthCodeURL(config.Conf.Credentials.State)

	c.Status(fiber.StatusSeeOther)
	c.Redirect(url)

	return c.JSON(url)
}
