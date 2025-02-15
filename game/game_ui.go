package game

import (
	"bufio"
	"fmt"
	"os"
)

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
