package gobotcore

import (
	"strconv"
	"strings"
)

type Location struct {
	col int8
	row int8
}

type PieceLocation struct {
	piece    Piece
	location Location
}

var alphabet string = "ABCDEF"
var alphabetReversed string = "FEDCBA"

func NewLocation(col, row int8) Location {
	location := Location{col: col, row: row}
	if location.IsOnBoard() {
		return location
	} else {
		panic("Cannot create location out of board")
	}
}

func NewLocationFromString(readable string) Location {
	// readable should be in the form "A2" where A corresponds to a letter row, and 2 corresponds to a column
	if len(readable) != 2 {
		panic("Length of readable must be 2.")
	}

	var rowLetter string = readable[:1]
	var colNumber string = readable[1:]

	indexOfLetterInAlphabet := strings.Index(alphabet, strings.ToUpper(rowLetter))
	if indexOfLetterInAlphabet > -1 {
		return NewLocation(int8(indexOfLetterInAlphabet), AtoiEZPZ(colNumber)-1)
	} else {
		// Return a bad location. This should get caught by the calling function
		return Location{-1, -1}
	}
}

func NewLocationsFromString(fullReadable string) (Location, Location) {
	return NewLocationFromString(fullReadable[:2]), NewLocationFromString(fullReadable[2:])
}

func (location Location) ToString() string {
	return string(alphabet[location.col]) + strconv.Itoa(int(location.row+1))
}

func ToStringMultipleLocations(source Location, destination Location) string {
	return string(alphabet[source.col]) + strconv.Itoa(int(source.row+1)) + string(alphabet[destination.col]) + strconv.Itoa(int(destination.row+1))
}

func ToStringMultipleLocationsFlipped(source Location, destination Location) string {
	return string(alphabetReversed[source.col]) + strconv.Itoa(int(boardRows-source.row)) + string(alphabetReversed[destination.col]) + strconv.Itoa(int(boardRows-destination.row))
}

func (location *Location) IsOnBoard() bool {
	return location.row < boardRows && location.row >= 0 && location.col < boardCols && location.col >= 0
}

func (location *Location) Equals(otherLocation *Location) bool {
	return location.row == otherLocation.row && location.col == otherLocation.col
}

func (location Location) Append(colsToAppendBy int8, rowsToAppendBy int8) Location {
	location.row = location.row + rowsToAppendBy
	location.col = location.col + colsToAppendBy
	return location
}
