/* Gobot V2
 * Kyle Szombathy
 * A game with a minimax implementation
 * CSC 180 "Morph" competition
 */
package gobotcore

import (
	"fmt"
	"time"
)

const (
	moveTime      time.Duration = time.Duration(5) // Duration of the move time
	boardCols     int           = 6
	boardRows     int           = 8
	numDiffPieces int           = 10

	//Minimax
	bestMax int = 9999999
	bestMin int = -9999999
	winMax  int = 2000000
	winMin  int = -2000000
)

var (
	board           [boardCols][boardRows]Piece
	humanGoingFirst bool
	curDepth        int
	curMaxDepth     int
	stopSearch      bool
)



func SetInitialPositions() {

}



func ExecuteGobotMove() {
	fmt.Println("Executing Gobot Move")

	//computeGobotBestMove()

}

func ComputeGobotBestMove() int {

	//var best int = bestMin
	//var bestMove int
	//var currentBestMove int
	curDepth = 0

	//computeMovesGobot()

	/*if isWinningMoveInFirstMove() {

	} else {
		//timeout := time.After(moveTime * time.Second)

		curMaxDepth = 2

	}*/

	return -1
}

func IsWinningMoveInFirstMove() bool {
	// TODO
	return false
}

func executeHumanMove() {

}
