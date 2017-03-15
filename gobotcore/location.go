package gobotcore

type location struct {
	row int
	col int
}

func newLocation(row, col int) location {
	return location{row: row, col: col}
}

func (location location) getRow() int {
	return location.row
}

func (location location) getCol() int {
	return location.col
}

func (location location) isOnBoard() bool {
	return location.row < boardRows && location.row >= 0 && location.col < boardCols && location.col >= 0
}

func(location location) equals (otherLocation location) bool {
	return location.row == otherLocation.row && location.col == otherLocation.col
}