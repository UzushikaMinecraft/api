package structs

import "time"

type Player struct {
	ID                   int       `gorm:"primaryKey;autoIncrement"`
	UUID                 string    `gorm:"type:char(36);unique;not null"`
	InitialLoginDateTime time.Time `gorm:"not null"`
	LastLoginDateTime    time.Time
	PlayTime             int `gorm:"default:0"`
	ExperiencePoints     int `gorm:"default:0"`
	Currency             int `gorm:"default:0"`
	TotalBuildBlocks     int `gorm:"default:0"`
	TotalDestroyBlocks   int `gorm:"default:0"`
	TotalMobKills        int `gorm:"default:0"`
}
