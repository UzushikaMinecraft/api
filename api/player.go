package api

import (
	"fmt"
	"strconv"

	"github.com/2mugi/uzsk-api/structs"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Get players with query parameters
// @Summary Get players
// @Description Get a list of players with optional filtering and sorting, etc.
// @Tags players
// @Accept  json
// @Produce  json
// @Param filter query string false "Filter criteria" example(550e8400-e29b-41d4-a716-446655440000) default()
// @Param sort query string false "Sort order" example(desc) default(asc)
// @Param offset query int false "Offset for pagination" example(0) default(0)
// @Param limit query int false "Limit for pagination" example(10) default(50)
// @Param order_by query string false "Order by field" example(play_time)
// @Success 200 {array} structs.Player
// @Router /players [get]
func GetPlayers(db *gorm.DB, m map[string]string) fiber.Map {
	var err error

	// Check if required parameters were provided
	filter, _ := m["filter"]
	sort, sort_ok := m["sort"]
	offset, offset_ok := m["offset"]
	limit, limit_ok := m["limit"]

	// Parse parameters
	// sort
	if !(sort_ok) {
		sort = "asc"
	}
	if !(sort == "desc" || sort == "asc") {
		return fiber.Map{
			"error": "Parameter `sort` is not correct",
		}
	}

	// order_by
	order_by, order_by_ok := m["order_by"]
	if !(order_by_ok) {
		order_by = "uuid"
	}

	// offset
	var o int
	if offset_ok {
		o, err = strconv.Atoi(offset)
		if err != nil {
			return fiber.Map{
				"error": "Parameter `offset` is not valid",
			}
		}
	}

	// limit
	var l int
	if limit_ok {
		l, err := strconv.Atoi(limit)
		if err != nil || l < 0 {
			return fiber.Map{
				"error": "Parameter `limit` is not valid",
			}
		}
		if l > 50 {
			l = 50
		}
	} else {
		l = 50
	}

	if !(order_by == "uuid" || order_by == "experience" || order_by == "currency" || order_by == "total_build_blocks" || order_by == "total_destroy_blocks" || order_by == "total_mob_kills" || order_by == "total_play_time") {
		return fiber.Map{
			"error": "Parameter `order_by` is not valid",
		}
	}

	var players []structs.Player
	db.
		Where("uuid LIKE ?", "%"+filter+"%").
		Order(fmt.Sprintf("%v %v", order_by, sort)).
		Offset(o).
		Limit(l).
		Find(&players)
	db.Find(&players)

	return fiber.Map{ 
		"players": players,
	}
}

// Get player by UUID
// @Summary Get player
// @Description Get a player by UUID
// @Tags players
// @Accept  json
// @Produce  json
// @Param uuid path string true "UUID of target player"
// @Success 200 {object} structs.Player
// @Router /players/{uuid} [get]
func GetPlayer(db *gorm.DB, uuid string) fiber.Map {
	var player structs.Player
	db.Where("uuid = ?", uuid).First(&player)

	if player.UUID == "" {
		return fiber.Map{
			"error": "No such player",
		}
	}

	return fiber.Map{
		"player": player,
	}
}
