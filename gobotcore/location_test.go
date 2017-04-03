package gobotcore

import (
	"strings"
	"testing"
)

func TestLocation_isOnBoard(t *testing.T) {
	location := Location{col: boardCols, row: 0}
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
	loc1 := Location{col: 3, row: 2}
	loc2 := Location{col: 3, row: 2}
	if !loc1.Equals(&loc2) {
		t.Error("Locations should be equal")
	}
}

func TestNewLocation(t *testing.T) {
	loc1 := Location{col: 5, row: 7}
	loc2 := NewLocation(5, 7)
	if !loc1.Equals(&loc2) {
		t.Error("Locations should be equal")
	}
}

func TestNewStringFromLocation(t *testing.T) {
	location := NewLocation(5, 7)
	locationStr := location.ToString()
	if strings.Compare("F8", locationStr) != 0 {
		t.Error("locationStr should be equal to F8, is actually " + locationStr)
	}
}

func TestNewStringFromLocations(t *testing.T) {
	location1 := NewLocation(5, 7)
	location2 := NewLocation(0, 0)
	str := ToStringMultipleLocations(location1, location2)
	if strings.Compare("F8A1", str) != 0 {
		t.Error("String is incorrect, is actually " + str)
	}
}

func TestNewStringFromLocationsFlipped(t *testing.T) {
	location1 := NewLocation(5, 7)
	location2 := NewLocation(0, 0)
	str := ToStringMultipleLocationsFlipped(location1, location2)
	if strings.Compare("A1F8", str) != 0 {
		t.Error("String is incorrect, is actually " + str)
	}
}

func TestNewLocationFromString(t *testing.T) {
	loc1 := NewLocation(5, 7)
	loc2 := NewLocationFromString("F8")
	if !loc1.Equals(&loc2) {
		t.Error("NewLocationFromString returning wrong value. Row: " + string(loc2.row) + " Col: " + string(loc2.col))
	}
}

func TestNewLocationsFromStrings(t *testing.T) {
	loc1, loc2 := NewLocationsFromString("A2C3")
	if !loc1.Equals(&Location{0, 1}) {
		t.Error("Location 1 incorrect")
	}
	if !loc2.Equals(&Location{2, 2}) {
		t.Error("Location 2 incorrect")
	}
}

func TestLocation_Append(t *testing.T) {
	location1 := Location{4, 2}
	location2 := &Location{5, 3}
	location1 = location1.Append(1, 1)
	if !location1.Equals(location2) {
		t.Error("Location append failed")
	}
}
