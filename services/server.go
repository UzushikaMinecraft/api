package services

import (
	"errors"

	"github.com/Craftserve/mcstatus"
	"github.com/uzushikaminecraft/api/structs"
)

// Get servers registered to uzsk-api
// @Summary Get servers
// @Description Get servers registered to uzsk-api
// @Tags servers
// @Accept  json
// @Produce  json
// @Success 200 {array} structs.Server
// @Router /servers [get]
func GetServers(config structs.Config) []structs.ServerStatus {
	servers := make([]structs.ServerStatus, 0)
	for k := range config.Servers {
		s, _ := GetServer(config, k)
		servers = append(servers, *s)
	}

	return servers
}

// Get servers registered to uzsk-api
// @Summary Get server
// @Description Get servers registered to uzsk-api
// @Tags servers
// @Accept  json
// @Produce  json
// @Param name path string true "Name of target server"
// @Success 200 {array} structs.Server
// @Failure 500 {object} structs.Error
// @Router /servers/{name} [get]
func GetServer(config structs.Config, name string) (*structs.ServerStatus, error) {
	v, ok := config.Servers[name]

	if !ok {
		return &structs.ServerStatus{
			Name:        name,
			Description: &v.Description,
			IsOnline:    false,
		}, errors.New("specified server is not registered")
	}

	// Resolve FQDN
	addr, err := mcstatus.Resolve(v.Address)
	if err != nil {
		return &structs.ServerStatus{
			Name:        name,
			Description: &v.Description,
			IsOnline:    false,
		}, errors.New("could not resolve hostname")
	}

	addr.Port = v.Port

	// Ping the server
	status, _, err := mcstatus.CheckStatus(addr)

	if err != nil {
		return &structs.ServerStatus{
			Name:        name,
			Description: &v.Description,
			IsOnline:    false,
		}, nil
	}

	return &structs.ServerStatus{
		Name:          name,
		Description:   &v.Description,
		IsOnline:      true,
		OnlinePlayers: &status.Players,
		MaxPlayers:    &status.Slots,
		Version:       &status.GameVersion,
		PlayersSample: &status.PlayersSample,
	}, nil
}
