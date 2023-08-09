package app_test

import (
	"bufio"
	"github.com/brcodingdev/cloudwalk-parser/internal/pkg/app"
	"github.com/brcodingdev/cloudwalk-parser/internal/pkg/model"
	"github.com/brcodingdev/cloudwalk-parser/internal/ports/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestOneMatchFileProcessLoad(t *testing.T) {
	cases := map[string]struct {
		file          string
		matchExpected model.Match
	}{
		"load_one_single_match": {
			file:          "assets/log_5.log",
			matchExpected: mocks.MatchLogFile5Fixture,
		},

		"load_match_without_kills": {
			file:          "assets/log_2.log",
			matchExpected: mocks.MatchLogFile2Fixture,
		},
	}

	for caseTitle, tc := range cases {
		t.Run(caseTitle, func(t *testing.T) {
			repository := mocks.Match{}
			repository.On(
				"Add", &tc.matchExpected).
				Return(nil).Times(1)

			filePath := getFilePath(tc.file)

			appMatch := app.NewMatchFileProcess(filePath, &repository)
			err := appMatch.Load()
			assert.NoError(t, err)

			require.Eventually(t, func() bool {
				return repository.AssertExpectations(t)
			}, 2*time.Second, 50*time.Millisecond)
		})
	}
}

func TestMoreMatchesFileProcessLoad(t *testing.T) {
	cases := map[string]struct {
		file                    string
		callFirstMatchExpected  model.Match
		callSecondMatchExpected model.Match
	}{
		"load_more_matches": {
			file:                    "assets/log_3.log",
			callFirstMatchExpected:  mocks.MatchLogFile3Fixture1,
			callSecondMatchExpected: mocks.MatchLogFile3Fixture2,
		},
	}

	for caseTitle, tc := range cases {
		t.Run(caseTitle, func(t *testing.T) {
			repository := mocks.Match{}
			repository.On(
				"Add", &tc.callFirstMatchExpected).
				Return(nil).Times(1).
				On("Add", &tc.callSecondMatchExpected).
				Return(nil).Times(1)

			filePath := getFilePath(tc.file)

			appMatch := app.NewMatchFileProcess(filePath, &repository)
			err := appMatch.Load()
			assert.NoError(t, err)

			require.Eventually(t, func() bool {
				return repository.AssertExpectations(t)
			}, 2*time.Second, 50*time.Millisecond)
		})
	}
}

func TestMatchFileProcess(t *testing.T) {
	cases := map[string]struct {
		file          string
		matchID       int32
		matchExpected model.Match
	}{
		"process_file": {
			matchID:       1,
			file:          "assets/log_4.log",
			matchExpected: mocks.MatchLogFile4Fixture,
		},
	}

	for caseTitle, tc := range cases {
		t.Run(caseTitle, func(t *testing.T) {
			repository := mocks.Match{}
			repository.On(
				"Add", mock.Anything).
				Return(nil).Times(1)

			filePath := getFilePath(tc.file)

			f, err := os.Open(filePath)
			assert.NoError(t, err)

			scanner := bufio.NewScanner(f)
			rows := make([]string, 0)
			for scanner.Scan() {
				rows = append(rows, scanner.Text())
			}

			appMatch := app.NewMatchFileProcess(filePath, &repository)
			match := appMatch.ProcessMatch(&rows, tc.matchID)
			assert.Equal(t, &tc.matchExpected, match)
		})
	}
}

func getFilePath(file string) string {
	testDir, _ := filepath.Abs("../../../")
	return filepath.Join(testDir, file)
}
