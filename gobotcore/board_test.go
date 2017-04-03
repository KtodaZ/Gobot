package gobotcore

import (
	"bytes"
	"runtime"
	"testing"
)

func TestNewBoardFromString(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("8   - K - - - -\n")
	buffer.WriteString("7   N B R R B N\n")
	buffer.WriteString("6   - - P P - -\n")
	buffer.WriteString("5   - - - - - -\n")
	buffer.WriteString("4   - - - - - -\n")
	buffer.WriteString("3   - - p p - -\n")
	buffer.WriteString("2   n b r r b n\n")
	buffer.WriteString("1   - - - - k -\n")
	buffer.WriteString("\n")
	buffer.WriteString("    A B C D E F")
	if NewBoardFromString(buffer.String()) != NewDefaultBoard() {
		t.Error("String board should match default board")
	}
}

var depth int8 = 7
/*func TestBoard_Minimax(t *testing.T) {
	board := NewDefaultBoard()
	player := Player(GOBOT)
	move := board.Minimax(&player, &depth)
	board.MakeMoveAndPrintMessage(&move)
}*/
func TestBoard_MinimaxMulti(t *testing.T) {
	board := NewDefaultBoard()
	player := Player(GOBOT)
	move := board.MinimaxMulti(&player, &depth)
	board.MakeMoveAndPrintMessage(&move)
}

func TestBoard_Minimax2(t *testing.T) {
	SetDebug(false)
	var buffer bytes.Buffer
	buffer.Reset()
	buffer.WriteString("8   - - - - - -\n")
	buffer.WriteString("7   - - - - - -\n")
	buffer.WriteString("6   - - - - - -\n")
	buffer.WriteString("5   - - - - - -\n")
	buffer.WriteString("4   - - - - - -\n")
	buffer.WriteString("3   - - - - - -\n")
	buffer.WriteString("2   - - B - - -\n")
	buffer.WriteString("1   - r - k - -\n")
	buffer.WriteString("\n")
	buffer.WriteString("    A B C D E F")
	board := NewBoardFromString(buffer.String())
	player := Player(GOBOT)
	move := board.MinimaxMulti(&player, &depth)
	moveExpected := Move{from: Location{2, 1}, to: Location{3, 0}}
	if !move.Equals(&moveExpected) {
		t.Error("Move " + move.ToString() + " should equal expected: " + moveExpected.ToString())
	}
}

var benchMove Move
/*func BenchmarkBoard_Minimax(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	board := NewDefaultBoard()
	var benchMoveTemp Move
	SetDebug(false)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		player := Player(GOBOT)
		benchMoveTemp = board.Minimax(&player, &depth)
	}
	benchMove = benchMoveTemp
}*/
func BenchmarkBoard_MinimaxMulti(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	board := NewDefaultBoard()
	var benchMoveTemp Move
	SetDebug(false)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		player := Player(GOBOT)
		benchMoveTemp = board.MinimaxMulti(&player, &depth)
	}
	benchMove = benchMoveTemp
}

func TestBoard_MakeMove(t *testing.T) {
	var buffer bytes.Buffer
	buffer.Reset()
	buffer.WriteString("8   - - - - - -\n")
	buffer.WriteString("7   - - - - - -\n")
	buffer.WriteString("6   - - - - - -\n")
	buffer.WriteString("5   - - - - - -\n")
	buffer.WriteString("4   - - - - - -\n")
	buffer.WriteString("3   - - - - - -\n")
	buffer.WriteString("2   - - - - - -\n")
	buffer.WriteString("1   - - - k - -\n")
	buffer.WriteString("\n")
	buffer.WriteString("    A B C D E F")
	boardBefore := NewBoardFromString(buffer.String())

	buffer.Reset()
	buffer.WriteString("8   - - - - - -\n")
	buffer.WriteString("7   - - - - - -\n")
	buffer.WriteString("6   - - - - - -\n")
	buffer.WriteString("5   - - - - - -\n")
	buffer.WriteString("4   - - - - - -\n")
	buffer.WriteString("3   - - - - - -\n")
	buffer.WriteString("2   - - - - - -\n")
	buffer.WriteString("1   - - k - - -\n")
	buffer.WriteString("\n")
	buffer.WriteString("    A B C D E F")
	boardAfter := NewBoardFromString(buffer.String())

	move := Move{from: Location{col: 3, row: 0}, to: Location{col: 2, row: 0}}
	boardBefore.MakeMoveAndGetTakenPiece(&move)

	if boardBefore != boardAfter {
		t.Error("Boards should match")
	}

}

