package services

import (
	"errors"

	"github.com/Craftserve/mcstatus"
	"github.com/uzushikaminecraft/api/config"
	"github.com/uzushikaminecraft/api/structs"
)

func GetServers() []structs.ServerStatus {
	servers := make([]structs.ServerStatus, 0)
	for k := range config.Conf.Servers {
		s, _ := GetServer(k)
		servers = append(servers, *s)
	}

	return servers
}

func GetServer(name string) (*structs.ServerStatus, error) {
	v, ok := config.Conf.Servers[name]

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
