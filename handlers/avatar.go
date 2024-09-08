package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uzushikaminecraft/api/services"
	"github.com/uzushikaminecraft/api/structs"
)

// @Summary Get player's skin image
// @Description Get the specified part of player's skin image
// @Tags avatar
// @Produce  png
// @Param part path string true "which part to retrieve"
// @Param xuid path string true "XUID of target Bedrock player"
// @Failure 500 {object} structs.Error
// @Router /avatar/{part}/bedrock/{xuid} [get]
func HandleRenderBedrockSkin(c *fiber.Ctx) error {
	c.Set("Content-Type", "image/png")

	img, err := services.RenderBedrockSkin(c.Params("xuid"), c.Params("part"))
	if err != nil {
		return c.Status(500).JSON(structs.Error{
			Error: "Error occured while rendering image",
		})
	}

	c.Response().Header.Set("Cache-Control", "public, max-age=86400, immutable")
	return c.SendStream(img)
}
