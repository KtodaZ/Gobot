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

//type board [boardCols][boardRows]Piece

const (
	moveTime      time.Duration = time.Duration(5) // Duration of the move time
	boardCols     int           = 6
	boardRows     int           = 8
	numDiffPieces int           = 10

	// Game Pieces -  i < 0 = gobot, i > 0 = human

	//Minimax
	bestMax int = 9999999
	bestMin int = -9999999
	winMax  int = 2000000
	winMin  int = -2000000
)

var (
	board [boardCols][boardRows]Piece
	humanGoingFirst bool
	curDepth int
	curMaxDepth int
	stopSearch bool
)

func GameLoop(gobotGoingFirst bool)   {
	defer fmt.Println("\nProgram finished execution")

	setInitialPositions()
	fmt.Print("\nInitial Board Position:")
	printBoard()

	if gobotGoingFirst {
		executeGobotMove()
		printBoard()
	}
	/*for {
		executeHumanMove()
		executeGobotMove()
	}*/
}
func setInitialPositions() {
	// Gobot pieces
	board[1][7] = KING_GOB
	board[0][6] = KNIGHT_GOB
	board[1][6] = BISHOP_GOB
	board[2][6] = ROOK_GOB
	board[3][6] = ROOK_GOB
	board[4][6] = BISHOP_GOB
	board[5][6] = KNIGHT_GOB
	board[2][5] = PAWN_GOB
	board[3][5] = PAWN_GOB
	// Human pieces
	board[2][2] = PAWN_HUM
	board[3][2] = PAWN_HUM
	board[0][1] = KNIGHT_HUM
	board[1][1] = KNIGHT_HUM
	board[2][1] = ROOK_HUM
	board[3][1] = ROOK_HUM
	board[4][1] = BISHOP_HUM
	board[5][1] = KNIGHT_HUM
	board[4][0] = KING_HUM
}

func printBoard() {
	fmt.Println()
	for row := boardRows - 1; row >= 0; row-- {
		fmt.Print(row, "  ")
		for col := 0; col < boardCols; col++ {
			// Print piece
			val := &board[col][row]
			fmt.Print(val.getName())

			// Add space and/or newline
			fmt.Print(" ")
			if col == boardCols-1 {
				fmt.Println()
			}
		}
	}
	fmt.Println("   A B C D E F\n")
}

func executeGobotMove() {
	fmt.Println("Executing Gobot Move")

	computeGobotBestMove()

}

func computeGobotBestMove() int {


	//var best int = bestMin
	//var bestMove int
	//var currentBestMove int
	curDepth = 0

	//computeMovesGobot()

	if isWinningMoveInFirstMove() {

	} else {
		//timeout := time.After(moveTime * time.Second)

		curMaxDepth = 2


	}

	return -1
}





func isWinningMoveInFirstMove() bool {
	// TODO
	return false
}

func executeHumanMove() {

}
