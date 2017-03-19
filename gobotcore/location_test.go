package gobotcore

import (
	"testing"
	"strings"
)

func TestLocation_isOnBoard(t *testing.T) {
	location := Location{ col: boardCols, row: 0}
	if location.IsOnBoard() {
		t.Error("Location is off board")
	}
	location = Location{col: 0, row: boardRows}
	if location.IsOnBoard() {
		t.Error("Location is off board")
	}
	location = Location{col: boardCols - 1, row: -1}
	if location.IsOnBoard() {
		t.Error("Location is off board")
	}
	location = Location{col: -1, row: boardRows - 1}
	if location.IsOnBoard() {
		t.Error("Location is off board")
	}
}

func TestLocation_equals(t *testing.T) {
	loc1 := Location{col: 3, row: 2 }
	loc2 := Location{col: 3, row: 2 }
	if !loc1.Equals(loc2) {
		t.Error("Locations should be equal")
	}
}

func TestNewLocation(t *testing.T) {
	loc1 := Location{col: 3, row: 2 }
	loc2 := NewLocation(3, 2)
	if !loc1.Equals(loc2) {
		t.Error("Locations should be equal")
	}
}

func TestNewStringFromLocation(t *testing.T) {
	location := NewLocation(2, 2)
	locationStr := location.ToString()
	if strings.Compare("C3", locationStr) != 0 {
		t.Error("locationStr should be equal to C3")
	}
}

func TestNewStringFromLocations(t *testing.T) {
	location1 := NewLocation(5, 2)
	location2 := NewLocation(1, 4 )
	str := ToStringMultipleLocations(location1, location2)
	if strings.Compare("C6E2", str) != 0 {
		t.Error("String is incorrect")
	}
}

func TestNewLocationFromString(t *testing.T) {
	loc1 := NewLocation(2, 2)
	loc2 := NewLocationFromString("C3")
	if !loc1.Equals(loc2) {
		t.Error("NewLocationFromString returning wrong value. Row: " + string(loc2.row) + " Col: " + string(loc2.col))
	}
}

func TestNewLocationsFromStrings(t *testing.T) {
	loc1, loc2 := NewLocationsFromString("A2C3")
	if !loc1.Equals(NewLocation(0, 1)) {
		t.Error("Location 1 incorrect")
	}
	if !loc2.Equals(NewLocation(2, 2)) {
		t.Error("Location 2 incorrect")
	}
}

func TestLocation_Append(t *testing.T) {
	location1 := NewLocation(4, 2)
	location2 := NewLocation(5, 3 )
	location1 = location1.Append(1,1)
	if !location1.Equals(location2) {
		t.Error("Location append failed")
	}
}