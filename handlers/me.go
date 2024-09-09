package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uzushikaminecraft/api/services"
	"github.com/uzushikaminecraft/api/structs"
)

// @Summary retrieve information of authenticated user
// @Description retrieve information of authenticated user
// @Tags auth
// @Accept json
// @Success 200 {object} structs.Me
// @Failure 400 {object} structs.Error
// @Failure 500 {object} structs.Error
// @Router /me [get]
func HandleMe(c *fiber.Ctx) error {
	cookie := new(structs.CoreCookie)
	if err := c.CookieParser(cookie); err != nil {
		c.Status(301).Redirect("/api/auth")
	}

	if cookie.JWT == "" {
		c.Status(301).Redirect("/api/auth")
	}

	res, err := services.GetMe(cookie.JWT)

	if err == nil {
		return c.Status(200).JSON(
			res,
		)
	}

	if err.Error() == "token is not provided" || err.Error() == "invalid token" {
		return c.Status(400).JSON(
			structs.Error{
				Error: err.Error(),
			},
		)
	}

	return c.Status(500).JSON(
		structs.Error{
			Error: err.Error(),
		},
	)
}

// @Summary update biography of authenticated user
// @Description update biography of authenticated user
// @Tags auth
// @Accept json
// @Param Biography body structs.Biography true "new biography"
// @Success 200 {object} structs.Me
// @Failure 400 {object} structs.Error
// @Failure 500 {object} structs.Error
// @Router /me/biography [post]
func HandleMeBiography(c *fiber.Ctx) error {
	var b structs.Biography
	if err := c.BodyParser(&b); err != nil {
		return c.Status(400).JSON(
			structs.Error{
				Error: "request body is not valid",
			},
		)
	}

	cookie := new(structs.CoreCookie)
	if err := c.CookieParser(cookie); err != nil {
		c.Status(301).Redirect("/api/auth")
	}

	if cookie.JWT == "" || cookie == nil {
		c.Status(301).Redirect("/api/auth")
	}

	res, err := services.UpdateBiography(
		cookie.JWT, b.Biography,
	)

	if err == nil {
		return c.Status(200).JSON(
			res,
		)
	}

	if err.Error() == "token is not provided" || err.Error() == "invalid token" {
		return c.Status(400).JSON(
			structs.Error{
				Error: err.Error(),
			},
		)
	}

	return c.Status(500).JSON(
		structs.Error{
			Error: err.Error(),
		},
	)
}
