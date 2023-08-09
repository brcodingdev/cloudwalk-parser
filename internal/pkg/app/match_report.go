package app

import (
	"encoding/json"
	"fmt"
	"github.com/brcodingdev/cloudwalk-parser/internal/pkg/model"
	"github.com/brcodingdev/cloudwalk-parser/internal/ports/repository"
	"sort"
)

// MatchReport structure of matches report
type MatchReport struct {
	repository repository.Match
}

// NewMatchReport creates new instance of matches report
func NewMatchReport(
	repository repository.Match,
) *MatchReport {
	return &MatchReport{
		repository: repository,
	}
}

// PrintMatch prints report in json in the console
func (a *MatchReport) PrintMatch() (string, error) {
	matches, err := a.repository.FindAll()

	if err != nil {
		return "", err
	}
	//sort by ID
	sort.Slice(matches, func(i, j int) bool {
		return matches[i].ID < matches[j].ID
	})
	// convert to format expected
	resultList := make([]map[string]model.Match, 0, len(matches))

	for _, match := range matches {
		result := map[string]model.Match{}
		// the ID is not expected in the report
		id := match.ID
		result[fmt.Sprintf("game_%d", id)] = match
		resultList = append(resultList, result)
	}

	resultJSON, err := json.MarshalIndent(resultList, "", "  ")
	if err != nil {
		fmt.Println("could not marshall report to json", err)
		return "", err
	}

	report := string(resultJSON)
	fmt.Println(report)
	return report, nil
}
