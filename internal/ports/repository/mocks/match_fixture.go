package mocks

import "github.com/brcodingdev/cloudwalk-parser/internal/pkg/model"

// MatchLogFile5Fixture fixture for log_5.log file
var MatchLogFile5Fixture = model.Match{
	ID:         1,
	TotalKills: 131,
	Players:    []model.Player{"Isgalamido", "Oootsimo", "Dono da Bola", "Assasinu Credi", "Zeh", "Mal"},
	Kills: map[model.Player]int{
		"Isgalamido":     17,
		"Oootsimo":       20,
		"Dono da Bola":   10,
		"Assasinu Credi": 13,
		"Zeh":            19,
		"Mal":            6,
	},
	PlayerRanking: []model.Ranking{
		{Position: 1, Player: "Oootsimo", TotalKills: 20},
		{Position: 2, Player: "Zeh", TotalKills: 19},
		{Position: 3, Player: "Isgalamido", TotalKills: 17},
		{Position: 4, Player: "Assasinu Credi", TotalKills: 13},
		{Position: 5, Player: "Dono da Bola", TotalKills: 10},
		{Position: 6, Player: "Mal", TotalKills: 6},
	},
	DeathCauseStatus: map[model.DeathCause]int{
		model.MOD_ROCKET:        37,
		model.MOD_TRIGGER_HURT:  14,
		model.MOD_RAILGUN:       9,
		model.MOD_ROCKET_SPLASH: 60,
		model.MOD_MACHINEGUN:    4,
		model.MOD_SHOTGUN:       4,
		model.MOD_FALLING:       3,
	},
}

// MatchLogFile2Fixture fixture for log_2.log file
var MatchLogFile2Fixture = model.Match{
	ID:         1,
	TotalKills: 0,
	Players:    []model.Player{"Isgalamido"},
	Kills: map[model.Player]int{
		"Isgalamido": 0,
	},
	PlayerRanking: []model.Ranking{
		{Position: 1, Player: "Isgalamido", TotalKills: 0},
	},
	DeathCauseStatus: map[model.DeathCause]int{},
}

// MatchLogFile3Fixture1 fixture for log_3.log file
var MatchLogFile3Fixture1 = model.Match{
	ID:         1,
	TotalKills: 0,
	Players:    []model.Player{"Isgalamido"},
	Kills: map[model.Player]int{
		"Isgalamido": 0,
	},
	PlayerRanking: []model.Ranking{
		{Position: 1, Player: "Isgalamido", TotalKills: 0},
	},
	DeathCauseStatus: map[model.DeathCause]int{},
}

// MatchLogFile3Fixture2 fixture for log_3.log file
var MatchLogFile3Fixture2 = model.Match{
	ID:         2,
	TotalKills: 131,
	Players:    []model.Player{"Isgalamido", "Oootsimo", "Dono da Bola", "Assasinu Credi", "Zeh", "Mal"},
	Kills: map[model.Player]int{
		"Isgalamido":     17,
		"Oootsimo":       20,
		"Dono da Bola":   10,
		"Assasinu Credi": 13,
		"Zeh":            19,
		"Mal":            6,
	},
	PlayerRanking: []model.Ranking{
		{Position: 1, Player: "Oootsimo", TotalKills: 20},
		{Position: 2, Player: "Zeh", TotalKills: 19},
		{Position: 3, Player: "Isgalamido", TotalKills: 17},
		{Position: 4, Player: "Assasinu Credi", TotalKills: 13},
		{Position: 5, Player: "Dono da Bola", TotalKills: 10},
		{Position: 6, Player: "Mal", TotalKills: 6},
	},
	DeathCauseStatus: map[model.DeathCause]int{
		model.MOD_ROCKET:        37,
		model.MOD_TRIGGER_HURT:  14,
		model.MOD_RAILGUN:       9,
		model.MOD_ROCKET_SPLASH: 60,
		model.MOD_MACHINEGUN:    4,
		model.MOD_SHOTGUN:       4,
		model.MOD_FALLING:       3,
	},
}

// MatchLogFile4Fixture fixture for log_4.log file
var MatchLogFile4Fixture = model.Match{
	ID:         1,
	TotalKills: 14,
	Players:    []model.Player{"Dono da Bola", "Isgalamido", "Zeh", "Assasinu Credi"},
	Kills: map[model.Player]int{
		"Assasinu Credi": -3,
		"Dono da Bola":   0,
		"Isgalamido":     2,
		"Zeh":            1},
	PlayerRanking: []model.Ranking{
		{Position: 1, Player: "Isgalamido", TotalKills: 2},
		{Position: 2, Player: "Zeh", TotalKills: 1},
		{Position: 3, Player: "Dono da Bola", TotalKills: 0},
		{Position: 4, Player: "Assasinu Credi", TotalKills: -3},
	},
	DeathCauseStatus: map[model.DeathCause]int{
		model.MOD_RAILGUN:       1,
		model.MOD_ROCKET:        4,
		model.MOD_ROCKET_SPLASH: 4,
		model.MOD_TRIGGER_HURT:  5,
	},
}
