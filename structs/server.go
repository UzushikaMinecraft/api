package structs

type Server struct {
	Address     string
	Port        int
	Description string
}

type ServerStatus struct {
	Name           string  `json:"name"`
	Description    *string `json:"description"`
	IsOnline       bool    `json:"is_online"`
	OnlinePlayers  *int    `json:"online_players"`
	MaxPlayers     *int    `json:"max_players"`
	Version        *string `json:"version"`
	PlayersSample  *[]string `json:"players_sample"`
}