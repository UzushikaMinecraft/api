package dev

import (
	"time"

	"github.com/uzushikaminecraft/api/db"
	"github.com/uzushikaminecraft/api/structs"
)

func CreateTestEntry() {
	profiles := []structs.Profile{
		{
			UUID:               "550e8400-e29b-41d4-a716-446655440000",
			InitialLoginDate:   time.Date(2024, 8, 25, 12, 0, 0, 0, time.UTC),
			LastLoginDate:      time.Date(2024, 8, 30, 16, 45, 0, 0, time.UTC),
			TotalPlayTime:      7200,
			Experience:         1500.0,
			Currency:           100,
			TotalBuildBlocks:   200,
			TotalDestroyBlocks: 100,
			TotalMobKills:      50,
		},
		{
			UUID:               "e29b5500-41d4-4a71-8400-446655440001",
			InitialLoginDate:   time.Date(2024, 8, 26, 14, 30, 0, 0, time.UTC),
			LastLoginDate:      time.Date(2024, 8, 29, 10, 30, 0, 0, time.UTC),
			TotalPlayTime:      5400,
			Experience:         1200.0,
			Currency:           150,
			TotalBuildBlocks:   250,
			TotalDestroyBlocks: 150,
			TotalMobKills:      75,
		},
		{
			UUID:               "d441b400-7164-550e-29b8-446655440002",
			InitialLoginDate:   time.Date(2024, 8, 27, 9, 15, 0, 0, time.UTC),
			LastLoginDate:      time.Date(2024, 8, 28, 18, 20, 0, 0, time.UTC),
			TotalPlayTime:      3600,
			Experience:         800.0,
			Currency:           200,
			TotalBuildBlocks:   150,
			TotalDestroyBlocks: 75,
			TotalMobKills:      25,
		},
		{
			UUID:               "a716e400-8400-29b5-41d4-446655440003",
			InitialLoginDate:   time.Date(2024, 8, 28, 11, 0, 0, 0, time.UTC),
			LastLoginDate:      time.Date(2024, 8, 30, 14, 0, 0, 0, time.UTC),
			TotalPlayTime:      4800,
			Experience:         1100.0,
			Currency:           80,
			TotalBuildBlocks:   175,
			TotalDestroyBlocks: 120,
			TotalMobKills:      40,
		},
		{
			UUID:               "29b54100-4466-55e4-8400-550ea7160004",
			InitialLoginDate:   time.Date(2024, 8, 29, 13, 45, 0, 0, time.UTC),
			LastLoginDate:      time.Date(2024, 8, 30, 15, 30, 0, 0, time.UTC),
			TotalPlayTime:      6000,
			Experience:         1400.0,
			Currency:           120,
			TotalBuildBlocks:   220,
			TotalDestroyBlocks: 160,
			TotalMobKills:      65,
		},
	}

	db.DB.Create(&profiles)
}
