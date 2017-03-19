/* Gobot V2
 * Kyle Szombathy
 * A game with a minimax implementation
 * CSC 180 "Morph" competition
 */
package main

import (
	"fmt"
	"github.com/ktodaz/gobot/gobotcore"
)

func main() {
	GameLoop(isGobotGoingFirst())
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

func GameLoop(gobotGoingFirst bool) {
	board := gobotcore.NewDefaultBoard()
	fmt.Print("\nInitial Board Position:")
	board.PrintBoard()

	if gobotGoingFirst {
		//gobotcore.ExecuteGobotMove()
		//gobotcore.PrintBoard()
	}
	/*for {
		executeHumanMove()
		executeGobotMove()
	}*/
}
