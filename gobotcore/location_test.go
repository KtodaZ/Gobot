package gobotcore

import (
	"testing"
)

func TestLocation_isOnBoard(t *testing.T) {
	location := Location{row: 0, col: boardCols}
	if location.isOnBoard() {
		t.Error("Location is off board")
	}
	location = Location{row: boardRows, col: 0}
	if location.isOnBoard() {
		t.Error("Location is off board")
	}
	location = Location{row: -1, col: boardCols-1}
	if location.isOnBoard() {
		t.Error("Location is off board")
	}
	location = Location{row: boardRows-1, col: -1}
	if location.isOnBoard() {
		t.Error("Location is off board")
	}
}

func TestLocation_equals(t *testing.T) {
	loc1 := Location{row: 2, col: 3}
	loc2 := Location{row: 2, col: 3}
	if !loc1.equals(loc2) {
		t.Error("Locations should be equal")
	}
}