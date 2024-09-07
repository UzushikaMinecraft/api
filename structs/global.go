package structs

type Tabler interface {
	TableName() string
}

type Error struct {
	Error string `json:"error"`
}
