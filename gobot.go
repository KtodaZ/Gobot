/* Gobot V2
 * Kyle Szombathy
 * A game with a minimax implementation
 * CSC 180 "Morph" competition
 */
package main

import (
	"fmt"
	"time"
	"math/rand"
)

const (
	moveTime          = time.Duration(5) // Duration of the move time
	boardCols     int = 6
	boardRows     int = 8
	numDiffPieces int = 10

	// Game Pieces -  i < 0 = gobot, i > 0 = human
	bishopGob int = -1
	bishopHum int = 1
	rookGob   int = -2
	rookHum   int = 2
	knightGob int = -3
	knightHum int = 3
	pawnGob   int = -4
	pawnHum   int = 4
	kingGob   int = -5
	kingHum   int = 5

	bishopGobName rune = 'B'
	bishopHumName rune = 'b'
	rookGobName   rune = 'R'
	rookHumName   rune = 'r'
	knightGobName rune = 'N'
	knightHumName rune = 'n'
	pawnGobName   rune = 'P'
	pawnHumName   rune = 'p'
	kingGobName   rune = 'K'
	kingHumName   rune = 'k'

	//Minimax
	bestMax = 9000000
	bestMin = -9000000
	winMax  = 2000000
	winMin  = -2000000
)

var (
	board           [boardCols][boardRows]int
	humanGoingFirst bool
	pieceNames      = [...]rune{kingGobName, pawnGobName, knightGobName, rookGobName, bishopGobName, '-',
								bishopHumName, rookHumName, knightHumName, pawnHumName, kingHumName}
	taunts = [...]string{"Gobot will prevail", "Fear the Gobot", "Do not underestimate the power of Gobot",
						 "Gobot will destroy you", "Feel the power of Gobot", "Come with me if you want to live"}
)

func main() {
	//timeout := time.After(moveTime * time.Second)
	defer fmt.Println("\nProgram finished execution")

	setInitialPositions()
	printBoard()

	if isGobotGoingFirst() {
		executeGobotMove()
	}
	/*for {
		executeHumanMove()
		executeGobotMove()
	}*/
}
func setInitialPositions() {
	// Gobot pieces
	board[1][7] = kingGob
	board[0][6] = knightGob
	board[1][6] = bishopGob
	board[2][6] = rookGob
	board[3][6] = rookGob
	board[4][6] = bishopGob
	board[5][6] = knightGob
	board[2][5] = pawnGob
	board[3][5] = pawnGob
	// Human pieces
	board[2][2] = pawnHum
	board[3][2] = pawnHum
	board[0][1] = knightHum
	board[1][1] = knightHum
	board[2][1] = rookHum
	board[3][1] = rookHum
	board[4][1] = bishopHum
	board[5][1] = knightHum
	board[4][0] = kingHum
}

func printBoard() {
	fmt.Println()
	for row := boardRows - 1; row >= 0; row-- {
		fmt.Print(row, "  ")
		for col := 0; col < boardCols; col++ {
			// Print piece
			val := &board[col][row]
			fmt.Print(string(pieceNames[*val+5]))

			// Add space and/or newline
			fmt.Print(" ")
			if col == boardCols-1 {
				fmt.Println()
			}
		}
	}
	fmt.Println("   A B C D E F\n")
}

func isGobotGoingFirst() bool {
	var input int
	for input != 1 && input != 2 {
		tauntOpponent()
		fmt.Print("Will Gobot go first or second? Enter 1 or 2: ")
		fmt.Scan(&input)
		if input != 1 && input != 2 {
			fmt.Println("Enter a valid input")
		}
	}
	return input == 1
}

func tauntOpponent() {
	fmt.Println("WARNING: " + taunts[int(rand.Float64()*float64(len(taunts)))] + "\n")
}

func executeGobotMove() {
	computeGobotBestMove()
	printBoard()
}

func computeGobotBestMove() {

}

func executeHumanMove() {

}
