package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uzushikaminecraft/api/structs"
)

// retrieve information of authenticated user
// @Summary retrieve information of authenticated user
// @Description retrieve information of authenticated user
// @Tags login
// @Accept  json
// @Produce  json
// @Param token query string true "JSON Web Token"
// @Router /me [get]
func GetMe(c *fiber.Ctx) error {
	if c.Params("token") == "" {
		return c.Status(400).JSON(structs.Error{
			Error: "token is not provided",
		})
	}

	claims, err := Validate(c.Params("token"))
	if err != nil {
		return c.Status(400).JSON(structs.Error{
			Error: "invalid token",
		})
	}

	return c.Status(200).JSON(structs.Me{
		UserId:          claims["user_id"].(string),
		SessionExpireAt: claims["exp"].(int64),
	})
}
