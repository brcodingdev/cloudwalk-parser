package app_test

import (
	"encoding/json"
	"github.com/brcodingdev/cloudwalk-parser/internal/pkg/app"
	"github.com/brcodingdev/cloudwalk-parser/internal/pkg/model"
	"github.com/brcodingdev/cloudwalk-parser/internal/ports/repository/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrintMatch(t *testing.T) {

	cases := map[string]struct {
		matches  []model.Match
		expected string
	}{
		"print_report": {
			matches: []model.Match{
				mocks.MatchLogFile2Fixture,
				mocks.MatchLogFile3Fixture2,
			},
			expected: `[{"game_1":{"id":1,"kills":{"Isgalamido":0},"kills_by_means":{},"players":["Isgalamido"],"ranking":[{"player":"Isgalamido","position":1,"total_kills":0}],"total_kills":0}},{"game_2":{"id":2,"kills":{"Assasinu Credi":13,"Dono da Bola":10,"Isgalamido":17,"Mal":6,"Oootsimo":20,"Zeh":19},"kills_by_means":{"MOD_FALLING":3,"MOD_MACHINEGUN":4,"MOD_RAILGUN":9,"MOD_ROCKET":37,"MOD_ROCKET_SPLASH":60,"MOD_SHOTGUN":4,"MOD_TRIGGER_HURT":14},"players":["Isgalamido","Oootsimo","Dono da Bola","Assasinu Credi","Zeh","Mal"],"ranking":[{"player":"Oootsimo","position":1,"total_kills":20},{"player":"Zeh","position":2,"total_kills":19},{"player":"Isgalamido","position":3,"total_kills":17},{"player":"Assasinu Credi","position":4,"total_kills":13},{"player":"Dono da Bola","position":5,"total_kills":10},{"player":"Mal","position":6,"total_kills":6}],"total_kills":131}}]`,
		},
	}

	for caseTitle, tc := range cases {
		t.Run(caseTitle, func(t *testing.T) {
			repository := mocks.Match{}
			repository.On("FindAll").Return(tc.matches, nil)

			appReport := app.NewMatchReport(&repository)

			report, err := appReport.PrintMatch()
			assert.NoError(t, err)

			var obj []interface{}
			// remove indentation
			err = json.Unmarshal([]byte(report), &obj)
			assert.NoError(t, err)

			bs, err := json.Marshal(obj)
			assert.NoError(t, err)

			assert.Equal(t, tc.expected, string(bs))
		})
	}
}
