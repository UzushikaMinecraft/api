package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uzushikaminecraft/api/services"
	"github.com/uzushikaminecraft/api/structs"
)

// @Summary Get profiles
// @Description Get a list of profiles with optional filtering and sorting, etc.
// @Tags profiles
// @Accept  json
// @Produce  json
// @Param filter query string false "Filter criteria" example(550e8400-e29b-41d4-a716-446655440000) default()
// @Param sort query string false "Sort order" example(desc) default(asc)
// @Param offset query int false "Offset for pagination" example(0) default(0)
// @Param limit query int false "Limit for pagination" example(10) default(50)
// @Param order_by query string false "Order by field" example(play_time)
// @Success 200 {array} structs.Profile
// @Failure 500 {object} structs.Error
// @Router /profiles [get]
func HandleProfiles(c *fiber.Ctx) error {
	m := c.Queries()

	res, err := services.GetProfiles(m)
	if err != nil {
		return c.Status(500).JSON(
			structs.Error{
				Error: err.Error(),
			},
		)
	}

	return c.JSON(res)
}

// @Summary Get profile
// @Description Get a profile by UUID
// @Tags profiles
// @Accept  json
// @Produce  json
// @Param uuid path string true "UUID of target profile"
// @Success 200 {object} structs.Profile
// @Failure 500 {object} structs.Error
// @Router /profiles/{uuid} [get]
func HandleProfile(c *fiber.Ctx) error {
	res, _ := services.GetProfile(c.Params("uuid"))
	return c.JSON(res)
}
