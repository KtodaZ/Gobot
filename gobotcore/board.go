package gobotcore

import (
	"sync"
	"fmt"
	"strings"
)

type Board [boardRows][boardCols]Piece

func NewEmptyBoard() Board {
	emptyBoard := Board{}
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			emptyBoard[row][col] = EMPTY
		}
	}
	return emptyBoard
}

/*Creates a board like so:
7   - K - - - -
6   N B R R B N
5   - - P P - -
4   - - - - - -
3   - - - - - -
2   - - p p - -
1   n n r r b n
0   - - - - k -

   A B C D E F
 */
func NewDefaultBoard() Board {
	defaultBoard := Board{}
	// Gobot pieces
	defaultBoard[7][1] = KING_GOB
	defaultBoard[6][0] = KNIGHT_GOB
	defaultBoard[6][1] = BISHOP_GOB
	defaultBoard[6][2] = ROOK_GOB
	defaultBoard[6][3] = ROOK_GOB
	defaultBoard[6][4] = BISHOP_GOB
	defaultBoard[6][5] = KNIGHT_GOB
	defaultBoard[5][2] = PAWN_GOB
	defaultBoard[5][3] = PAWN_GOB
	// Human pieces
	defaultBoard[2][2] = PAWN_HUM
	defaultBoard[2][3] = PAWN_HUM
	defaultBoard[1][0] = KNIGHT_HUM
	defaultBoard[1][1] = KNIGHT_HUM
	defaultBoard[1][2] = ROOK_HUM
	defaultBoard[1][3] = ROOK_HUM
	defaultBoard[1][4] = BISHOP_HUM
	defaultBoard[1][5] = KNIGHT_HUM
	defaultBoard[0][4] = KING_HUM
	return defaultBoard
}

func NewBoardFromString(boardString string) Board {
	/*Board format is like so:
	7   - K - - - -
	6   N B R R B N
	5   - - P P - -
	4   - - - - - -
	3   - - - - - -
	2   - - p p - -
	1   n n r r b n
	0   - - - - k -

	   A B C D E F
	*/
	board := Board{}
	boardStringRows := strings.Split(boardString, "\n")
	if len(boardStringRows) != 10 {
		panic("Incorrect input row size")
	}
	if len(boardStringRows[0]) != 15 {
		panic("Incorrect input col size")
	}

	// These are not the magic numbers you're looking for
	for i := len(boardStringRows) - 3; i >= 0; i-- {
		boardStringRow := boardStringRows[i]
		for j := 4; j < len(boardStringRow); j = j + 2 {
			colIndex := (j - 4) / 2
			rowIndex := 7 - i
			board[rowIndex][colIndex] = GetPieceByName(string(boardStringRow[j]))
		}
	}
	return board
}

func (board Board) PrintBoard() {
	fmt.Println()
	for row := boardRows - 1; row >= 0; row-- {
		fmt.Print(row, "   ")
		for col := 0; col < boardCols; col++ {
			// Print piece
			val := &board[row][col]
			fmt.Print(val.GetName())

			// Add space and/or newline
			fmt.Print(" ")
			if col == boardCols-1 {
				fmt.Println()
			}
		}
	}
	fmt.Println("\n   A B C D E F\n")
}

func (board Board) GetMovesForPlayer(player Player) []Move {
	totalMoves := []Move{}
	var waitGroup sync.WaitGroup

	for row := 0; row < boardRows; row++ {
		for col := 0; col < boardCols; col++ {

			location := NewLocation(col, row)
			piece := board.PieceAt(location)

			if !piece.IsEmpty() && piece.OwnedBy(player) {

				// Create goRoutines so we can quickly find all the moves
				waitGroup.Add(1)
				go func(newBoard Board, newLocation Location) {
					defer waitGroup.Done()
					currentMoves := newBoard.FindMovesForPlayersPieceAtLocation(player, newLocation)
					totalMoves = append(totalMoves, currentMoves...)
				}(board, location)
			}
		}
	}

	// Wait until they are all done, then return
	waitGroup.Wait()
	return totalMoves
}

func (board Board) FindMovesForPlayersPieceAtLocation(player Player, location Location) []Move {
	piece := board.PieceAt(location)

	switch piece {
	case BISHOP_GOB:
		return board.FindMovesForBishopAtLocation(player, location)
	case BISHOP_HUM:
		return board.FindMovesForBishopAtLocation(player, location)
	case ROOK_GOB:
		return board.FindMovesForRookAtLocation(player, location)
	case ROOK_HUM:
		return board.FindMovesForRookAtLocation(player, location)
	case KNIGHT_GOB:
		return board.FindMovesForKnightAtLocation(player, location)
	case KNIGHT_HUM:
		return board.FindMovesForKnightAtLocation(player, location)
	case PAWN_GOB:
		return board.FindMovesForPawnAtLocation(player, location)
	case PAWN_HUM:
		return board.FindMovesForPawnAtLocation(player, location)
	case KING_GOB:
		return board.FindMovesForKingAtLocation(player, location)
	case KING_HUM:
		return board.FindMovesForKingAtLocation(player, location)
	}
	return []Move{}
}

func (board Board) FindMovesForBishopAtLocation(player Player, originalLocation Location) []Move {
	moves := []Move{}

	// NE direction
	moves = append(moves, board.getDirectionalMoves(player, originalLocation, 1, 1)...)
	// NW direction
	moves = append(moves, board.getDirectionalMoves(player, originalLocation, -1, 1)...)
	// SW direction
	moves = append(moves, board.getDirectionalMoves(player, originalLocation, -1, -1)...)
	// SE direction
	moves = append(moves, board.getDirectionalMoves(player, originalLocation, 1, -1)...)

	return moves
}

func (board Board) FindMovesForRookAtLocation(player Player, originalLocation Location) []Move {
	moves := []Move{}

	// N direction
	moves = append(moves, board.getDirectionalMoves(player, originalLocation, 0, 1)...)
	// W direction
	moves = append(moves, board.getDirectionalMoves(player, originalLocation, -1, 0)...)
	// S direction
	moves = append(moves, board.getDirectionalMoves(player, originalLocation, 0, -1)...)
	// E direction
	moves = append(moves, board.getDirectionalMoves(player, originalLocation, 1, 0)...)

	return moves
}

func (board Board) getDirectionalMoves(player Player, originalLocation Location, colsToAppendBy int, rowsToAppendBy int) []Move {
	moves := []Move{}
	moveLocation := originalLocation.Append(colsToAppendBy, rowsToAppendBy)
	for moveLocation.IsOnBoard() {
		piece := board.PieceAt(moveLocation)
		if piece.OwnedBy(player) {
			break
		}

		move := Move{
			from: originalLocation,
			to:   moveLocation,
		}

		moves = append(moves, move)

		if piece.OwnedBy(player.Opponent()) {
			break
		}
		moveLocation = moveLocation.Append(colsToAppendBy, rowsToAppendBy)
	}
	return moves
}

func (board Board) FindMovesForKnightAtLocation(player Player, originalLocation Location) []Move {
	return []Move{}
}

func (board Board) FindMovesForPawnAtLocation(player Player, originalLocation Location) []Move {
	return []Move{}
}

func (board Board) FindMovesForKingAtLocation(player Player, originalLocation Location) []Move {
	return []Move{}
}

func (board Board) PieceAt(location Location) Piece {
	if location.IsOnBoard() {
		return board[location.row][location.col]
	} else {
		return -1
	}
}