func TestBoard_LegalMovesForPlayer(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("8   - K - - - -\n")
	buffer.WriteString("7   N B R R B N\n")
	buffer.WriteString("6   - - P - - -\n")
	buffer.WriteString("5   - - - - - -\n")
	buffer.WriteString("4   - - - - - -\n")
	buffer.WriteString("3   - - - p - -\n")
	buffer.WriteString("2   n b r r b n\n")
	buffer.WriteString("1   - - - - k -\n")
	buffer.WriteString("\n")
	buffer.WriteString("    A B C D E F")

	board := NewBoardFromString(buffer.String())
	boardCopy := board
	moves := board.LegalMovesForPlayer(HUMAN)

	if board != boardCopy {
		board.PrintBoard()
		boardCopy.PrintBoard()
		t.Error("Boards changed after legalmoves")
	}
	if len(moves) != 15 {
		t.Error("Returned wrong move size : " + string(len(moves)))
	}
	bishopLocation := Location{col: 4, row: 1}
	if !NewMove(bishopLocation, bishopLocation.Append(1, 1)).IsContainedIn(&moves) {
		t.Error("Move is valid")
	}
	knightLocation := Location{col: 0, row: 1}
	if !NewMove(knightLocation, knightLocation.Append(1, 2)).IsContainedIn(&moves) {
		t.Error("Move is valid")
	}
	rookLocation := Location{col: 2, row: 1}
	if !NewMove(rookLocation, rookLocation.Append(0, 1)).IsContainedIn(&moves) {
		t.Error("Move is valid")
	}
	pawnLocation := Location{col: 3, row: 2}
	if !NewMove(pawnLocation, pawnLocation.Append(0, 1)).IsContainedIn(&moves) {
		t.Error("Move is valid")
	}
	kingLocation := Location{col: 4, row: 0}
	if !NewMove(kingLocation, kingLocation.Append(-1, 0)).IsContainedIn(&moves) {
		t.Error("Move is valid")
	}

	moves = board.LegalMovesForPlayer(GOBOT)
	bishopLocation = Location{col: 4, row: 6}
	if !NewMove(bishopLocation, bishopLocation.Append(1, -1)).IsContainedIn(&moves) {
		t.Error("Move is valid")
	}
	knightLocation = Location{col: 5, row: 6}
	if !NewMove(knightLocation, knightLocation.Append(-1, -2)).IsContainedIn(&moves) {
		t.Error("Move is valid")
	}
	rookLocation = Location{col: 3, row: 6}
	if !NewMove(rookLocation, rookLocation.Append(0, -1)).IsContainedIn(&moves) {
		t.Error("Move is valid")
	}
	pawnLocation = Location{col: 2, row: 5}
	if !NewMove(pawnLocation, pawnLocation.Append(0, -1)).IsContainedIn(&moves) {
		t.Error("Move is valid")
	}
	kingLocation = Location{col: 1, row: 7}
	if !NewMove(kingLocation, kingLocation.Append(1, 0)).IsContainedIn(&moves) {
		t.Error("Move is valid")
	}
}

func TestBoard_LegalMovesForPlayerSort(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("8   - - - - - -\n")
	buffer.WriteString("7   - - - - - -\n")
	buffer.WriteString("6   - - - - - -\n")
	buffer.WriteString("5   - n - - - -\n")
	buffer.WriteString("4   - - - P - -\n")
	buffer.WriteString("3   - - K - - -\n")
	buffer.WriteString("2   - - - - - -\n")
	buffer.WriteString("1   - - - - - -\n")
	buffer.WriteString("\n")
	buffer.WriteString("    A B C D E F")

	board := NewBoardFromString(buffer.String())
	moves := board.LegalMovesForPlayer(HUMAN)

	if moves[0].weight != 0 {
		t.Error("Should make king capturing moves first")
	}
	if moves[1].weight != 1 {
		t.Error("Should make capturing moves second")
	}
	if moves[2].weight != 2 {
		t.Error("Should make regular moves third")
	}
}

