package dev

import (
	"time"

	"github.com/2mugi/uzsk-api/structs"
	"gorm.io/gorm"
)

func CreateTestEntry(db *gorm.DB) {
	players := []structs.Player{
		{
			UUID:                 "550e8400-e29b-41d4-a716-446655440000",
			InitialLoginDateTime: time.Date(2024, 8, 25, 12, 0, 0, 0, time.UTC),
			LastLoginDateTime:    time.Date(2024, 8, 30, 16, 45, 0, 0, time.UTC),
			PlayTime:             7200,
			ExperiencePoints:     1500,
			Currency:             100,
			TotalBuildBlocks:     200,
			TotalDestroyBlocks:   100,
			TotalMobKills:        50,
		},
		{
			UUID:                 "e29b5500-41d4-4a71-8400-446655440001",
			InitialLoginDateTime: time.Date(2024, 8, 26, 14, 30, 0, 0, time.UTC),
			LastLoginDateTime:    time.Date(2024, 8, 29, 10, 30, 0, 0, time.UTC),
			PlayTime:             5400,
			ExperiencePoints:     1200,
			Currency:             150,
			TotalBuildBlocks:     250,
			TotalDestroyBlocks:   150,
			TotalMobKills:        75,
		},
		{
			UUID:                 "d441b400-7164-550e-29b8-446655440002",
			InitialLoginDateTime: time.Date(2024, 8, 27, 9, 15, 0, 0, time.UTC),
			LastLoginDateTime:    time.Date(2024, 8, 28, 18, 20, 0, 0, time.UTC),
			PlayTime:             3600,
			ExperiencePoints:     800,
			Currency:             200,
			TotalBuildBlocks:     150,
			TotalDestroyBlocks:   75,
			TotalMobKills:        25,
		},
		{
			UUID:                 "a716e400-8400-29b5-41d4-446655440003",
			InitialLoginDateTime: time.Date(2024, 8, 28, 11, 0, 0, 0, time.UTC),
			LastLoginDateTime:    time.Date(2024, 8, 30, 14, 0, 0, 0, time.UTC),
			PlayTime:             4800,
			ExperiencePoints:     1100,
			Currency:             80,
			TotalBuildBlocks:     175,
			TotalDestroyBlocks:   120,
			TotalMobKills:        40,
		},
		{
			UUID:                 "29b54100-4466-55e4-8400-550ea7160004",
			InitialLoginDateTime: time.Date(2024, 8, 29, 13, 45, 0, 0, time.UTC),
			LastLoginDateTime:    time.Date(2024, 8, 30, 15, 30, 0, 0, time.UTC),
			PlayTime:             6000,
			ExperiencePoints:     1400,
			Currency:             120,
			TotalBuildBlocks:     220,
			TotalDestroyBlocks:   160,
			TotalMobKills:        65,
		},
	}

	db.Create(&players)
}
