package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uzushikaminecraft/api/auth"
	"github.com/uzushikaminecraft/api/structs"
)

// @Summary Login with Discord
// @Description Login with Discord
// @Tags login
// @Accept  json
// @Produce  json
// @Router /auth [get]
// @Success 200
// @Header 200 {string} Location
func HandleAuth(c *fiber.Ctx) error {
	url := auth.Auth(c)

	c.Status(fiber.StatusSeeOther)
	c.Redirect(url)

	return c.JSON(url)
}

// @Summary callback endpoint for Discord login
// @Description callback endpoint for Discord login
// @Tags login
// @Accept  json
// @Produce  json
// @Param code query string true "Bearer token"
// @Param state query string true "Random state for validating request"
// @Success 200 {array} structs.JWTResponse
// @Header 200 {string} X-Auth-Token
// @Failure 500 {object} structs.Error
// @Router /auth/callback [get]
func HandleAuthCallback(c *fiber.Ctx) error {
	jwtAccessToken, err := auth.Callback(
		c.Params("state"), c.Params("code"),
	)

	if err == nil {
		c.Response().Header.Add("X-Auth-Token", *jwtAccessToken)
		return c.JSON(
			structs.JWTResponse{
				Success: true,
			},
		)
	}

	e := err.Error()

	if e == "state string does not match" || e == "required parameter is not provided" {
		return c.Status(400).JSON(
			structs.Error{
				Error: e,
			},
		)
	}

	return c.Status(500).JSON(
		structs.Error{
			Error: e,
		},
	)
}
