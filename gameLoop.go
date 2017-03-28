/* Gobot V2
 * Kyle Szombathy
 * A game with a minimax implementation
 * CSC 180 "Morph" competition
 */
package main

import (
	"fmt"
	"github.com/ktodaz/gobot/gobotcore"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		gobotcore.SetDebug(false)
		GameLoop(isGobotGoingFirst())
	} else if os.Args[1] == "test" {
		fmt.Println("Test entered")
	}
}

var board gobotcore.Board

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
	board = gobotcore.NewDefaultBoard()
	fmt.Print("\nInitial Board Position:")
	board.PrintBoard()

	if gobotGoingFirst {
		executeGobotMove()
	}
	for {
		executeHumanMove()
		executeGobotMove()
	}
}

func executeGobotMove() {
	move := board.MinimaxMulti(gobotcore.GOBOT, 4)
	board.MakeMoveAndPrintMessage(move)
	board.PrintBoard()
}

func executeHumanMove() {
	move := getHumanInput()
	board.MakeMoveAndGetTakenPiece(move)
}

func getHumanInput() gobotcore.Move {
	var input string

	for !IsValidInput(input) {
		fmt.Print("Enter a move: ")
		fmt.Scan(&input)
	}
	return gobotcore.NewMove(gobotcore.NewLocationsFromString(input))
}

func IsValidInput(input string) bool {
	if len(input) != 4 {
		return false
	}
	loc1, loc2 := gobotcore.NewLocationsFromString(input)
	move := gobotcore.NewMove(loc1, loc2)

	isOnBoard := move.To().IsOnBoard() && move.From().IsOnBoard()
	return isOnBoard && board.IsValidHumanMove(move)
}

