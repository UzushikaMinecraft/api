package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uzushikaminecraft/api/services"
)

// Get servers registered to uzsk-api
// @Summary Get server
// @Description Get servers registered to uzsk-api
// @Tags servers
// @Accept  json
// @Produce  json
// @Param name path string true "Name of target server"
// @Success 200 {object} structs.Server
// @Failure 500 {object} structs.Error
// @Router /servers/{name} [get]
func HandleServer(c *fiber.Ctx) error {
	res, _ := services.GetServer(c.Params("name"))
	return c.JSON(res)
}

// @Summary Get servers
// @Description Get servers registered to uzsk-api
// @Tags servers
// @Accept  json
// @Produce  json
// @Success 200 {array} structs.Server
// @Router /servers [get]
func HandleServers(c *fiber.Ctx) error {
	return c.JSON(services.GetServers())
}
