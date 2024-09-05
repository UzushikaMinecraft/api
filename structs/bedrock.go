type Bedrock struct {
	ID                   int       `gorm:"column:id;primaryKey;autoIncrement"json:"id"`
	FUID                 string    `gorm:"column:fuid;type:char(36);unique;not null"json:"uuid"`
	XUID				 string    `gorm:"column:xuid;type:char(36);unique;not null"json:"xuid"`
}

type Tabler interface {
    TableName() string
}

func (Bedrock) TableName() string {
	return "bedrock"
}