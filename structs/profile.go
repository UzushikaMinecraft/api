package structs

import "time"

type Profile struct {
	ID                   int       `gorm:"column:id;primaryKey;autoIncrement"json:"id"`
	UUID                 string    `gorm:"column:uuid;type:char(36);unique;not null"json:"uuid"`
	InitialLoginDate     time.Time `gorm:"column:initial_login_date;not null"json:"initial_login_date"`
	LastLoginDate        time.Time `gorm:"column:last_login_date;not null"json:"last_login_date"`
	TotalPlayTime        int64     `gorm:"column:total_play_time;default:0"json:"total_play_time"`
	Experience           float64   `gorm:"column:experience;default:0.0"json:"experience"`
	Currency             int       `gorm:"column:currency;default:0"json:"currency"`
	TotalBuildBlocks     int       `gorm:"column:total_build_blocks;default:0"json:"total_build_blocks"`
	TotalDestroyBlocks   int       `gorm:"column:total_destroy_blocks;default:0"json:"total_destroy_blocks"`
	TotalMobKills        int       `gorm:"column:total_mob_kills;default:0"json:"total_mob_kills"`
}

type Tabler interface {
    TableName() string
}

func (Profile) TableName() string {
	return "profile"
}