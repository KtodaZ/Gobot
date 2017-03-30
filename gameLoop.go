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
	"runtime"
)

var (
	board      gobotcore.Board
	depth      int = 5
	isGobotGoingFirst bool = true
)

// Default: no args
// Testing: Arg[1] = "test", Arg[2] = "nameOfFile", Arg[3] = "true"/"false"
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	board = gobotcore.NewDefaultBoard()

	if len(os.Args) == 1 {
		gobotcore.SetDebug(false)
		isGobotGoingFirst = IsGobotGoingFirst()
		GameLoop(isGobotGoingFirst)
	} else if os.Args[1] == "test" {
		if os.Args[2] == "false" {
			isGobotGoingFirst = false
		}
		gobotcore.SetDebug(false)
		testGameLoop()
	}
}
func testGameLoop() {
	if isGobotGoingFirst {
		testGobotMove()
	}
	for {
		testHumanMove()
		testGobotMove()
		if isGameOver() {
			break
		}
	}
}
func testHumanMove() {
	var input string
	//fmt.Println("Awaiting Input")
	fmt.Scanf("%s", &input)
	//fmt.Println("Input Received")
	move := gobotcore.NewMove(gobotcore.NewLocationsFromString(input))
	board.MakeMoveAndGetTakenPiece(move)

}

func testGobotMove() {
	move := board.MinimaxMulti(gobotcore.GOBOT, depth)
	board.MakeMoveAndGetTakenPiece(move)
	fmt.Println(move.ToString())
}

func isGameOver() bool {
	if board.IsGameOverForPlayer(gobotcore.GOBOT, board.LegalMovesForPlayer(gobotcore.GOBOT)) {
		fmt.Println("Lost")
		return true
	}
	if board.IsGameOverForPlayer(gobotcore.HUMAN, board.LegalMovesForPlayer(gobotcore.HUMAN)) {
		fmt.Println("Won")
		return true
	} else {
		return false
	}
}

func IsGobotGoingFirst() bool {
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
	move := board.MinimaxMulti(gobotcore.GOBOT, depth)
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
