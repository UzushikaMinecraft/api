package api

import (
	"fmt"
	"strconv"

	"github.com/2mugi/uzsk-api/structs"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetPlayers(db *gorm.DB, m map[string]string) fiber.Map {
	// Check if required parameters were provided
	filter, filter_ok := m["filter"]
	sort, sort_ok := m["sort"]
	offset, offset_ok := m["offset"]
	if !(filter_ok && sort_ok && offset_ok) {
		return fiber.Map{
			"error": "Required parameters not provided",
		}
	}

	// Parse parameters
	if !(sort == "desc" || sort == "asc") {
		return fiber.Map{
			"error": "Parameter `sort` is not correct",
		}
	}

	o, err := strconv.Atoi(offset)
	if err != nil {
		return fiber.Map{
			"error": "Parameter `offset` is not valid",
		}
	}

	var players []structs.Player
	db.
		Where("uuid LIKE ?", "%"+filter+"%").
		Order(fmt.Sprintf("uuid %v", sort)).
		Limit(30).
		Offset(o).
		Find(&players)

	return fiber.Map{
		"players": players,
	}
}
