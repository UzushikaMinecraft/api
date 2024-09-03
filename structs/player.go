package structs

import "time"

type Player struct {
	ID                   int       `gorm:"column:id;primaryKey;autoIncrement"`
	UUID                 string    `gorm:"column:uuid;type:char(36);unique;not null"`
	InitialLoginDate     time.Time `gorm:"column:initial_login_date;not null"`
	LastLoginDate        time.Time `gorm:"column:last_login_date;not null"`
	TotalPlayTime        int64     `gorm:"column:total_play_time;default:0"`
	Experience           float64   `gorm:"column:experience;default:0.0"`
	Currency             int       `gorm:"column:currency;default:0"`
	TotalBuildBlocks     int       `gorm:"column:total_build_blocks;default:0"`
	TotalDestroyBlocks   int       `gorm:"column:total_destroy_blocks;default:0"`
	TotalMobKills        int       `gorm:"column:total_mob_kills;default:0"`
}
