package game

import (
	"fmt"
)

func provideHint(game *GameInfo, input int) {
	game.MinimumRange, game.MaximumRange = updateRange(game.MinimumRange, game.MaximumRange, input, game.Answer)
	if input < game.Answer {
		fmt.Printf("\n%vIncorrect!%v The number is %vgreater%v than %v.\n", redColor, resetColor, cyanColor, resetColor, input)
	} else if input > game.Answer {
		fmt.Printf("\n%vIncorrect!%v The number is %vless%v than %v.\n", redColor, resetColor, cyanColor, resetColor, input)
	}

	printRange(game.MinimumRange, game.MaximumRange)

	if utils.Abs(input-game.Answer) <= 3 {
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

func printRange(min, max int) {
	rangeString := "✦━━━━━━━━━━━━━━━━━━━━━━✦"
	fmt.Printf("%v %d %v %v %s %v %v %d%v\n", blueColor, min, resetColor, greenColor, rangeString, resetColor, blueColor, max, resetColor)
}
