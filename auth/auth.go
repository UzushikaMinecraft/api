package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uzushikaminecraft/api/config"
)

func Auth(c *fiber.Ctx) string {
	url := oauthConf.AuthCodeURL(config.Conf.Credentials.State)

	return url
}
