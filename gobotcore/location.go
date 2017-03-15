package gobotcore

type Location struct {
	row int
	col int
}

func newLocation(row, col int) Location {
	return Location{row: row, col: col}
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