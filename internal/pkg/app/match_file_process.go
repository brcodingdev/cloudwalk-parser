package app

import (
	"bufio"
	"fmt"
	"github.com/brcodingdev/cloudwalk-parser/internal/pkg/model"
	"github.com/brcodingdev/cloudwalk-parser/internal/ports/repository"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
)

// regexp to find pattern for player
var regexpPlayer = regexp.MustCompile(`n\\(.*?)\\t`)

// regexp to find pattern for kills
var regexpKills = regexp.MustCompile(`(.+):\s+(.+) killed (.+) by (\w+)`)

// MatchFileProcess structure to load matches in the file
type MatchFileProcess struct {
	absFileLog string
	repository repository.Match
}

// NewMatchFileProcess creates new instance of match file process
func NewMatchFileProcess(
	absFileLog string,
	repository repository.Match,
) *MatchFileProcess {
	return &MatchFileProcess{
		absFileLog: absFileLog,
		repository: repository,
	}
}

// Load starts the process to read file and save matches in the database
func (a *MatchFileProcess) Load() error {
	f, err := os.Open(a.absFileLog)
	if err != nil {
		return err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	rowsMatch := make([]string, 0)
	var wg sync.WaitGroup
	var matchID int32

	for {
		scanned := scanner.Scan()
		row := scanner.Text()
		// header str, start and finnish match
		if shouldSkipRow(&row) {
			continue
		}

		if len(rowsMatch) > 0 &&
			(shouldInitNewMatch(&row) || !scanned) {
			wg.Add(1)
			// process the copied rows
			go func(rows []string) {
				defer wg.Done()

				a.ProcessMatch(&rows,
					atomic.AddInt32(&matchID, 1))
			}(rowsMatch)

			// execute concurrently and sequentially
			wg.Wait()

			// cleaning buffer after call goroutine
			rowsMatch = make([]string, 0)
		}

		rowsMatch = append(rowsMatch, row)

		if !scanned {
			break
		}
	}
	return nil
}

// ProcessMatch process matches with rows grouped
func (a *MatchFileProcess) ProcessMatch(
	rows *[]string,
	matchID int32,
) *model.Match {
	match := model.NewMatch(matchID)

	for _, row := range *rows {
		if shouldFillPlayer(&row) {
			fillPlayer(&row, match)
		}

		if shouldFillKills(&row) {
			match.TotalKills++
			fillKills(&row, match)
		}
	}

	fillRanking(match)
	err := a.repository.Add(match)
	if err != nil {
		fmt.Printf("could not save match err: %s\nset HOST_REDIS env var\n", err)
		return nil
	}

	fmt.Println("match id ", match.ID, "saved successfully")
	return match
}

func fillPlayer(row *string, match *model.Match) {
	// 20:38 ClientUserinfoChanged: 2 n\Isgalamido\t\0\model\...
	// regex group
	// 1. (.*?) Username Isgalamido
	matches := regexpPlayer.FindStringSubmatch(*row)
	if len(matches) < 2 {
		return
	}

	player := model.Player(matches[1])
	hasFound := false
	for _, p := range match.Players {
		if p == player {
			hasFound = true
			break
		}
	}
	if !hasFound {
		match.Players = append(match.Players, player)
		match.Kills[player] = 0
	}
}

func fillKills(row *string, match *model.Match) {
	// 21:42 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT
	// regex group
	// 1. (.+): 21:42 Kill: 1022 2 22
	// 2. \s+(.+)     killer:     <world>
	// 3. (.+)        killed by:  Isgalamido
	// 4. (\w+)       kill means: MOD_TRIGGER_HURT
	matches := regexpKills.FindStringSubmatch(*row)
	if len(matches) < 4 {
		return
	}

	killer := model.Player(matches[2])
	killedBy := model.Player(matches[3])
	deathCause := model.DeathCause(matches[4])

	// if kill by means is not mapped, then it should be unknown
	if _, ok := model.MappedDeathCause[deathCause]; !ok {
		match.DeathCauseStatus[model.MOD_UNKNOWN]++
	} else {
		match.DeathCauseStatus[deathCause]++
	}

	// analysing the match clearly suicide have
	// the same behavior as <world> kills
	// if not subtract kills the score will not match
	// 13:55 score: 20  ping: 8  client: 3 Oootsimo
	// 13:55 score: 19  ping: 14  client: 6 Zeh
	// 13:55 score: 17  ping: 1  client: 2 Isgalamido
	// 13:55 score: 13  ping: 0  client: 5 Assasinu Credi
	// 13:55 score: 10  ping: 8  client: 4 Dono da Bola
	// 13:55 score: 6  ping: 19  client: 7 Mal
	if killer == "<world>" || killer == killedBy {
		match.Kills[killedBy]--

	} else {
		match.Kills[killer]++
	}
}

func fillRanking(match *model.Match) {
	for player, kills := range match.Kills {
		match.PlayerRanking = append(match.PlayerRanking,
			model.Ranking{
				Player:     player,
				TotalKills: kills,
			})
	}

	sort.Slice(match.PlayerRanking, func(i, j int) bool {
		return match.PlayerRanking[i].TotalKills > match.PlayerRanking[j].TotalKills
	})

	for i := range match.PlayerRanking {
		match.PlayerRanking[i].Position = i + 1
	}
}

func shouldSkipRow(row *string) bool {
	return strings.LastIndex(
		*row,
		"------------------------------------------------------------") > -1
}

func shouldFillPlayer(row *string) bool {
	return strings.LastIndex(*row, "ClientUserinfoChanged: ") > -1
}

func shouldFillKills(row *string) bool {
	return strings.LastIndex(*row, "Kill: ") > -1
}

func shouldInitNewMatch(row *string) bool {
	return strings.LastIndex(*row, "InitGame: ") > -1
}
