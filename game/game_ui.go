package game

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mattn/go-colorable"
)

var (
	out = colorable.NewColorableStdout()
)

func printLogo() {
	art := `
  _______ _______ _______ _______ _______        _______ _______ _______ _______ 
|     __|   |   |    ___|     __|     __|______|     __|   _   |   |   |    ___|
|    |  |   |   |    ___|__     |__     |______|    |  |       |       |    ___|
|_______|_______|_______|_______|_______|      |_______|___|___|__|_|__|_______|
                                                                                `
	fmt.Fprintf(out, "%v%v\n%v", cyanColor, art, resetColor)
	fmt.Fprintf(out, "\t\t\t\t%vMade by github.com/datslim%v\n\n", redColor, resetColor)
}

func cheering() {
	printLogo()
	fmt.Fprintf(out, "%vWelcome to the%v %vNumber Guessing Game!%v\n%vI'm thinking of a number between%v %v1 and 100.%v\n",
		whiteColor, resetColor, cyanColor, resetColor,
		whiteColor, resetColor, greenColor, resetColor)
	fmt.Fprintf(out, "%vYour%v %vnumber of attempts%v %vto guess the number%v %vdepends on the difficulty%v %vlevel, so...%v\n\n",
		whiteColor, resetColor, blueColor, resetColor, whiteColor,
		resetColor, blueColor, resetColor, whiteColor, resetColor)

}

func difficultyChoosing() int {
	var difficultyChoice, totalAttempts int
	var difficultyName, keyColor string

	for {
		fmt.Fprintf(out, "%vSelect the difficulty level:%v\n%v1. Easy (10 chances)%v\n%v2. Medium (5 chances)%v\n%v3. Hard (3 chances)%v\n\n",
			whiteColor, resetColor, greenColor, resetColor,
			yellowColor, resetColor, redColor, resetColor)

		fmt.Fprintf(out, "%vEnter your choice: %v", whiteColor, resetColor)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		if n, err := fmt.Sscanf(scanner.Text(), "%d", &difficultyChoice); err != nil || n == 0 {
			fmt.Fprintf(out, "%v%s%v\n", redColor, invalidInputMsg, resetColor)
			continue //
		}

		if difficultyChoice < 1 || difficultyChoice > 3 {
			fmt.Fprintf(out, "\n%vPlease enter 1, 2 or 3.%v\n", redColor, resetColor)
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

	fmt.Fprintf(out, "\n%vGreat!%v %vYou have selected the %v%v%s%v %vdifficulty level.%v\n\n",
		greenColor, resetColor, whiteColor, resetColor,
		keyColor, difficultyName, resetColor, whiteColor, resetColor)

	return totalAttempts
}

func getUserGuess(inputsList []int) int {
	var input int
	for {
		fmt.Fprintf(out, "%vEnter your guess: %v", blueColor, resetColor)

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		if n, err := fmt.Sscanf(scanner.Text(), "%d", &input); err != nil || n == 0 {
			fmt.Fprintf(out, "%v%s%v\n", redColor, invalidInputMsg, resetColor)
			continue
		}

		if input < 1 || input > 99 {
			fmt.Fprintf(out, "%vPlease enter a number%v %vbetween 1 and 100.%v\n", whiteColor, resetColor, redColor, resetColor)
			continue
		}

		if numberAlreadyTried(inputsList, input) {
			fmt.Fprintf(out, "%vYou've already tried this number!%v\n", redColor, resetColor)
			continue
		}

		break
	}

	return input
}

func askToPlayAgain() string {
	fmt.Fprintf(out, "%vDo you want to play another round?%v\n%v1. Yes%v\n%v2. No%v\n\n", cyanColor, resetColor, greenColor, resetColor, redColor, resetColor)
	fmt.Fprintf(out, "%vEnter your choice: %v", whiteColor, resetColor)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Text() != continuePlaying && scanner.Text() != stopPlaying {
		fmt.Fprintf(out, "%vPlease enter 1 or 2.%v\n", redColor, resetColor)
		return askToPlayAgain()
	}
	return scanner.Text()
}

func goodBye() {
	clearScreen()
	printLogo()
	catArt := `  /\_/\  (
 ( ^.^ ) _)
   \"/  (
 ( | | )
(__d b__)`
	fmt.Fprintf(out, "%vSee you later! Hope you had a good time here!%v\n", magentaColor, resetColor)
	fmt.Fprintf(out, "%v%v%v", cyanColor, catArt, resetColor)
}
