package gobot

import (
	"testing"
)

func TestLocation_isOnBoard(t *testing.T) {
	location := location{row: 0, col: boardCols}
	if location.isOnBoard() {
		t.Error("Location is off board")
	}
	location = location{row: boardRows, col: 0}
	if location.isOnBoard() {
		t.Error("Location is off board")
	}
	location = location{row: -1, col: boardCols-1}
	if location.isOnBoard() {
		t.Error("Location is off board")
	}
	location = location{row: boardRows-1, col: -1}
	if location.isOnBoard() {
		t.Error("Location is off board")
	}
}

func TestLocation_equals(t *testing.T) {
	loc1 := location{row: 2, col: 3}
	loc2 := location{row: 2, col: 3}
	if !loc1.equals(loc2) {
		t.Error("Locations should be equal")
	}
}