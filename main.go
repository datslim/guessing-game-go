package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	continuePlaying   = "1"
	stopPlaying       = "2"
	EASY_DIFFICULTY   = 1
	MEDIUM_DIFFICULTY = 2
	HARD_DIFFICULTY   = 3
	resetColor        = "\033[0m"
	redColor          = "\033[31;1m"
	greenColor        = "\033[32;1m"
	cyanColor         = "\033[36;1m"
	blueColor         = "\033[34;1m"
	whiteColor        = "\033[97;1m"
	yellowColor       = "\033[33;1m"
	magentaColor      = "\033[35;1m"
	invalidInputMsg   = "Invalid input! Please enter a valid number."
)

type GameInfo struct {
	TotalAttempts int
	Answer        int
	StartTime     time.Time
	MinimumRange  int
	MaximumRange  int
}

func main() {
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

func cheering() {
	printLogo()
	fmt.Printf("%vWelcome to the%v %vNumber Guessing Game!%v\n%vI'm thinking of a number between%v %v1 and 100.%v\n",
		whiteColor, resetColor, cyanColor, resetColor,
		whiteColor, resetColor, greenColor, resetColor)
	fmt.Printf("%vYour%v %vnumber of attempts%v %vto guess the number%v %vdepends on the difficulty%v %vlevel, so...%v\n\n",
		whiteColor, resetColor, blueColor, resetColor, whiteColor,
		resetColor, blueColor, resetColor, whiteColor, resetColor)

}

func printLogo() {
	art := `
  _______ _______ _______ _______ _______        _______ _______ _______ _______ 
|     __|   |   |    ___|     __|     __|______|     __|   _   |   |   |    ___|
|    |  |   |   |    ___|__     |__     |______|    |  |       |       |    ___|
|_______|_______|_______|_______|_______|      |_______|___|___|__|_|__|_______|
                                                                                `
	fmt.Printf("%v%v\n%v", cyanColor, art, resetColor)
	fmt.Printf("\t\t\t\t%vMade by github.com/datslim%v\n\n", redColor, resetColor)
}

func goodBye() {
	clearScreen()
	printLogo()
	catArt := `  /\_/\  (
 ( ^.^ ) _)
   \"/  (
 ( | | )
(__d b__)`
	fmt.Printf("%vSee you later! Hope you had a good time here!%v\n", magentaColor, resetColor)
	fmt.Printf("%v%v%v", cyanColor, catArt, resetColor)
}

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

func difficultyChoosing() int {
	var difficultyChoice, totalAttempts int
	var difficultyName, keyColor string

	for {
		fmt.Printf("%vSelect the difficulty level:%v\n%v1. Easy (10 chances)%v\n%v2. Medium (5 chances)%v\n%v3. Hard (3 chances)%v\n\n",
			whiteColor, resetColor, greenColor, resetColor,
			yellowColor, resetColor, redColor, resetColor)

		fmt.Printf("%vEnter your choice: %v", whiteColor, resetColor)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		if n, err := fmt.Sscanf(scanner.Text(), "%d", &difficultyChoice); err != nil || n == 0 {
			fmt.Printf("%v%s%v\n", redColor, invalidInputMsg, resetColor)
			continue //
		}

		if difficultyChoice < 1 || difficultyChoice > 3 {
			fmt.Printf("\n%vPlease enter 1, 2 or 3.%v\n", redColor, resetColor)
			continue //
		}

		break
	}

	switch difficultyChoice {
	case EASY_DIFFICULTY:
		totalAttempts = 10
		difficultyName = "Easy"
		keyColor = greenColor

	case MEDIUM_DIFFICULTY:
		totalAttempts = 5
		difficultyName = "Medium"
		keyColor = yellowColor

	case HARD_DIFFICULTY:
		totalAttempts = 3
		difficultyName = "Hard"
		keyColor = redColor
	}

	fmt.Printf("\n%vGreat!%v %vYou have selected the %v%v%s%v %vdifficulty level.%v\n\n",
		greenColor, resetColor, whiteColor, resetColor,
		keyColor, difficultyName, resetColor, whiteColor, resetColor)

	return totalAttempts
}

func getUserGuess(inputsList []int) int {
	var input int
	for {
		fmt.Printf("%vEnter your guess: %v", blueColor, resetColor)

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		if n, err := fmt.Sscanf(scanner.Text(), "%d", &input); err != nil || n == 0 {
			fmt.Printf("%v%s%v\n", redColor, invalidInputMsg, resetColor)
			continue
		}

		if input < 1 || input > 99 {
			fmt.Printf("%vPlease enter a number%v %vbetween 1 and 100.%v\n", whiteColor, resetColor, redColor, resetColor)
			continue
		}

		if numberAlreadyTried(inputsList, input) {
			fmt.Printf("%vYou've already tried this number!%v\n", redColor, resetColor)
			continue
		}

		break
	}

	return input
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
			fmt.Printf("%vCongratulations! You guessed the correct number in %v attempts and %v.\n%v", greenColor, attempt+1, formatedTime, resetColor)
			return
		} else {
			fmt.Printf("%vYour guess was: %d%v", whiteColor, input, resetColor)
			provideHint(game, input)
		}

		remainingAttempts := game.TotalAttempts - attempt - 1

		if remainingAttempts != 0 {
			fmt.Printf("You have %v%d attempts%v left!\n\n", yellowColor, remainingAttempts, resetColor)
		}

		if remainingAttempts == 1 {
			oddsOrEvenHint(game)
		}

	}
	fmt.Printf("%vYou're out of chances.%v %v Nice try! The number to guess was %v%v%d%v\n\n", redColor, resetColor, whiteColor, resetColor, greenColor, game.Answer, resetColor)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func askToPlayAgain() string {
	fmt.Printf("%vDo you want to play another round?%v\n%v1. Yes%v\n%v2. No%v\n\n", cyanColor, resetColor, greenColor, resetColor, redColor, resetColor)
	fmt.Printf("%vEnter your choice: %v", whiteColor, resetColor)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Text() != continuePlaying && scanner.Text() != stopPlaying {
		fmt.Printf("%vPlease enter 1 or 2.%v\n", redColor, resetColor)
		return askToPlayAgain()
	}
	return scanner.Text()
}

