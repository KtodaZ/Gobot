package gobotcore

import (
	"fmt"
	"strings"
	"time"
)

type Board [boardRows][boardCols]Piece

const (
	// Duration of the move time
	moveTime  time.Duration = time.Duration(5)

	// Board info
	boardCols int           = 6
	boardRows int           = 8

	//Minimax
	bestMax float64 = 9999999.0
	bestMin float64 = -9999999.0
	winMax  float64 = 2000000.0
	winMin  float64 = -2000000.0
)

var (
	curDepth    int
	curMaxDepth int
	stopSearch  bool
	debug   bool    = true
)

// ================== Getters / Utility ==================

func NewEmptyBoard() Board {
	emptyBoard := Board{}
	for row := 0; row < boardRows; row++ {
		for col := 0; col < boardCols; col++ {
			emptyBoard[row][col] = EMPTY
		}
	}
	return emptyBoard
}

func NewDefaultBoard() Board {
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
	defaultBoard[1][1] = BISHOP_HUM
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

func (board *Board) PrintBoard() {
	/*Prints board like so
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
	fmt.Println()
	for row := boardRows - 1; row >= 0; row-- {
		fmt.Print(row+1, "   ")
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
	fmt.Println("\n    A B C D E F\n")
}

func (board *Board) PieceAt(location Location) Piece {
	if location.IsOnBoard() {
		return board[location.row][location.col]
	} else {
		return -1
	}
}

func (board *Board) RetractMove(move Move, takenPiece Piece) {
	board.SetPieceAtLocation(move.from, board.PieceAt(move.to).UnMorph())
	board.SetPieceAtLocation(move.to, takenPiece)
}

func (board *Board) MakeMoveAndGetTakenPiece(move Move) Piece {
	takenPiece := board.PieceAt(move.to)
	board.SetPieceAtLocation(move.to, board.PieceAt(move.from).Morph())
	board.SetPieceAtLocation(move.from, EMPTY)
	return takenPiece
}

func (board *Board) MakeMoveAndPrintMessage(move Move) {
	piece := board.MakeMoveAndGetTakenPiece(move)
	fmt.Printf("\nMade move %s and took piece %s\n", move.ToString(), piece.GetName())
}

func (board *Board) SetPieceAtLocation(location Location, pieceToSet Piece) {
	board[location.row][location.col] = pieceToSet
}

func (board *Board) IsValidHumanMove(move Move) bool {
	return move.IsContainedIn(board.LegalMovesForPlayer(HUMAN))
}

func SetDebug(bool bool)  {
	debug = bool
}

func (board *Board) makeCopy() Board {
	newBoard := Board{}
	for row := 0; row < boardRows; row++ {
		for col := 0; col < boardCols; col++ {
			newBoard[row][col] = board[row][col]
		}
	}
	return newBoard
}

// ================== Minimax ==================
func (board *Board) Minimax(player Player, depth int) Move {
	best := ScoredMove{score: bestMin}
	bestIndex := 0 //Remove
	playerMoves := board.LegalMovesForPlayer(player)

	for i, move := range playerMoves {
		takenPiece := board.MakeMoveAndGetTakenPiece(move)

		curScore := board.Min(player.Opponent(), depth)
		if curScore > best.score {
			best.move = move
			best.score = curScore
			bestIndex = i
		}
		board.RetractMove(move, takenPiece)
	}
	if debug {
		fmt.Printf("Minimax: Found best move with score %f move %s at index %d", best.score, best.move.ToString(), bestIndex)
	}
	return best.move
}
func (board *Board) Max(player Player, depth int) float64 {
	bestScore := bestMin
	var bestMove Move
	playerMoves := board.LegalMovesForPlayer(player)

	if board.IsGameOverForPlayer(player, playerMoves) {
		return winMin
	}

	if depth == 0 {
		return board.GetWeightedScoreForPlayer(player)
	}

	for _, move := range playerMoves {
		takenPiece := board.MakeMoveAndGetTakenPiece(move)
		curScore := board.Min(player.Opponent(), depth-1)

		if curScore > bestScore {
			bestScore = curScore
			bestMove = move
		}

		board.RetractMove(move, takenPiece)
	}

	if debug {
		fmt.Printf("MAX: Found bestscore %f for depth %d moves left %d with move %s \n", bestScore, depth, len(playerMoves), bestMove.ToString())
	}

	return bestScore
}
func (board *Board) Min(player Player, depth int) float64 {
	bestScore := bestMax
	var bestMove Move
	playerMoves := board.LegalMovesForPlayer(player)

	if board.IsGameOverForPlayer(player, playerMoves) {
		return winMax
	}

	if depth == 0 {
		return board.GetWeightedScoreForPlayer(player)
	}

	for _, move := range playerMoves {
		takenPiece := board.MakeMoveAndGetTakenPiece(move)
		curScore := board.Max(player.Opponent(), depth-1)
		if curScore < bestScore {
			bestScore = curScore
			bestMove = move
		}
		board.RetractMove(move, takenPiece)
	}

	if debug {
		fmt.Printf("MIN: Found bestscore %f for depth %d moves left %d with move %s \n", bestScore, depth, len(playerMoves), bestMove.ToString())
	}
	return bestScore
}

// ================== Minimax with Goroutines ==================
func (board *Board) MinimaxMulti(player Player, depth int) Move {
	best := ScoredMove{score: bestMin}
	bestIndex := 0 //Remove
	playerMoves := board.LegalMovesForPlayer(player)
	var scoreChan = make(chan ScoredMove)

	for _, move := range playerMoves {
		takenPiece := board.MakeMoveAndGetTakenPiece(move)

		boardCopy := board.makeCopy()
		scoredMove := ScoredMove{move: move}
		go func() {
			curScore := boardCopy.MinMulti(player.Opponent(), depth)
			scoredMove.score = curScore
			scoreChan <- scoredMove
		}()

		board.RetractMove(move, takenPiece)
	}
	for i := 0; i < len(playerMoves); i++ {
		cur := <-scoreChan
		if cur.score > best.score {
			best.move = cur.move
			best.score = cur.score
			bestIndex = i
		}
	}
	if debug {
		fmt.Printf("Minimax: Found best move with score %f move %s at index %d", best.score, best.move.ToString(), bestIndex)
	}
	return best.move
}
func (board *Board) MaxMulti(player Player, depth int) float64 {
	bestScore := bestMin
	var bestMove Move
	playerMoves := board.LegalMovesForPlayer(player)
	var scoreChan = make(chan ScoredMove)

	if board.IsGameOverForPlayer(player, playerMoves) {
		return winMin
	}

	if depth == 0 {
		return board.GetWeightedScoreForPlayer(player)
	}

	for _, move := range playerMoves {
		takenPiece := board.MakeMoveAndGetTakenPiece(move)

		boardCopy := board.makeCopy()
		scoredMove := ScoredMove{move: move}
		go func() {
			curScore := boardCopy.Min(player.Opponent(), depth-1)
			scoredMove.score = curScore
			scoreChan <- scoredMove
		}()

		board.RetractMove(move, takenPiece)
	}
	for i := 0; i < len(playerMoves); i++ {
		cur := <-scoreChan
		if cur.score > bestScore {
			bestScore = cur.score
			bestMove = cur.move
		}
	}
	if debug {
		fmt.Printf("MAX: Found bestscore %f for depth %d moves left %d with move %s \n", bestScore, depth, len(playerMoves), bestMove.ToString())
	}

	return bestScore
}
func (board *Board) MinMulti(player Player, depth int) float64 {
	bestScore := bestMax
	var bestMove Move
	playerMoves := board.LegalMovesForPlayer(player)
	var scoreChan = make(chan ScoredMove)

	if board.IsGameOverForPlayer(player, playerMoves) {
		return winMax
	}

	if depth == 0 {
		return board.GetWeightedScoreForPlayer(player)
	}

	for _, move := range playerMoves {
		takenPiece := board.MakeMoveAndGetTakenPiece(move)

		boardCopy := board.makeCopy()
		scoredMove := ScoredMove{move: move}
		go func() {
			curScore := boardCopy.Max(player.Opponent(), depth-1)
			scoredMove.score = curScore
			scoreChan <- scoredMove
		}()

		board.RetractMove(move, takenPiece)
	}
	for i := 0; i < len(playerMoves); i++ {
		cur := <-scoreChan
		if cur.score < bestScore {
			bestScore = cur.score
			bestMove = cur.move
		}
	}
	
	if debug {
		fmt.Printf("MIN: Found bestscore %f for depth %d moves left %d with move %s \n", bestScore, depth, len(playerMoves), bestMove.ToString())
	}
	return bestScore
}

func (board *Board) IsGameOverForPlayer(player Player, playerMoves []Move) bool {
	return board.isKingDeadForPlayer(player) || len(playerMoves) == 0
}

func (board *Board) isKingDeadForPlayer(player Player) bool {
	for row := 0; row < boardRows; row++ {
		for col := 0; col < boardCols; col++ {
			piece := board.PieceAt(Location{row: row, col: col})
			if piece.IsKing() && piece.IsOwnedBy(player) {
				return false
			}
		}
	}
	return true
}

// ================== Heuristic ==================
func (board *Board) GetWeightedScoreForPlayer(player Player) float64 {
	score := 0.0
	for row := 0; row < boardRows; row++ {
		for col := 0; col < boardCols; col++ {
			location := Location{row: row, col: col}
			piece := board.PieceAt(location)

			if piece.IsOwnedBy(player) {
				score += piece.Weight()
			} else {
				score -= piece.Weight()
			}
		}
	}
	return score
}

// ================== Legal Moves ==================

func (board *Board) LegalMovesForPlayer(player Player) []Move {
	totalMoves := []Move{}

	for row := 0; row < boardRows; row++ {
		for col := 0; col < boardCols; col++ {

			location := NewLocation(col, row)
			piece := board.PieceAt(location)

			if !piece.IsEmpty() && piece.IsOwnedBy(player) {
				currentMoves := board.FindMovesForPlayersPieceAtLocation(player, location)
				totalMoves = append(totalMoves, currentMoves...)
			}
		}
	}

	return totalMoves
}

func (board *Board) LegalMovesForPlayerMulti(player Player) []Move {
	totalMoves := []Move{}
	var countGoRoutines int = 0
	var movesChan = make(chan []Move)

	for row := 0; row < boardRows; row++ {
		for col := 0; col < boardCols; col++ {

			location := NewLocation(col, row)
			piece := board.PieceAt(location)

			if !piece.IsEmpty() && piece.IsOwnedBy(player) {

				// Create goRoutines so we can quickly find all the moves
				countGoRoutines++
				go func() {
					currentMoves := board.FindMovesForPlayersPieceAtLocation(player, location)
					movesChan <- currentMoves
				}()
			}
		}
	}
	for i := 0; i < countGoRoutines; i++ {
		totalMoves = append(totalMoves, <-movesChan...)
	}

	return totalMoves
}

func (board *Board) FindMovesForPlayersPieceAtLocation(player Player, location Location) []Move {
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

func (board *Board) FindMovesForBishopAtLocation(player Player, originalLocation Location) []Move {
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

func (board *Board) FindMovesForRookAtLocation(player Player, originalLocation Location) []Move {
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

func (board *Board) FindMovesForKnightAtLocation(player Player, originalLocation Location) []Move {
	moves := []Move{}

	applyKnightMove := func(colsToAppendBy, rowsToAppendBy int) {
		isMovingBackward := (rowsToAppendBy < 0 && player == HUMAN) || (rowsToAppendBy > 0 && player == GOBOT)

		move := Move{from: originalLocation, to: originalLocation.Append(colsToAppendBy, rowsToAppendBy)}
		piece := board.PieceAt(move.to)

		isValidMove := board.isValidMove(move, player)
		isValidForwardMove := isValidMove && !isMovingBackward
		isValidBackwardsMove := isValidMove && isMovingBackward && piece.IsOwnedBy(player.Opponent())

		if isValidForwardMove || isValidBackwardsMove {
			moves = append(moves, move)
		}
	}

	//N direction
	applyKnightMove(1, 2)
	applyKnightMove(-1, 2)
	// E direction
	applyKnightMove(-2, 1)
	applyKnightMove(-2, -1)
	// S direction
	applyKnightMove(1, -2)
	applyKnightMove(-1, -2)
	// W direction
	applyKnightMove(2, 1)
	applyKnightMove(2, -1)

	return moves
}

func (board *Board) FindMovesForPawnAtLocation(player Player, originalLocation Location) []Move {
	moves := []Move{}
	var move Move

	// Get direction
	var direction int = 1
	if player == GOBOT {
		direction = -1
	}

	// Forward move
	move = Move{originalLocation, originalLocation.Append(0, 1*direction)}
	if board.PieceAt(move.to).IsEmpty() {
		moves = append(moves, move)
	}

	// Attack moves
	applyAttackMove := func(colsToAppendBy, rowsToAppendBy int) {
		move = Move{originalLocation, originalLocation.Append(colsToAppendBy, rowsToAppendBy)}
		if board.PieceAt(move.to).IsOwnedBy(player.Opponent()) {
			moves = append(moves, move)
		}
	}
	applyAttackMove(1, 1*direction)
	applyAttackMove(-1, 1*direction)

	return moves
}

func (board *Board) FindMovesForKingAtLocation(player Player, originalLocation Location) []Move {
	moves := []Move{}

	// Get direction
	var direction int = 1
	if player == GOBOT {
		direction = -1
	}

	move := Move{originalLocation, originalLocation.Append(-1*direction, 0)}
	if board.PieceAt(move.to).IsEmpty() {
		moves = append(moves, move)
	}

	move = Move{originalLocation, originalLocation.Append(1, 0)}
	if board.PieceAt(move.to).IsOwnedBy(player.Opponent()) {
		moves = append(moves, move)
	}
	move = Move{originalLocation, originalLocation.Append(-1, 0)}
	if board.PieceAt(move.to).IsOwnedBy(player.Opponent()) {
		moves = append(moves, move)
	}

	return moves
}

func (board *Board) getDirectionalMoves(player Player, originalLocation Location, colsToAppendBy int, rowsToAppendBy int) []Move {
	moves := []Move{}
	moveLocation := originalLocation
	movingBackward := (rowsToAppendBy < 0 && player == HUMAN) || (rowsToAppendBy > 0 && player == GOBOT)
	moveToNextLocation := func() { moveLocation = moveLocation.Append(colsToAppendBy, rowsToAppendBy) }

	moveToNextLocation()
	for moveLocation.IsOnBoard() {
		pieceToMoveTo := board.PieceAt(moveLocation)

		if movingBackward && pieceToMoveTo.IsEmpty() {
			moveToNextLocation()
			continue
		}
		if pieceToMoveTo.IsOwnedBy(player) {
			break // Don't do the move
		}

		moves = append(moves, Move{from: originalLocation, to: moveLocation})

		if pieceToMoveTo.IsOwnedBy(player.Opponent()) {
			break
		}
		moveToNextLocation()
	}
	return moves
}

func (board *Board) isValidMove(move Move, player Player) bool {
	piece := board.PieceAt(move.to)
	return !piece.IsOwnedBy(player) && move.to.IsOnBoard()
}
