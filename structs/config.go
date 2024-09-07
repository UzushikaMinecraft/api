package structs

type Config struct {
	Servers     map[string]Server
	MySQL       MySQL
	General     General
	Credentials Credentials
}

type MySQL struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
}

type General struct {
	JWTSecret   string
	CallbackURL string `toml:"callback_url"`
}

type Credentials struct {
	State   string
	Discord CredentialsDiscord
}

type CredentialsDiscord struct {
	ClientID     string `toml:"client_id"`
	ClientSecret string `toml:"client_secret"`
}
