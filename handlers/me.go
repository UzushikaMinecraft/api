package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uzushikaminecraft/api/services"
	"github.com/uzushikaminecraft/api/structs"
)

// @Summary retrieve information of authenticated user
// @Description retrieve information of authenticated user
// @Tags login
// @Accept json
// @Param X-Auth-Token header string true "JSON Web Token"
// @Success 200 {object} structs.Me
// @Failure 400 {object} structs.Error
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
