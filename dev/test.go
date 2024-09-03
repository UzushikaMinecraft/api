package dev

import (
	"time"

	"github.com/2mugi/uzsk-api/structs"
	"gorm.io/gorm"
)

func CreateTestEntry(db *gorm.DB) {
	players := []structs.Player{
		{
			uuid:                    "550e8400-e29b-41d4-a716-446655440000",
			initial_login_date_time: time.Date(2024, 8, 25, 12, 0, 0, 0, time.UTC),
			last_login_date_time:    time.Date(2024, 8, 30, 16, 45, 0, 0, time.UTC),
			total_play_time:         7200,
			experience:              1500,
			currency:                100,
			total_build_blocks:      200,
			total_destroy_blocks:    100,
			total_mob_kills:         50,
		},
		{
			uuid:                    "e29b5500-41d4-4a71-8400-446655440001",
			initial_login_date_time: time.Date(2024, 8, 26, 14, 30, 0, 0, time.UTC),
			last_login_date_time:    time.Date(2024, 8, 29, 10, 30, 0, 0, time.UTC),
			total_play_time:         5400,
			experience:              1200,
			currency:                150,
			total_build_blocks:      250,
			total_destroy_blocks:    150,
			total_mob_kills:         75,
		},
		{
			uuid:                    "d441b400-7164-550e-29b8-446655440002",
			initial_login_date_time: time.Date(2024, 8, 27, 9, 15, 0, 0, time.UTC),
			last_login_date_time:    time.Date(2024, 8, 28, 18, 20, 0, 0, time.UTC),
			total_play_time:         3600,
			experience:              800,
			currency:                200,
			total_build_blocks:      150,
			total_destroy_blocks:    75,
			total_mob_kills:         25,
		},
		{
			uuid:                    "a716e400-8400-29b5-41d4-446655440003",
			initial_login_date_time: time.Date(2024, 8, 28, 11, 0, 0, 0, time.UTC),
			last_login_date_time:    time.Date(2024, 8, 30, 14, 0, 0, 0, time.UTC),
			total_play_time:         4800,
			experience:              1100,
			currency:                80,
			total_build_blocks:      175,
			total_destroy_blocks:    120,
			total_mob_kills:         40,
		},
		{
			uuid:                    "29b54100-4466-55e4-8400-550ea7160004",
			initial_login_date_time: time.Date(2024, 8, 29, 13, 45, 0, 0, time.UTC),
			last_login_date_time:    time.Date(2024, 8, 30, 15, 30, 0, 0, time.UTC),
			total_play_time:         6000,
			experience:              1400,
			currency:                120,
			total_build_blocks:      220,
			total_destroy_blocks:    160,
			total_mob_kills:         65,
		},
	}

	db.Create(&players)
}
