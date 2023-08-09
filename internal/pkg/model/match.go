package model

// Player player name
type Player string

// Match structure to store info of matches
type Match struct {
	ID               int32              `json:"id"`
	TotalKills       int                `json:"total_kills"`
	Players          []Player           `json:"players"`
	Kills            map[Player]int     `json:"kills"`
	PlayerRanking    []Ranking          `json:"ranking"`
	DeathCauseStatus map[DeathCause]int `json:"kills_by_means"`
}

// Ranking player ranking
type Ranking struct {
	Position   int    `json:"position"`
	Player     Player `json:"player"`
	TotalKills int    `json:"total_kills"`
}

// NewMatch new match
func NewMatch(id int32) *Match {
	return &Match{
		ID:               id,
		TotalKills:       0,
		Players:          make([]Player, 0),
		Kills:            make(map[Player]int),
		PlayerRanking:    make([]Ranking, 0),
		DeathCauseStatus: make(map[DeathCause]int),
	}
}
