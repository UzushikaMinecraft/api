package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uzushikaminecraft/api/services"
	"github.com/uzushikaminecraft/api/structs"
)

// @Summary fetch UUID from provided Discord ID
// @Description fetch UUID from provided Discord ID
// @Tags discord
// @Produce json
// @Param uuid path string true "who to retrieve"
// @Failure 200 {object} structs.UUID
// @Failure 500 {object} structs.Error
// @Router /api/discord/{uuid} [get]
func HandleDiscordUuid(c *fiber.Ctx) error {
	uuid, err := services.GetDiscordByUUID(c.Params("uuid"))

	if err == nil {
		return c.JSON(
			structs.UUID{
				UUID: *uuid,
			},
		)
	}

	e := err.Error()

	if e == "uuid is not provided" {
		return c.Status(400).JSON(
			structs.Error{
				Error: e,
			},
		)
	}

	if e == "no players found" {
		return c.Status(404).JSON(
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
