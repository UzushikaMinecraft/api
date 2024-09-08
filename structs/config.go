package structs

type Config struct {
	Servers     map[string]RegisteredServer
	MySQL       MySQL
	General     General
	Credentials Credentials
}

type MySQL struct {
	Core       MySQLCore
	DiscordSRV MySQLDiscordSRV
}

type MySQLCore struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
}

type MySQLDiscordSRV struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
	Prefix   string
}

type General struct {
	CallbackURL string `toml:"callback_url"`
}

type Credentials struct {
	State     string
	JWTSecret string `toml:"jwt_secret"`
	Discord   CredentialsDiscord
}

type CredentialsDiscord struct {
	ClientID     string `toml:"client_id"`
	ClientSecret string `toml:"client_secret"`
}

type RegisteredServer struct {
	Address     string
	Port        int
	Description string
}
