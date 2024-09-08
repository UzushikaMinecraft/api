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
// @Param X-Auth-Token header string true "JSON Web Token"
// @Success 200 {object} structs.Me
// @Failure 400 {object} structs.Error
// @Failure 500 {object} structs.Error
// @Router /me [get]
func HandleMe(c *fiber.Ctx) error {
	res, err := services.GetMe(c.Get("X-Auth-Token"))

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
// @Param X-Auth-Token header string true "JSON Web Token"
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

	res, err := services.UpdateBiography(
		c.Get("X-Auth-Token"), b.Biography,
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
