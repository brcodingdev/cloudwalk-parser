package app

// MatchProcess contract to load matches
type MatchProcess interface {
	// Load to start process of read and store matches
	Load() error
}
