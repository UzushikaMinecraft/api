package structs

type Tabler interface {
	TableName() string
}

type Error struct {
	Error string `json:"error"`
}

type UUID struct {
	UUID string `json:"uuid"`
}
