package structs

type Server struct {
	Address     string
	Port        int
	Description string
}

type ServerStatus struct {
	Name           string  `json:"name"`
	Description    string `json:"description,omitempty"`
	IsOnline       bool    `json:"is_online"`
	OnlinePlayers  int    `json:"online_players,omitempty"`
	MaxPlayers     int    `json:"max_players,omitempty"`
	Version        string `json:"version,omitempty"`
	PlayersSample  []string `json:"players_sample,omitempty"`
}