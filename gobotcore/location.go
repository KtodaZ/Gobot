package gobotcore

import (
	"strings"
	"strconv"
)

type Location struct {
	row int
	col int
}

func NewLocation(row, col int) Location {
	return Location{row: row, col: col}
}

func NewLocationFromString(readable string) Location {
	// readable should be in the form "A2" where A corresponds to a letter row, and 2 corresponds to a column
	if len(readable) > 2 {
		panic("Length of readable must be 2.")
	}

	var alphabet string= "ABCDEF"
	var rowLetter string = readable[:1]
	var colNumber string = readable[1:]

	indexOfLetterInAlphabet := strings.Index(alphabet, rowLetter)
	if indexOfLetterInAlphabet > -1 {
		return NewLocation(indexOfLetterInAlphabet, AtoiEZPZ(colNumber))
	} else {
		panic("Inputed Row index out of bounds")
	}
}

func AtoiEZPZ(str string) int {
	if i, err := strconv.Atoi(str); err == nil {
		return i
	}
	return -1
}

func (location Location) getRow() int {
	return location.row
}

func (location Location) getCol() int {
	return location.col
}

func (location Location) isOnBoard() bool {
	return location.row < boardRows && location.row >= 0 && location.col < boardCols && location.col >= 0
}

func(location Location) equals (otherLocation Location) bool {
	return location.row == otherLocation.row && location.col == otherLocation.col
}