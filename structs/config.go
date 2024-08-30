package structs

type MySQL struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
}

type Config struct {
	Servers map[string]Server
	MySQL   MySQL
}
