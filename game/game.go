package game

import (
	"time"
)

type GameInfo struct {
	TotalAttempts int
	Answer        int
	StartTime     time.Time
	MinimumRange  int
	MaximumRange  int
}

func StartGame() {
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
	cheering()
	return GameInfo{
		TotalAttempts: 0,
		Answer:        0,
		StartTime:     time.Time{},
		MinimumRange:  0,
		MaximumRange:  100,
	}
}
