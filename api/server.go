package api

import (
	"log"

	"github.com/2mugi/uzsk-api/structs"
	"github.com/Craftserve/mcstatus"

	"github.com/gofiber/fiber/v2"
)

func GetServer(config structs.Config, name string) fiber.Map {
	v, ok := config.Servers[name]

	if !ok {
		return fiber.Map{
			"error": "No such server",
		}
	}

	// Resolve FQDN
	addr, err := mcstatus.Resolve(v.Address)
	if err != nil {
		log.Println(err)
		return fiber.Map{
			"error": "Failed to resolve server address",
		}
	}

	addr.Port = v.Port

	// Ping the server
	status, _, err := mcstatus.CheckStatus(addr)

	if err != nil {
		return fiber.Map{
			"name":      name,
			"is_online": false,
		}
	}

	return fiber.Map{
		"name":           name,
		"is_online":      true,
		"online_players": status.Players,
		"max_players":    status.Slots,
		"version":        status.GameVersion,
		"players_sample": status.PlayersSample,
	}
}
