package gobotcore

import (
	"testing"
	"bytes"
)

func TestNewBoardFromString(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("7   - K - - - -\n")
	buffer.WriteString("6   N B R R B N\n")
	buffer.WriteString("5   - - P P - -\n")
	buffer.WriteString("4   - - - - - -\n")
	buffer.WriteString("3   - - - - - -\n")
	buffer.WriteString("2   - - p p - -\n")
	buffer.WriteString("1   n n r r b n\n")
	buffer.WriteString("0   - - - - k -\n")
	buffer.WriteString("\n")
	buffer.WriteString("    A B C D E F")
	if NewBoardFromString(buffer.String()) != NewDefaultBoard() {
		t.Error("String board should match default board")
	}
}

func TestBoard_FindMovesForBishopAtLocation(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("7   - - - - - -\n")
	buffer.WriteString("6   - - - - - -\n")
	buffer.WriteString("5   - - - - - -\n")
	buffer.WriteString("4   - - - - - -\n")
	buffer.WriteString("3   - p p p p p\n")
	buffer.WriteString("2   - - - - - -\n")
	buffer.WriteString("1   - - B - - -\n")
	buffer.WriteString("0   - - - K - -\n")
	buffer.WriteString("\n")
	buffer.WriteString("    A B C D E F")

	board := NewBoardFromString(buffer.String())
	bishopLocation := Location{col: 2, row: 1}
	moves := board.FindMovesForBishopAtLocation(GOBOT, bishopLocation)

	if len(moves) == 0 {
		t.Error("Returned no moves")
	}
	if !NewMove(bishopLocation, bishopLocation.Append(1, 1)).IsContainedIn(moves) {
		t.Error("Move is valid")
	}
	if !NewMove(bishopLocation, bishopLocation.Append(2, 2)).IsContainedIn(moves) {
		t.Error("Move should be able to capture another piece")
	}
	if !NewMove(bishopLocation, bishopLocation.Append(-1, 1)).IsContainedIn(moves) {
		t.Error("Move is valid")
	}
	if !NewMove(bishopLocation, bishopLocation.Append(-1, -1)).IsContainedIn(moves) {
		t.Error("Move is valid")
	}
	if NewMove(bishopLocation, bishopLocation.Append(1, -1)).IsContainedIn(moves) {
		t.Error("Cannot move to owned piece")
	}
	if NewMove(bishopLocation, bishopLocation.Append(3, 3)).IsContainedIn(moves) {
		t.Error("Move should not be able to go past a piece after it gets captured")
	}
	if NewMove(bishopLocation, bishopLocation.Append(-3, 3)).IsContainedIn(moves) {
		t.Error("Move is outside of board")
	}
	if NewMove(bishopLocation, bishopLocation.Append(-1, 2)).IsContainedIn(moves) {
		t.Error("Move is invalid")
	}
}

func TestBoard_FindMovesForRookAtLocation(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("7   - - - - - -\n")
	buffer.WriteString("6   - - - - - -\n")
	buffer.WriteString("5   - - - - - -\n")
	buffer.WriteString("4   - - p - - -\n")
	buffer.WriteString("3   p p - p p p\n")
	buffer.WriteString("2   - - - - - -\n")
	buffer.WriteString("1   - - R - - K\n")
	buffer.WriteString("0   - - - - - -\n")
	buffer.WriteString("\n")
	buffer.WriteString("    A B C D E F")

	board := NewBoardFromString(buffer.String())
	rookLocation := Location{col: 2, row: 1}
	moves := board.FindMovesForRookAtLocation(GOBOT, rookLocation)

	if len(moves) == 0 {
		t.Error("Returned no moves")
	}
	if !NewMove(rookLocation, rookLocation.Append(1, 0)).IsContainedIn(moves) {
		t.Error("Move is valid")
	}
	if NewMove(rookLocation, rookLocation.Append(3, 0)).IsContainedIn(moves) {
		t.Error("Cannot capture owned piece")
	}
	if !NewMove(rookLocation, rookLocation.Append(2, 0)).IsContainedIn(moves) {
		t.Error("Move is valid")
	}
	if !NewMove(rookLocation, rookLocation.Append(0, 2)).IsContainedIn(moves) {
		t.Error("Move is valid")
	}
	if !NewMove(rookLocation, rookLocation.Append(0, 3)).IsContainedIn(moves) {
		t.Error("Move should be able capture an opponent")
	}
	if NewMove(rookLocation, rookLocation.Append(0, 4)).IsContainedIn(moves) {
		t.Error("Move should not continue after capturing an opponent")
	}
	if !NewMove(rookLocation, rookLocation.Append(-2, 0)).IsContainedIn(moves) {
		t.Error("Move is valid")
	}
	if NewMove(rookLocation, rookLocation.Append(-3, 0)).IsContainedIn(moves) {
		t.Error("Move should not be able to move outside board")
	}
	if !NewMove(rookLocation, rookLocation.Append(0, -1)).IsContainedIn(moves) {
		t.Error("Move is valid")
	}
	if NewMove(rookLocation, rookLocation.Append(0, -2)).IsContainedIn(moves) {
		t.Error("Move should not be able to move outside board")
	}
}

