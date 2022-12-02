package two

import (
	"bytes"
	"embed"
)

type Score int

const (
	Rock     Score = 1
	Paper    Score = 2
	Scissors Score = 3
	// Outcomes
	Lost Score = 0
	Draw Score = 3
	Won  Score = 6
)

//go:embed *.txt
var inputData embed.FS

type Play [3]Score

func (p Play) PlayerScore() int {
	return int(p[1] + p[2])
}

// This is so ugly and dumb
var outcomes = map[Score]map[Score]Score{
	Rock: {
		Rock:     Draw,
		Paper:    Won,
		Scissors: Lost,
	},
	Paper: {
		Rock:     Lost,
		Paper:    Draw,
		Scissors: Won,
	},
	Scissors: {
		Rock:     Won,
		Paper:    Lost,
		Scissors: Draw,
	},
}

var strategy = map[Score]map[Score]Score{
	Rock: {
		Won:  Paper,
		Lost: Scissors,
		Draw: Rock,
	},
	Paper: {
		Won:  Scissors,
		Lost: Rock,
		Draw: Paper,
	},
	Scissors: {
		Won:  Rock,
		Lost: Paper,
		Draw: Scissors,
	},
}

var plays = map[byte]Score{
	'A': Rock,
	'B': Paper,
	'C': Scissors,
	'X': Rock,
	'Y': Paper,
	'Z': Scissors,
}

var playsPartTwo = map[byte]Score{
	'A': Rock,
	'B': Paper,
	'C': Scissors,
	'X': Lost,
	'Y': Draw,
	'Z': Won,
}

func ParseInput() []Play {
	data, err := inputData.ReadFile("input.txt")
	if err != nil {
		panic("Failed to read the input data!")
	}

	rows := bytes.Split(data, []byte("\n"))
	results := make([]Play, 0, len(rows))
	for _, row := range rows {
		if len(row) == 0 {
			continue
		}
		opponent := plays[row[0]]
		player := plays[row[len(row)-1]]
		results = append(results, Play{
			opponent,
			player,
			outcomes[opponent][player],
		})
	}

	return results
}

func ParseInputPartTwo() []Play {
	data, err := inputData.ReadFile("input.txt")
	if err != nil {
		panic("Failed to read the input data!")
	}

	rows := bytes.Split(data, []byte("\n"))
	results := make([]Play, 0, len(rows))
	for _, row := range rows {
		if len(row) == 0 {
			continue
		}
		opponent := playsPartTwo[row[0]]
		player := playsPartTwo[row[len(row)-1]]
		playerMove := strategy[opponent][player]
		results = append(results, Play{
			opponent,
			playerMove,
			outcomes[opponent][playerMove],
		})
	}

	return results
}

func TotalScore(all []Play) int {
	var total int
	for _, play := range all {
		total += play.PlayerScore()
	}
	return total
}
