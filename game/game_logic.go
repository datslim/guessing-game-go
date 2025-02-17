package game

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandomNumber(game *GameInfo) {
	game.Answer = rand.Intn(99) + 1
	game.MinimumRange = 0
	game.MaximumRange = 100
}

func getElapsedTime(startTime time.Time) time.Duration {
	return time.Since(startTime)
}

func formatTime(d time.Duration) string {
	minutes := int(d.Minutes())
	seconds := int(d.Seconds()) % 60
	return fmt.Sprintf("%01d minutes and %01d seconds", minutes, seconds)
}

func playGame(game *GameInfo) {
	inputsList := make([]int, game.TotalAttempts)
	game.StartTime = time.Now()
	for attempt := 0; attempt < game.TotalAttempts; attempt++ {
		input := getUserGuess(inputsList)
		inputsList[attempt] = input

		clearScreen()
		printLogo()

		if input == game.Answer {
			gameDuration := getElapsedTime(game.StartTime)
			formatedTime := formatTime(gameDuration)
			fmt.Fprintf(out, "%vCongratulations! 🏆 You guessed the correct number in %v attempts and %v.\n%v", greenColor, attempt+1, formatedTime, resetColor)
			return
		} else {
			fmt.Fprintf(out, "%vYour guess was: %d%v", whiteColor, input, resetColor)
			provideHint(game, input)
		}

		remainingAttempts := game.TotalAttempts - attempt - 1

		if remainingAttempts != 0 {
			fmt.Fprintf(out, "You have %v%d attempts%v left!\n\n", yellowColor, remainingAttempts, resetColor)
		}

		if remainingAttempts == 1 {
			oddsOrEvenHint(game)
		}

	}
	fmt.Fprintf(out, "%vYou're out of chances.%v %v Nice try! The number to guess was %v%v%d%v\n\n", redColor, resetColor, whiteColor, resetColor, greenColor, game.Answer, resetColor)
}

func updateRange(min, max, input, answer int) (int, int) {
	if input < answer {
		if input > min {
			min = input
		}
	} else if input > answer {
		if input < max {
			max = input
		}
	}
	return min, max
}