func TestBoard_FindMovesForKnightAtLocation(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("7   - - - - - -\n")
	buffer.WriteString("6   - - - - - -\n")
	buffer.WriteString("5   - - - - - -\n")
	buffer.WriteString("4   - N - - - -\n")
	buffer.WriteString("3   - - - p - -\n")
	buffer.WriteString("2   - - K - - -\n")
	buffer.WriteString("1   - - - - - -\n")
	buffer.WriteString("0   - - - - - -\n")
	buffer.WriteString("\n")
	buffer.WriteString("    A B C D E F")

	board := NewBoardFromString(buffer.String())
	knightLocation := Location{col: 1, row: 4}
	moves := board.FindMovesForKnightAtLocation(GOBOT, knightLocation)

	if len(moves) == 0 {
		t.Error("Returned no moves")
	}
	if !NewMove(knightLocation, knightLocation.Append(2, 1)).IsContainedIn(moves) {
		t.Error("Move is valid")
	}
	if !NewMove(knightLocation, knightLocation.Append(-1, 2)).IsContainedIn(moves) {
		t.Error("Move is valid")
	}
	if !NewMove(knightLocation, knightLocation.Append(1, 2)).IsContainedIn(moves) {
		t.Error("Move is valid")
	}
	if NewMove(knightLocation, knightLocation.Append(-2, 1)).IsContainedIn(moves) {
		t.Error("Move is off board")
	}
	if NewMove(knightLocation, knightLocation.Append(-2, -1)).IsContainedIn(moves) {
		t.Error("Move is off board")
	}
	if !NewMove(knightLocation, knightLocation.Append(-1, -2)).IsContainedIn(moves) {
		t.Error("Move is valid")
	}
	if NewMove(knightLocation, knightLocation.Append(1, -2)).IsContainedIn(moves) {
		t.Error("Cannot capture owned piece")
	}
	if !NewMove(knightLocation, knightLocation.Append(2, -1)).IsContainedIn(moves) {
		t.Error("Should be able capture an opponent")
	}
}

func TestBoard_FindMovesForPawnAtLocation(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("7   - - p - - -\n")
	buffer.WriteString("6   - - - - - -\n")
	buffer.WriteString("5   - - - - - -\n")
	buffer.WriteString("4   P - - - K -\n")
	buffer.WriteString("3   p n - p - -\n")
	buffer.WriteString("2   p - - - - -\n")
	buffer.WriteString("1   - - - - - -\n")
	buffer.WriteString("0   - - - - - -\n")
	buffer.WriteString("\n")
	buffer.WriteString("    A B C D E F")

	board := NewBoardFromString(buffer.String())

	pawnLocation := Location{col: 3, row: 3}
	moves := board.FindMovesForPawnAtLocation(HUMAN, pawnLocation)
	if len(moves) == 0 {
		t.Error("Returned no moves")
	}
	if !NewMove(pawnLocation, pawnLocation.Append(1, 1)).IsContainedIn(moves) {
		t.Error("Move is valid")
	}
	if !NewMove(pawnLocation, pawnLocation.Append(0, 1)).IsContainedIn(moves) {
		t.Error("Move is valid")
	}
	if NewMove(pawnLocation, pawnLocation.Append(-1, 1)).IsContainedIn(moves) {
		t.Error("Cannot capture empty piece")
	}

	pawnLocation = Location{col: 2, row: 7}
	moves = board.FindMovesForPawnAtLocation(HUMAN, pawnLocation)
	if NewMove(pawnLocation, pawnLocation.Append(0, 1)).IsContainedIn(moves) {
		t.Error("Move is off board")
	}

	pawnLocation = Location{col: 0, row: 2}
	moves = board.FindMovesForPawnAtLocation(HUMAN, pawnLocation)
	if NewMove(pawnLocation, pawnLocation.Append(0, 1)).IsContainedIn(moves) {
		t.Error("Cannot move to held location")
	}
	if NewMove(pawnLocation, pawnLocation.Append(1, 1)).IsContainedIn(moves) {
		t.Error("Cannot capture own piece")
	}
	if NewMove(pawnLocation, pawnLocation.Append(-1, 1)).IsContainedIn(moves) {
		t.Error("Cannot move off board")
	}

	pawnLocation = Location{col: 0, row: 4}
	moves = board.FindMovesForPawnAtLocation(GOBOT, pawnLocation)
	if len(moves) == 0 {
		t.Error("Returned no moves")
	}
	if NewMove(pawnLocation, pawnLocation.Append(-1, 1)).IsContainedIn(moves) {
		t.Error("Move is valid")
	}
}

func TestBoard_FindMovesForKingAtLocation(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("7   - K - - - -\n")
	buffer.WriteString("6   N B R R B N\n")
	buffer.WriteString("5   - - P P - -\n")
	buffer.WriteString("4   - - - - - -\n")
	buffer.WriteString("3   - - - - - -\n")
	buffer.WriteString("2   - - p p - -\n")
	buffer.WriteString("1   n n r r b n\n")
	buffer.WriteString("0   - - - - k -\n")
	buffer.WriteString("\n")
	buffer.WriteString("    A B C D E F")

	board := NewBoardFromString(buffer.String())

	kingLocation := Location{col: 4, row: 0}
	moves := board.FindMovesForKingAtLocation(HUMAN, kingLocation)
	if len(moves) == 0 {
		t.Error("Returned no moves")
	}
	if !NewMove(kingLocation, kingLocation.Append(-1, 0)).IsContainedIn(moves) {
		t.Error("Move is valid")
	}
	if NewMove(kingLocation, kingLocation.Append(1, 0)).IsContainedIn(moves) {
		t.Error("Move is not valid")
	}

	kingLocation = Location{col: 1, row: 7}
	moves = board.FindMovesForKingAtLocation(GOBOT, kingLocation)
	if len(moves) == 0 {
		t.Error("Returned no moves")
	}
	if !NewMove(kingLocation, kingLocation.Append(1, 0)).IsContainedIn(moves) {
		t.Error("Move is valid")
	}
	if NewMove(kingLocation, kingLocation.Append(-1, 0)).IsContainedIn(moves) {
		t.Error("Move is not valid")
	}
}
