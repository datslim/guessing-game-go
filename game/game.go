package game

import (
	"time"
)

const (
	continuePlaying   = "1"
	stopPlaying       = "2"
	EASY_DIFFICULTY   = 1
	MEDIUM_DIFFICULTY = 2
	HARD_DIFFICULTY   = 3
	invalidInputMsg   = "Invalid input! Please enter a valid number."
)

type GameInfo struct {
	TotalAttempts int
	Answer        int
	StartTime     time.Time
	MinimumRange  int
	MaximumRange  int
}

func StartGame() {
	cheering()
	currentGameSession := initializeGame()
	for {
		currentGameSession.TotalAttempts = difficultyChoosing()
		getRandomNumber(&currentGameSession)
		playGame(&currentGameSession)
		if askToPlayAgain() == stopPlaying {
			break
		}
	}
	goodBye()
}

func initializeGame() GameInfo {
	return GameInfo{
		TotalAttempts: 0,
		Answer:        0,
		StartTime:     time.Time{},
		MinimumRange:  0,
		MaximumRange:  100,
	}
}