var result Moves

func BenchmarkBoard_LegalMovesForPlayer(b *testing.B) {
	board := NewDefaultBoard()
	var r Moves
	b.ResetTimer()
	// run the test function b.N times
	for n := 0; n < b.N; n++ {
		player := Player(HUMAN)
		r = board.LegalMovesForPlayer(player)
	}
	result = r
}
func BenchmarkBoard_LegalMovesForPlayerMulti(b *testing.B) {
	board := NewDefaultBoard()
	var r Moves
	b.ResetTimer()
	// run the test function b.N times
	for n := 0; n < b.N; n++ {
		player := Player(HUMAN)
		r = *board.LegalMovesForPlayerMulti(player)
	}
	result = r
}

func TestBoard_FindMovesForBishopAtLocation(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("8   - - - K - -\n")
	buffer.WriteString("7   - - B K - -\n")
	buffer.WriteString("6   - - - - - -\n")
	buffer.WriteString("5   - - p - - -\n")
	buffer.WriteString("4   - P P P P P\n")
	buffer.WriteString("3   - - - - - -\n")
	buffer.WriteString("2   - - b - - -\n")
	buffer.WriteString("1   - P - - - -\n")
	buffer.WriteString("\n")
	buffer.WriteString("    A B C D E F")

	board := NewBoardFromString(buffer.String())
	bishopLocation := Location{col: 2, row: 1}
	player := Player(HUMAN)
	moves := board.FindMovesForBishopAtLocation(player, bishopLocation)

	if len(moves) != 5 {
		t.Error("Returned wrong move size : " + string(len(moves)))
	}
	if !NewMove(bishopLocation, bishopLocation.Append(1, 1)).IsContainedIn(&moves) {
		t.Error("Move is valid")
	}
	if !NewMove(bishopLocation, bishopLocation.Append(2, 2)).IsContainedIn(&moves) {
		t.Error("Move should be able to capture another piece")
	}
	if !NewMove(bishopLocation, bishopLocation.Append(-1, 1)).IsContainedIn(&moves) {
		t.Error("Move is valid")
	}
	if !NewMove(bishopLocation, bishopLocation.Append(-1, -1)).IsContainedIn(&moves) {
		t.Error("Move is valid")
	}
	if NewMove(bishopLocation, bishopLocation.Append(1, -1)).IsContainedIn(&moves) {
		t.Error("Move is invalid")
	}
	if NewMove(bishopLocation, bishopLocation.Append(3, 3)).IsContainedIn(&moves) {
		t.Error("Move should not be able to go past a piece after it gets captured")
	}
	if NewMove(bishopLocation, bishopLocation.Append(-3, 3)).IsContainedIn(&moves) {
		t.Error("Move is outside of board")
	}
	if NewMove(bishopLocation, bishopLocation.Append(-1, 2)).IsContainedIn(&moves) {
		t.Error("Move is invalid")
	}
	bishopLocation = Location{col: 2, row: 6}
	moves = board.FindMovesForBishopAtLocation(GOBOT, bishopLocation)
	if NewMove(bishopLocation, bishopLocation.Append(1, 1)).IsContainedIn(&moves) {
		t.Error("Move is invalid")
	}
}

