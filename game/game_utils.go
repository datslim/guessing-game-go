package game

import (
	"fmt"
)

func provideHint(game *GameInfo, input int) {
	game.MinimumRange, game.MaximumRange = updateRange(game.MinimumRange, game.MaximumRange, input, game.Answer)
	if input < game.Answer {
		fmt.Fprintf(out, "\n%vIncorrect!%v The number is %vgreater%v than %v.\n", redColor, resetColor, cyanColor, resetColor, input)
	} else if input > game.Answer {
		fmt.Fprintf(out, "\n%vIncorrect!%v The number is %vless%v than %v.\n", redColor, resetColor, cyanColor, resetColor, input)
	}

	printRange(game.MinimumRange, game.MaximumRange)

	if abs(input-game.Answer) <= 3 {
		fmt.Fprintf(out, "%vGetting hot! You're close!%v\n", magentaColor, resetColor)
	}
}

func oddsOrEvenHint(game *GameInfo) {
	if game.Answer%2 == 0 {
		fmt.Fprintf(out, "%vHint:%v the number is %veven.%v\n", greenColor, resetColor, greenColor, resetColor)
	} else {
		fmt.Fprintf(out, "%vHint:%v the number is %vodd.%v\n", greenColor, redColor, greenColor, resetColor)
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

func printRange(min, max int) {
	rangeString := "✦━━━━━━━━━━━━━━━━━━━━━━✦"
	fmt.Fprintf(out, "%v %d %v %v %s %v %v %d%v\n", blueColor, min, resetColor, greenColor, rangeString, resetColor, blueColor, max, resetColor)
}

func clearScreen() {
	fmt.Fprintf(out, "\033[H\033[2J")
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
