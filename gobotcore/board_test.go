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
	buffer.WriteString("0   - - - - - -\n")
	buffer.WriteString("\n")
	buffer.WriteString("    A B C D E F")

	board := NewBoardFromString(buffer.String())
	moves := board.GetMovesForPlayer(GOBOT)
	bishopLocation := Location{col: 2, row: 1}

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
	if !NewMove(bishopLocation, bishopLocation.Append(1, -1)).IsContainedIn(moves) {
		t.Error("Move is valid")
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
	buffer.WriteString("1   - - R - - -\n")
	buffer.WriteString("0   - - - - - -\n")
	buffer.WriteString("\n")
	buffer.WriteString("    A B C D E F")

	board := NewBoardFromString(buffer.String())
	moves := board.GetMovesForPlayer(GOBOT)
	rookLocation := Location{col: 2, row: 1}

	if !NewMove(rookLocation, rookLocation.Append(1, 0)).IsContainedIn(moves) {
		t.Error("Move is valid")
	}
	if !NewMove(rookLocation, rookLocation.Append(3, 0)).IsContainedIn(moves) {
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