func TestBoard_FindMovesForRookAtLocation(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("8   - - - - R -\n")
	buffer.WriteString("7   p - R - N -\n")
	buffer.WriteString("6   - - - - - -\n")
	buffer.WriteString("5   - - P - - -\n")
	buffer.WriteString("4   P P - P P P\n")
	buffer.WriteString("3   - - - - - -\n")
	buffer.WriteString("2   - - r - - k\n")
	buffer.WriteString("1   - - - - - -\n")
	buffer.WriteString("\n")
	buffer.WriteString("    A B C D E F")

	board := NewBoardFromString(buffer.String())
	rookLocation := Location{col: 2, row: 1}
	moves := board.FindMovesForRookAtLocation(HUMAN, rookLocation)

	if len(moves) != 7 {
		t.Error("Returned wrong move size : " + string(len(moves)))
	}
	if !NewMove(rookLocation, rookLocation.Append(1, 0)).IsContainedIn(&moves) {
		t.Error("Move is valid")
	}
	if NewMove(rookLocation, rookLocation.Append(3, 0)).IsContainedIn(&moves) {
		t.Error("Cannot capture owned piece")
	}
	if !NewMove(rookLocation, rookLocation.Append(2, 0)).IsContainedIn(&moves) {
		t.Error("Move is valid")
	}
	if !NewMove(rookLocation, rookLocation.Append(0, 2)).IsContainedIn(&moves) {
		t.Error("Move is valid")
	}
	if !NewMove(rookLocation, rookLocation.Append(0, 3)).IsContainedIn(&moves) {
		t.Error("Move should be able capture an opponent")
	}
	if NewMove(rookLocation, rookLocation.Append(0, 4)).IsContainedIn(&moves) {
		t.Error("Move should not continue after capturing an opponent")
	}
	if !NewMove(rookLocation, rookLocation.Append(-2, 0)).IsContainedIn(&moves) {
		t.Error("Move is valid")
	}
	if NewMove(rookLocation, rookLocation.Append(-3, 0)).IsContainedIn(&moves) {
		t.Error("Move should not be able to move outside board")
	}
	if NewMove(rookLocation, rookLocation.Append(0, -1)).IsContainedIn(&moves) {
		t.Error("Move is invalid")
	}
	if NewMove(rookLocation, rookLocation.Append(0, -2)).IsContainedIn(&moves) {
		t.Error("Move should not be able to move outside board")
	}

	rookLocation = Location{col: 2, row: 6}
	moves = board.FindMovesForRookAtLocation(GOBOT, rookLocation)
	if NewMove(rookLocation, rookLocation.Append(0, -2)).IsContainedIn(&moves) {
		t.Error("Move is invalid")
	}
	if !NewMove(rookLocation, rookLocation.Append(-2, 0)).IsContainedIn(&moves) {
		t.Error("Move is valid")
	}

	rookLocation = Location{col: 4, row: 7}
	moves = board.FindMovesForRookAtLocation(GOBOT, rookLocation)
	if NewMove(rookLocation, rookLocation.Append(0, -1)).IsContainedIn(&moves) {
		t.Error("Move is invalid")
	}
}

