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
	board             gobotcore.Board
	depth             int8 = 6
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
		gobotMoveSimple()
	}
	for {
		humanMoveSimple()
		gobotMoveSimple()
		if isGameOver() {
			break
		}
	}
}
func humanMoveSimple() {
	var input string
	//fmt.Println("Awaiting Input")
	fmt.Scanf("%s", &input)
	//fmt.Println("Input Received")
	move := gobotcore.NewMove(gobotcore.NewLocationsFromString(input))
	board.MakeMoveAndGetTakenPiece(&move)

}

func gobotMoveSimple() {
	gobot := gobotcore.Player(gobotcore.GOBOT)
	move := board.MinimaxMulti(&gobot, &depth)
	board.MakeMoveAndGetTakenPiece(move.Move())
	fmt.Println(move.Move().ToStringFlipped())
}

func isGameOver() bool {
	gobot := gobotcore.Player(gobotcore.GOBOT)
	gobotMoves := board.LegalMovesForPlayer(gobot)
	if board.IsGameOverForPlayer(&gobot, &gobotMoves) {
		fmt.Println("Lost")
		return true
	}
	human := gobotcore.Player(gobotcore.HUMAN)
	humanMoves := board.LegalMovesForPlayer(human)
	if board.IsGameOverForPlayer(&human, &humanMoves) {
		fmt.Println("Won")
		return true
	} else {
		return false
	}
}

func IsGobotGoingFirst() bool {
	var input int8
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
		gobotMoveFriendly()
	}
	for {
		humanMoveFriendly()
		gobotMoveFriendly()
		if isGameOverFriendly() {
			break
		}
	}
}

func humanMoveFriendly() {
	move := getHumanInput()
	board.MakeMoveAndGetTakenPiece(move)
}

func getHumanInput() *gobotcore.Move {
	var input string

	for !IsValidInput(input) {
		fmt.Print("Enter a move: ")
		fmt.Scan(&input)
	}
	loc1, loc2 := gobotcore.NewLocationsFromString(input)
	move := gobotcore.NewMove(loc1, loc2)
	return &move
}

func IsValidInput(input string) bool {
	if len(input) != 4 {
		return false
	}
	loc1, loc2 := gobotcore.NewLocationsFromString(input)
	move := gobotcore.NewMove(loc1, loc2)

	isOnBoard := move.To().IsOnBoard() && move.From().IsOnBoard()
	return isOnBoard && board.IsValidHumanMove(&move)
}

func gobotMoveFriendly() {
	gobot := gobotcore.Player(gobotcore.GOBOT)
	move := board.MinimaxMulti(&gobot, &depth)
	board.MakeMoveAndPrintMessage(move.Move())
	board.PrintBoard()
}

func isGameOverFriendly() bool {
	gobot := gobotcore.Player(gobotcore.GOBOT)
	gobotMoves := board.LegalMovesForPlayer(gobot)
	if board.IsGameOverForPlayer(&gobot, &gobotMoves) {
		fmt.Println("Human Won")
		return true
	}
	human := gobotcore.Player(gobotcore.HUMAN)
	humanMoves := board.LegalMovesForPlayer(human)
	if board.IsGameOverForPlayer(&human, &humanMoves) {
		fmt.Println("Gobot Won")
		return true
	} else {
		return false
	}
}
