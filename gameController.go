package main

import (
	"fmt"
	"github.com/ktodaz/gobot/gobotcore"
)

func main() {
	gobotcore.GameLoop(isGobotGoingFirst())
}

func isGobotGoingFirst() bool {
	var input int
	for input != 1 && input != 2 {
		fmt.Print("Will Gobot go first or second? Enter 1 or 2: ")
		fmt.Scan(&input)
		if input != 1 && input != 2 {
			fmt.Println("Enter a valid input")
		}
	}
	return input == 1
}
