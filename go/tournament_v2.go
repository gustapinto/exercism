package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
	"errors"
)

type Score struct {
	Team          string
	MatchesPlayed uint
	MatchesWon    uint
	MatchesDrawn  uint
	MatchesLost   uint
	Points        uint
}

type Scores map[string]*Score

var (
	ValidResults = map[string]bool{
		"win": true,
		"loss": true,
		"draw": true,
	}

	ScorePerTeam = map[string]*Score{}
)

func Tally(reader io.Reader, writer io.Writer) error {
	scorePerTeam := make(Scores)

	lines, err := GetLines(reader)
	if err != nil {
		return err
	}

	// Initialize map
	for _, line := range lines {
		teamName, oposingTeamName := line[0], line[1]

		if _, teamExists := scorePerTeam[teamName]; !teamExists {
			scorePerTeam[teamName] = &Score{Team: teamName}
		}
		if _, oposingTeamExists := scorePerTeam[oposingTeamName]; !oposingTeamExists {
			scorePerTeam[oposingTeamName] = &Score{Team: oposingTeamName}
		}
	}

	// Build the score
	for _, line := range lines {
		ParseLineData(line, scorePerTeam)
	}

	sortedScore := SortScores(scorePerTeam)
	output := BuildTable(sortedScore)

	if _, err := writer.Write([]byte(output)); err != nil {
		return err
	}

	return nil
}

func GetLines(reader io.Reader) ([][]string, error) {
	var lines [][]string

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if string(line[0]) == "#" {
			continue
		}

		data := strings.Split(line, ";")
		if len(data) < 3 {
			return nil, errors.New("Invalid input")
		}
		if _, exists := ValidResults[data[2]]; !exists {
			return nil, errors.New("Invalid result")
		}

		lines = append(lines, data)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func ParseLineData(line []string, scores Scores) {
	teamName, oposingTeamName, result := line[0], line[1], line[2]
	teamScore := scores[teamName]
	oposingTeamScore := scores[oposingTeamName]

	switch result {
	case "win":
		teamScore.MatchesWon += 1
		teamScore.Points += 3
		oposingTeamScore.MatchesLost += 1
	case "loss":
		teamScore.MatchesLost += 1
		oposingTeamScore.MatchesWon += 1
		oposingTeamScore.Points += 3
	case "draw":
		teamScore.MatchesDrawn += 1
		teamScore.Points += 1
		oposingTeamScore.MatchesDrawn += 1
		oposingTeamScore.Points += 1
	}

	teamScore.MatchesPlayed += 1
	oposingTeamScore.MatchesPlayed += 1
}

func SortScores(scores Scores) []Score {
	var scoreList []Score
	for _, score := range scores {
		scoreList = append(scoreList, *score)
	}

	sort.Slice(scoreList, func(i, j int) bool {
		if scoreList[i].Points == scoreList[j].Points {
			return scoreList[i].Team < scoreList[j].Team
		}

		return scoreList[i].Points > scoreList[j].Points
	})

	return scoreList
}

func BuildTable(scores []Score) string {
	output := "Team                           | MP |  W |  D |  L |  P\n"

	for _, score := range scores {
		for len(score.Team) < 30 {
			score.Team += " "
		}

		output += fmt.Sprintf("%s |  %d |  %d |  %d |  %d |  %d\n", score.Team,
			score.MatchesPlayed, score.MatchesWon, score.MatchesDrawn,
			score.MatchesLost, score.Points)
	}

	return output
}
