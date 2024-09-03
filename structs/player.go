package structs

import "time"

type Player struct {
	id                   int       `gorm:"primaryKey;autoIncrement"`
	uuid                 string    `gorm:"type:char(36);unique;not null"`
	initial_login_date   time.Time `gorm:"not null"`
	last_login_date      time.Time
	total_play_time      bigint  `gorm:"default:0"`
	experience           float64 `gorm:"default:0.0"`
	currency             int     `gorm:"default:0"`
	total_build_blocks   int     `gorm:"default:0"`
	total_destroy_blocks int     `gorm:"default:0"`
	total_mob_kills      int     `gorm:"default:0"`
}
