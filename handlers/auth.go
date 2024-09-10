package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/uzushikaminecraft/api/auth"
	"github.com/uzushikaminecraft/api/config"
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
// @Tags auth
// @Accept  json
// @Produce  json
// @Param code query string true "Bearer token"
// @Param state query string true "Random state for validating request"
// @Success 200 {array} structs.JWTResponse
// @Header 200 {string} Location
// @Failure 400 {object} structs.Error
// @Failure 500 {object} structs.Error
// @Router /auth/callback [get]
func HandleAuthCallback(c *fiber.Ctx) error {
	jwtCallback, err := auth.Callback(
		c.Query("state"), c.Query("code"),
	)

	if err == nil {
		cookie := new(fiber.Cookie)
		cookie.Name = "accessToken"
		cookie.Value = jwtCallback.AccessToken
		cookie.SameSite = "Strict"
		cookie.Secure = true
		cookie.HTTPOnly = true
		cookie.Expires = jwtCallback.Claims["exp"].(time.Time)
		c.Cookie(cookie)

		c.Status(301).Redirect("/?loggedIn=success")

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

// @Summary refresh token with provided access token
// @Description refresh token with provided access token
// @Tags auth
// @Accept  json
// @Produce  json
// @Success 200
// @Header 200 {string} Location
// @Failure 400 {object} structs.Error
// @Failure 500 {object} structs.Error
// @Router /auth/token/refresh [get]
func HandleAuthTokenRefresh(c *fiber.Ctx) error {
	cookie := new(structs.CoreCookie)
	if err := c.CookieParser(cookie); err != nil {
		return c.Status(400).JSON(
			structs.Error{
				Error: "provided Cookie is not valid",
			},
		)
	}

	if cookie.AccessToken == "" {
		c.Status(400).JSON(
			structs.Error{
				Error: "no token is provided",
			},
		)
	}

	claims, err := auth.Validate(
		cookie.AccessToken,
	)

	if err != nil {
		c.Status(400).JSON(
			structs.Error{
				Error: "provided token is not valid",
			},
		)
	}

	c.ClearCookie("accessToken")

	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newTokenString, err := newToken.SignedString([]byte(config.Conf.Credentials.JWTSecret))
	if err != nil {
		c.Status(500).JSON(
			structs.Error{
				Error: fmt.Sprintf("failed to generate new token: %v", err),
			},
		)
	}

	newCookie := new(fiber.Cookie)
	newCookie.Name = "accessToken"
	newCookie.Value = newTokenString
	newCookie.SameSite = "Strict"
	newCookie.Secure = true
	newCookie.HTTPOnly = true
	newCookie.Expires = claims["exp"].(time.Time)
	c.Cookie(newCookie)

	c.Status(301).Redirect("/?loggedIn=success")

	return c.SendStatus(200)
}
