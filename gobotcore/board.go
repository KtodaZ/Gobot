package gobotcore

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type Board [boardRows][boardCols]Piece

const (
	// Duration of the move time
	moveTime time.Duration = 5 * time.Second

	// Board info
	boardCols int8 = 6
	boardRows int8 = 8

	//Minimax
	bestMax float32 = 9999999.0
	bestMin float32 = -9999999.0
	winMax  float32 = 2000000.0
	winMin  float32 = -2000000.0
)

var (
	curDepth      int8
	curMaxDepth   int8
	stopSearch    bool
	debug         bool = true
	timeOver      bool
	numGoRoutines int
)

// ================== Getters / Utility ==================

func NewEmptyBoard() Board {
	emptyBoard := Board{}
	var row, col int8
	for row = 0; row < boardRows; row++ {
		for col = 0; col < boardCols; col++ {
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
	var row, col int8
	for row = boardRows - 1; row >= 0; row-- {
		fmt.Print(row+1, "   ")
		for col = 0; col < boardCols; col++ {
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

func (board *Board) PieceAt(location *Location) Piece {
	if location.IsOnBoard() {
		return board[location.row][location.col]
	} else {
		return -1
	}
}

func (board *Board) RetractMove(move *Move, takenPiece Piece) {
	board.SetPieceAtLocation(&move.from, board.PieceAt(&move.to).UnMorph())
	board.SetPieceAtLocation(&move.to, takenPiece)
}

func (board *Board) MakeMoveAndGetTakenPiece(move *Move) *Piece {
	takenPiece := board.PieceAt(&move.to)
	board.SetPieceAtLocation(&move.to, board.PieceAt(&move.from).Morph())
	board.SetPieceAtLocation(&move.from, EMPTY)
	return &takenPiece
}

func (board *Board) MakeMoveAndPrintMessage(move *Move) {
	piece := board.MakeMoveAndGetTakenPiece(move)
	fmt.Printf("\nGobot made move %s (%s)", move.ToString(), move.ToStringFlipped())
	if *piece != EMPTY {
		fmt.Printf(" and captured Human piece %s", piece.GetName())
	}
	fmt.Println()
}

func (board *Board) SetPieceAtLocation(location *Location, pieceToSet Piece) {
	board[location.row][location.col] = pieceToSet
}

func (board *Board) IsValidHumanMove(move *Move) bool {
	player := Player(HUMAN)
	moves := board.LegalMovesForPlayer(player)
	return move.IsContainedIn(&moves)
}

func SetDebug(bool bool) {
	debug = bool
}

// ================== Minimax with Goroutines ==================

// Goroutines are essentially lightweight pseudo-threads.
// I am creating many goRoutines in the below "multi" functions.
// There is some performance impact by creating many goRoutines because we are creating copies of the board object
// Therefore, we end the goroutine recursion at the second level, and switch to an iterative approach
func (board *Board) MinimaxMulti(player *Player, depth *int8) ScoredMove {
	best := ScoredMove{score: bestMin}
	playerMoves := board.LegalMovesForPlayer(*player)
	opponent := player.Opponent()
	numGoRoutines = 0

	// This go channel is the communication link between the goRoutines and this function
	// Go primarily uses message passing between goRoutines and their parents
	scoreChan := make(chan ScoredMove)

	// Channel that will carry on to the goRoutines to tell them to stop early if needed
	timeOver = false
	go func() {
		time.Sleep(moveTime)
		timeOver = true
	}()

	fmt.Printf("Going to depth %d\n", int(*depth))

	for _, move := range playerMoves {
		boardCopy := *board
		boardCopy.MakeMoveAndGetTakenPiece(&move)

		scoredMove := ScoredMove{move: move}

		go func() { // Initiate a goRoutine.
			// &bestScore passes a pointer to the ever-changing bestScore variable.
			// This will ensure that no matter what stage the goRoutine is in it has the ability to
			// see what its parents best score is.
			curScore := boardCopy.MinMulti(opponent, depth, &best.score)
			scoredMove.score = curScore
			scoreChan <- scoredMove // Pass the scoredMove object back to the scoreChan channel
		}()
	}

	for i := 0; i < len(playerMoves); i++ { // Loop until all goRoutines are done
		cur := <-scoreChan // Execution will halt here and will wait until next goRoutine is done
		if cur.score > best.score {
			/*if debug {
				fmt.Print("NumGoRoutines: ")
				fmt.Println(runtime.NumGoroutine())
			}*/

			best.move = cur.move
			best.score = cur.score
		}
	}

	if !timeOver {
		newDepth := *depth + 1
		newBest := board.MinimaxMulti(player, &newDepth)
		if newBest.score > best.score {
			best = newBest
		}
	}

	return best
}

// I Found that ending the goroutine recursion at the second level is the most optimal. That is why there is no MaxMulti function.
func (board *Board) MinMulti(player *Player, depth *int8, parentsBestScore *float32) float32 {
	bestScore := bestMax
	//	var bestMove Move
	playerMoves := board.LegalMovesForPlayer(*player)
	scoreChan := make(chan ScoredMove)
	stopChan := make(chan struct{})
	newDepth := *depth - 1
	//var bestMove Move

	if board.IsGameOverForPlayer(player, &playerMoves) {
		return winMax
	}

	if newDepth == 0 {
		return board.GetWeightedScoreForPlayer(player)
	}

	for _, move := range playerMoves {
		boardCopy := *board
		boardCopy.MakeMoveAndGetTakenPiece(&move)

		scoredMove := ScoredMove{move: move}
		go func() {
			curScore := boardCopy.MaxMulti(player.Opponent(), &newDepth, &bestScore, stopChan, len(playerMoves))
			scoredMove.score = curScore
			scoreChan <- scoredMove
		}()
	}

	for i := 0; i < len(playerMoves); i++ {

		cur := <-scoreChan
		/*if debug {
			fmt.Print("NumGoRoutines: ")
			fmt.Println(runtime.NumGoroutine())
		}*/

		if cur.score < bestScore {
			bestScore = cur.score
			//bestMove = cur.move
		}

		// alpha-beta pruning
		if cur.score < *parentsBestScore {
			/*if debug {
				fmt.Printf("Stopping goRoutines because curScore %f is less than parents best score %f\n", bestScore, *parentsBestScore)
				fmt.Print("NumGoRoutines: ")
				fmt.Println(runtime.NumGoroutine())
			}*/
			close(stopChan) // Send message to all the goRoutines to tell them to stop. We don't care about their output now
			return bestScore
		}
	}

	/*if debug {
		fmt.Printf("MIN%d: Found bestscore %f moves left %d with move %s \n", depth, bestScore, len(playerMoves), bestMove.ToString())
	}*/

	return bestScore
}

func (board *Board) MaxMulti(player *Player, depth *int8, parentsBestScore *float32, parentStopChan <-chan struct{}, numParentMoves int) float32 {
	bestScore := bestMin
	//	var bestMove Move
	playerMoves := board.LegalMovesForPlayer(*player)
	scoreChan := make(chan ScoredMove)
	stopChan := make(chan struct{})
	newDepth := *depth - 1
	//var bestMove Move

	if board.IsGameOverForPlayer(player, &playerMoves) {
		return winMin
	}

	if newDepth == 0 {
		return float32(len(playerMoves)*2) - float32(numParentMoves*2) + board.GetWeightedScoreForPlayer(player)
	}

	for _, move := range playerMoves {
		boardCopy := *board
		boardCopy.MakeMoveAndGetTakenPiece(&move)

		scoredMove := ScoredMove{move: move}
		go func() {
			// Call min because we are done doing recursion with goRoutines
			curScore := boardCopy.Min(player.Opponent(), &newDepth, &bestScore, stopChan, len(playerMoves))
			scoredMove.score = curScore
			scoreChan <- scoredMove
		}()
	}

	for i := 0; i < len(playerMoves); i++ {
		select {
		default:
			cur := <-scoreChan
			/*if debug {
				fmt.Print("NumGoRoutines: ")
				fmt.Println(runtime.NumGoroutine())
			}*/

			if cur.score > bestScore {
				bestScore = cur.score
				//bestMove = cur.move
			}

			// alpha-beta pruning
			if cur.score > *parentsBestScore {
				/*if debug {
					fmt.Printf("Stopping goRoutines because curScore %f is less than parents best score %f\n", bestScore, *parentsBestScore)
					fmt.Print("NumGoRoutines: ")
					fmt.Println(runtime.NumGoroutine())
				}*/
				close(stopChan) // Send message to all the goRoutines to tell them to stop. We don't care about their output now
				return bestScore
			}
		case <-parentStopChan:
			// Parent told us to stop execution.. must have been a bad child
			/*if debug {
				fmt.Printf("Max%d: stopped execution\n", newDepth)
			}*/
			close(stopChan)
			return bestScore // Returning this score shouldn't do anything
		}

	}

	/*if debug {
		fmt.Printf("MAX%d: Found bestscore %f moves left %d with move %s \n", depth, bestScore, len(playerMoves), bestMove.ToString())
	}*/

	return bestScore
}

func (board *Board) Max(player *Player, depth *int8, parentsBestScore *float32, stopChan <-chan struct{}, numParentMoves int) float32 {
	playerMoves := board.LegalMovesForPlayer(*player)

	if board.IsGameOverForPlayer(player, &playerMoves) {
		return winMin
	}

	newDepth := *depth - 1

	if newDepth == 0 {
		return float32(len(playerMoves)*2) - float32(numParentMoves*2) + board.GetWeightedScoreForPlayer(player)
	}

	//var bestMove Move
	bestScore := bestMin
	sort.Sort(playerMoves)

	for _, move := range playerMoves {
		takenPiece := *board.MakeMoveAndGetTakenPiece(&move)
		curScore := board.Min(player.Opponent(), &newDepth, &bestScore, stopChan, len(playerMoves))

		select {
		default:
			if curScore > bestScore {
				bestScore = curScore
				//bestMove = move

				if timeOver {
					/*if debug {
						fmt.Printf("Max%d: stopped execution. Out of time\n", newDepth)
					}*/
					return bestScore
				}

				// alpha-beta pruning
				if bestScore > *parentsBestScore {
					board.RetractMove(&move, takenPiece)
					/*if debug {
						fmt.Printf("MAX%d: AB Pruning because curScore %f is more than parents best score %f\n", newDepth, bestScore, *parentsBestScore)
					}*/
					return bestScore
				}
			}

			board.RetractMove(&move, takenPiece)

		case <-stopChan:
			// Parent told us to stop execution.. must have been a bad child
			/*if debug {
				fmt.Printf("Max%d: stopped execution\n", newDepth)
			}*/
			return bestScore // Returning this score shouldn't do anything
		}
	}

	/*if debug {
		fmt.Printf("MAX%d: Found bestscore %f moves left %d with move %s \n", newDepth, bestScore, len(playerMoves), bestMove.ToString())
	}*/

	return bestScore
}
func (board *Board) Min(player *Player, depth *int8, parentsBestScore *float32, parentStopChan <-chan struct{}, numParentMoves int) float32 {
	playerMoves := board.LegalMovesForPlayer(*player)

	if board.IsGameOverForPlayer(player, &playerMoves) {
		return winMax
	}

	newDepth := *depth - 1

	if newDepth == 0 {
		return float32(len(playerMoves)*2) - float32(numParentMoves*2) + board.GetWeightedScoreForPlayer(player)
	}

	//var bestMove Move
	bestScore := bestMax
	sort.Sort(playerMoves)

	for _, move := range playerMoves {
		takenPiece := *board.MakeMoveAndGetTakenPiece(&move)
		curScore := board.Max(player.Opponent(), &newDepth, &bestScore, parentStopChan, len(playerMoves))

		select {
		default:
			if curScore < bestScore {
				bestScore = curScore
				//bestMove = move

				if timeOver {
					/*if debug {
						fmt.Printf("Max%d: stopped execution. Out of time\n", newDepth)
					}*/
					return bestScore
				}

				// alpha-beta pruning
				if bestScore < *parentsBestScore {
					board.RetractMove(&move, takenPiece)
					/*if debug {
						fmt.Printf("MAX%d: AB Pruning because curScore %f is less than parents best score %f\n", newDepth, bestScore, *parentsBestScore)
					}*/
					return bestScore
				}

			}

			board.RetractMove(&move, takenPiece)

		case <-parentStopChan:
			// Parent told us to stop execution.. must have been a bad child
			/*if debug {
				fmt.Printf("Min%d: stopped execution\n", newDepth)
			}*/
			return bestScore // Returning this score shouldn't do anything
		}

	}

	/*if debug {
		fmt.Printf("MIN%d: Found bestscore %f moves left %d with move %s \n", newDepth, bestScore, len(playerMoves), bestMove.ToString())
	}*/
	return bestScore
}

func (board *Board) IsGameOverForPlayer(player *Player, playerMoves *Moves) bool {
	return board.isKingDeadForPlayer(player) || len(*playerMoves) == 0
}

func (board *Board) isKingDeadForPlayer(player *Player) bool {
	var row, col int8
	for row = 0; row < boardRows; row++ {
		for col = 0; col < boardCols; col++ {
			location := Location{row: row, col: col}
			piece := board.PieceAt(&location)
			if piece.IsKing() && piece.IsOwnedBy(player) {
				return false
			}
		}
	}
	return true
}

// ================== Heuristic ==================
func (board *Board) GetWeightedScoreForPlayer(player *Player) float32 {
	var score float32
	var row, col int8
	score = 0.0
	for row = 0; row < boardRows; row++ {
		for col = 0; col < boardCols; col++ {
			location := Location{row: row, col: col}
			piece := board.PieceAt(&location)

			if piece.IsOwnedBy(player) {
				score += piece.Weight()
			} else {
				score -= piece.Weight()
			}
			//score += getPositionScore(player, &location)
		}
	}
	return score
}

func getPositionScore(player *Player, location *Location) float32 {
	if *player == HUMAN {
		if location.row < 3 {
			return -0.4
		}
	} else {
		if location.row > 4 {
			return -0.4
		}
	}
	return 0
}

// ================== Legal Moves ==================

func (board *Board) LegalMovesForPlayer(player Player) Moves {
	var row, col int8
	totalMoves := Moves{}

	for row = 0; row < boardRows; row++ {
		for col = 0; col < boardCols; col++ {

			location := NewLocation(col, row)
			piece := board.PieceAt(&location)

			if !piece.IsEmpty() && piece.IsOwnedBy(&player) {
				currentMoves := board.FindMovesForPlayersPieceAtLocation(piece, player, location)
				totalMoves = append(totalMoves, currentMoves...)
			}
		}
	}

	return totalMoves
}

func (board *Board) LegalMovesForPlayerMulti(player Player) *Moves {
	totalMoves := Moves{}
	var countGoRoutines int8 = 0
	var movesChan = make(chan Moves)
	var row, col, i int8

	for row = 0; row < boardRows; row++ {
		for col = 0; col < boardCols; col++ {

			location := NewLocation(col, row)
			piece := board.PieceAt(&location)

			if !piece.IsEmpty() && piece.IsOwnedBy(&player) {

				// Create goRoutines so we can quickly find all the moves
				countGoRoutines++
				go func() {
					currentMoves := board.FindMovesForPlayersPieceAtLocation(piece, player, location)
					movesChan <- currentMoves
				}()
			}
		}
	}
	for i = 0; i < countGoRoutines; i++ {
		totalMoves = append(totalMoves, <-movesChan...)
	}

	return &totalMoves
}

func (board *Board) FindMovesForPlayersPieceAtLocation(piece Piece, player Player, location Location) Moves {

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
	moves := Moves{}
	return moves
}

func (board *Board) FindMovesForBishopAtLocation(player Player, originalLocation Location) Moves {
	moves := Moves{}

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

func (board *Board) FindMovesForRookAtLocation(player Player, originalLocation Location) Moves {
	moves := Moves{}

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

func (board *Board) FindMovesForKnightAtLocation(player Player, originalLocation Location) Moves {
	moves := Moves{}

	applyKnightMove := func(colsToAppendBy, rowsToAppendBy int8) {
		isMovingBackward := (rowsToAppendBy < 0 && player == HUMAN) || (rowsToAppendBy > 0 && player == GOBOT)

		move := Move{from: originalLocation, to: originalLocation.Append(colsToAppendBy, rowsToAppendBy)}
		piece := board.PieceAt(&move.to)
		move.weight = piece.MoveWeight()

		isValidMove := board.isValidMove(&move, &player)
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

func (board *Board) FindMovesForPawnAtLocation(player Player, originalLocation Location) Moves {
	moves := Moves{}
	var move Move

	// Get direction
	var direction int8 = 1
	if player == GOBOT {
		direction *= -1
	}

	// Forward move
	move = Move{from: originalLocation, to: originalLocation.Append(0, direction)}
	piece := board.PieceAt(&move.to)
	move.weight = piece.MoveWeight()

	if piece.IsEmpty() {
		moves = append(moves, move)
	}

	// Attack moves
	applyAttackMove := func(colsToAppendBy, rowsToAppendBy int8) {
		move = Move{from: originalLocation, to: originalLocation.Append(colsToAppendBy, rowsToAppendBy)}
		piece := board.PieceAt(&move.to)
		move.weight = piece.MoveWeight()
		if piece.IsOwnedBy(player.Opponent()) {
			moves = append(moves, move)
		}
	}

	applyAttackMove(1, direction)
	applyAttackMove(-1, direction)

	return moves
}

func (board *Board) FindMovesForKingAtLocation(player Player, originalLocation Location) Moves {
	moves := Moves{}

	// Get direction
	var direction int8 = 1
	if player == GOBOT {
		direction *= -1
	}

	move := Move{from: originalLocation, to: originalLocation.Append(-1*direction, 0)}
	piece := board.PieceAt(&move.to)
	move.weight = piece.MoveWeight()

	if piece.IsEmpty() {
		moves = append(moves, move)
	}

	move = Move{from: originalLocation, to: originalLocation.Append(1, 0)}
	piece = board.PieceAt(&move.to)
	move.weight = piece.MoveWeight()
	if piece.IsOwnedBy(player.Opponent()) {
		moves = append(moves, move)
	}

	move = Move{from: originalLocation, to: originalLocation.Append(-1, 0)}
	piece = board.PieceAt(&move.to)
	move.weight = piece.MoveWeight()
	if piece.IsOwnedBy(player.Opponent()) {
		moves = append(moves, move)
	}

	return moves
}

func (board *Board) getDirectionalMoves(player Player, originalLocation Location, colsToAppendBy int8, rowsToAppendBy int8) Moves {
	moves := Moves{}
	moveLocation := originalLocation
	movingBackward := (rowsToAppendBy < 0 && player == HUMAN) || (rowsToAppendBy > 0 && player == GOBOT)
	moveToNextLocation := func() { moveLocation = moveLocation.Append(colsToAppendBy, rowsToAppendBy) }

	moveToNextLocation()
	for moveLocation.IsOnBoard() {
		pieceToMoveTo := board.PieceAt(&moveLocation)

		if movingBackward && pieceToMoveTo.IsEmpty() {
			moveToNextLocation()
			continue
		}
		if pieceToMoveTo.IsOwnedBy(&player) {
			break // Don't do the move
		}

		moves = append(moves, Move{from: originalLocation, to: moveLocation, weight: pieceToMoveTo.MoveWeight()})

		if pieceToMoveTo.IsOwnedBy(player.Opponent()) {
			break
		}
		moveToNextLocation()
	}
	return moves
}

func (board *Board) isValidMove(move *Move, player *Player) bool {
	piece := board.PieceAt(&move.to)
	return !piece.IsOwnedBy(player) && move.to.IsOnBoard()
}