func TestBoard_FindMovesForKnightAtLocation(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("8   - - - - - -\n")
	buffer.WriteString("7   - - - - - -\n")
	buffer.WriteString("6   - - - - - -\n")
	buffer.WriteString("5   - n - - - -\n")
	buffer.WriteString("4   - - - P - -\n")
	buffer.WriteString("3   - - k - - -\n")
	buffer.WriteString("2   - - - - - -\n")
	buffer.WriteString("1   - - - - - -\n")
	buffer.WriteString("\n")
	buffer.WriteString("    A B C D E F")

	board := NewBoardFromString(buffer.String())
	knightLocation := Location{col: 1, row: 4}
	moves := board.FindMovesForKnightAtLocation(HUMAN, knightLocation)

	if len(moves) != 4 {
		t.Error("Returned wrong move size : " + string(len(moves)))
	}
	if !NewMove(knightLocation, knightLocation.Append(2, 1)).IsContainedIn(&moves) {
		t.Error("Move is valid")
	}
	if !NewMove(knightLocation, knightLocation.Append(-1, 2)).IsContainedIn(&moves) {
		t.Error("Move is valid")
	}
	if !NewMove(knightLocation, knightLocation.Append(1, 2)).IsContainedIn(&moves) {
		t.Error("Move is valid")
	}
	if NewMove(knightLocation, knightLocation.Append(-2, 1)).IsContainedIn(&moves) {
		t.Error("Move is off board")
	}
	if NewMove(knightLocation, knightLocation.Append(-2, -1)).IsContainedIn(&moves) {
		t.Error("Move is off board")
	}
	if NewMove(knightLocation, knightLocation.Append(-1, -2)).IsContainedIn(&moves) {
		t.Error("Move is invalid")
	}
	if NewMove(knightLocation, knightLocation.Append(1, -2)).IsContainedIn(&moves) {
		t.Error("Cannot capture owned piece")
	}
	if !NewMove(knightLocation, knightLocation.Append(2, -1)).IsContainedIn(&moves) {
		t.Error("Should be able capture an opponent")
	}

	buffer.Reset()
	buffer.WriteString("8   - - - - - -\n")
	buffer.WriteString("7   - - - - - -\n")
	buffer.WriteString("6   - - - - - -\n")
	buffer.WriteString("5   - - - K - -\n")
	buffer.WriteString("4   - - p - - -\n")
	buffer.WriteString("3   - - - - N -\n")
	buffer.WriteString("2   - - - - - -\n")
	buffer.WriteString("1   - - - - - -\n")
	buffer.WriteString("\n")
	buffer.WriteString("    A B C D E F")

	board = NewBoardFromString(buffer.String())
	knightLocation = Location{col: 4, row: 2}
	moves = board.FindMovesForKnightAtLocation(GOBOT, knightLocation)

	if len(moves) != 4 {
		t.Error("Returned wrong move size")
	}
	if !NewMove(knightLocation, knightLocation.Append(-2, -1)).IsContainedIn(&moves) {
		t.Error("Move is valid")
	}
	if !NewMove(knightLocation, knightLocation.Append(-1, -2)).IsContainedIn(&moves) {
		t.Error("Move is valid")
	}
	if !NewMove(knightLocation, knightLocation.Append(1, -2)).IsContainedIn(&moves) {
		t.Error("Move is valid")
	}
	if NewMove(knightLocation, knightLocation.Append(2, 1)).IsContainedIn(&moves) {
		t.Error("Move is off board")
	}
	if NewMove(knightLocation, knightLocation.Append(2, -1)).IsContainedIn(&moves) {
		t.Error("Move is off board")
	}
	if NewMove(knightLocation, knightLocation.Append(-1, 2)).IsContainedIn(&moves) {
		t.Error("Move is invalid")
	}
	if !NewMove(knightLocation, knightLocation.Append(-2, 1)).IsContainedIn(&moves) {
		t.Error("Should be able capture an opponent")
	}
	if NewMove(knightLocation, knightLocation.Append(-1, 2)).IsContainedIn(&moves) {
		t.Error("Cannot capture owned piece")
	}
}

func TestBoard_FindMovesForPawnAtLocation(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("8   - - p - - -\n")
	buffer.WriteString("7   - - - - - -\n")
	buffer.WriteString("6   - - - - - -\n")
	buffer.WriteString("5   P - - - K -\n")
	buffer.WriteString("4   p n - p - -\n")
	buffer.WriteString("3   p - - - - -\n")
	buffer.WriteString("2   - - - - - -\n")
	buffer.WriteString("1   - - - - - -\n")
	buffer.WriteString("\n")
	buffer.WriteString("    A B C D E F")

	board := NewBoardFromString(buffer.String())

	pawnLocation := Location{col: 3, row: 3}
	moves := board.FindMovesForPawnAtLocation(HUMAN, pawnLocation)
	if len(moves) != 2 {
		t.Error("Returned wrong move size : " + string(len(moves)))
	}
	if !NewMove(pawnLocation, pawnLocation.Append(1, 1)).IsContainedIn(&moves) {
		t.Error("Move is valid")
	}
	if !NewMove(pawnLocation, pawnLocation.Append(0, 1)).IsContainedIn(&moves) {
		t.Error("Move is valid")
	}
	if NewMove(pawnLocation, pawnLocation.Append(-1, 1)).IsContainedIn(&moves) {
		t.Error("Cannot capture empty piece")
	}

	pawnLocation = Location{col: 2, row: 7}
	moves = board.FindMovesForPawnAtLocation(HUMAN, pawnLocation)
	if NewMove(pawnLocation, pawnLocation.Append(0, 1)).IsContainedIn(&moves) {
		t.Error("Move is off board")
	}

	pawnLocation = Location{col: 0, row: 2}
	moves = board.FindMovesForPawnAtLocation(HUMAN, pawnLocation)
	if NewMove(pawnLocation, pawnLocation.Append(0, 1)).IsContainedIn(&moves) {
		t.Error("Cannot move to held location")
	}
	if NewMove(pawnLocation, pawnLocation.Append(1, 1)).IsContainedIn(&moves) {
		t.Error("Cannot capture own piece")
	}
	if NewMove(pawnLocation, pawnLocation.Append(-1, 1)).IsContainedIn(&moves) {
		t.Error("Cannot move off board")
	}

	pawnLocation = Location{col: 0, row: 4}
	moves = board.FindMovesForPawnAtLocation(GOBOT, pawnLocation)
 	if len(moves) == 0 {
		t.Error("Returned no moves")
	}
	if NewMove(pawnLocation, pawnLocation.Append(-1, 1)).IsContainedIn(&moves) {
		t.Error("Move is valid")
	}
}

