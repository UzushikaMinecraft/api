package services

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/uzushikaminecraft/api/db"
	"github.com/uzushikaminecraft/api/structs"
)

// fetch UUID from provided Discord ID
// @Summary fetch UUID from provided Discord ID
// @Description fetch UUID from provided Discord ID
// @Tags discord
// @Produce json
// @Param uuid path string true "who to retrieve"
// @Failure 200 {object} structs.UUID
// @Failure 500 {object} structs.Error
// @Router /api/discord/{uuid} [get]
func GetUUIDByDiscord(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return c.Status(400).JSON(structs.Error{
			Error: "uuid is not provided",
		})
	}

	var u structs.DiscordSRVUser

	res := db.DB.
		Where("discord == ?", uuid).
		First(&u)

	if res.RowsAffected == 0 {
		return c.Status(404).JSON(structs.Error{
			Error: "no players found",
		})
	}
	if res.Error != nil {
		return c.Status(500).JSON(structs.Error{
			Error: fmt.Sprintf("error occured while processing request: %v", res.Error),
		})
	}

	return c.Status(200).JSON(
		structs.UUID{
			UUID: u.UUID,
		},
	)
}