func oddsOrEvenHint(game *GameInfo) {
	if game.Answer%2 == 0 {
		fmt.Printf("%vHint:%v the number is %veven.%v\n", greenColor, resetColor, greenColor, resetColor)
	} else {
		fmt.Printf("%vHint:%v the number is %vodd.%v\n", greenColor, redColor, greenColor, resetColor)
	}
}

func provideHint(game *GameInfo, input int) {
	game.MinimumRange, game.MaximumRange = updateRange(game.MinimumRange, game.MaximumRange, input, game.Answer)
	if input < game.Answer {
		fmt.Printf("\n%vIncorrect!%v The number is %vgreater%v than %v.\n", redColor, resetColor, cyanColor, resetColor, input)
	} else if input > game.Answer {
		fmt.Printf("\n%vIncorrect!%v The number is %vless%v than %v.\n", redColor, resetColor, cyanColor, resetColor, input)
	}

	printRange(game.MinimumRange, game.MaximumRange)

	if abs(input-game.Answer) <= 3 {
		fmt.Printf("%vGetting hot! You're close!%v\n", magentaColor, resetColor)
	}
}

func numberAlreadyTried(inputsList []int, input int) bool {
	for _, listValue := range inputsList {
		if listValue == input {
			return true
		}
	}
	return false
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

func printRange(min, max int) {
	rangeString := "✦━━━━━━━━━━━━━━━━━━━━━━✦"
	fmt.Printf("%v %d %v %v %s %v %v %d%v\n", blueColor, min, resetColor, greenColor, rangeString, resetColor, blueColor, max, resetColor)
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