func TestBoard_FindMovesForKingAtLocation(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("8   - K - - - -\n")
	buffer.WriteString("7   N B R R B N\n")
	buffer.WriteString("6   - - P P - -\n")
	buffer.WriteString("5   - - - - - -\n")
	buffer.WriteString("4   - - - - - -\n")
	buffer.WriteString("3   - - p p - -\n")
	buffer.WriteString("2   n n r r b n\n")
	buffer.WriteString("1   - - - - k -\n")
	buffer.WriteString("\n")
	buffer.WriteString("    A B C D E F")

	board := NewBoardFromString(buffer.String())

	kingLocation := Location{col: 4, row: 0}
	moves := board.FindMovesForKingAtLocation(HUMAN, kingLocation)
	if len(moves) == 0 {
		t.Error("Returned no moves")
	}
	if !NewMove(kingLocation, kingLocation.Append(-1, 0)).IsContainedIn(&moves) {
		t.Error("Move is valid")
	}
	if NewMove(kingLocation, kingLocation.Append(1, 0)).IsContainedIn(&moves) {
		t.Error("Move is not valid")
	}

	kingLocation = Location{col: 1, row: 7}
	moves = board.FindMovesForKingAtLocation(GOBOT, kingLocation)
	if len(moves) == 0 {
		t.Error("Returned no moves")
	}
	if !NewMove(kingLocation, kingLocation.Append(1, 0)).IsContainedIn(&moves) {
		t.Error("Move is valid")
	}
	if NewMove(kingLocation, kingLocation.Append(-1, 0)).IsContainedIn(&moves) {
		t.Error("Move is not valid")
	}

	buffer.Reset()
	buffer.WriteString("8   - - - - - -\n")
	buffer.WriteString("7   - - - - - -\n")
	buffer.WriteString("6   - - - - - -\n")
	buffer.WriteString("5   - - - - - -\n")
	buffer.WriteString("4   - - - - - -\n")
	buffer.WriteString("3   - - - - - -\n")
	buffer.WriteString("2   k - - - - -\n")
	buffer.WriteString("1   - R k R - -\n")
	buffer.WriteString("\n")
	buffer.WriteString("    A B C D E F")

	board = NewBoardFromString(buffer.String())
	kingLocation = Location{col: 2, row: 0}
	moves = board.FindMovesForKingAtLocation(HUMAN, kingLocation)
	if len(moves) != 2 {
		t.Error("Returned incorrect number of moves")
	}
	if !NewMove(kingLocation, kingLocation.Append(1, 0)).IsContainedIn(&moves) {
		t.Error("Move is valid")
	}
	if !NewMove(kingLocation, kingLocation.Append(-1, 0)).IsContainedIn(&moves) {
		t.Error("Move is valid")
	}

	kingLocation = Location{col: 0, row: 1}
	moves = board.FindMovesForKingAtLocation(HUMAN, kingLocation)
	if len(moves) != 0 {
		t.Error("Returned incorrect number of moves")
	}
}