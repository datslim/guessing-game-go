package utils

import (
	"fmt"
)

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
