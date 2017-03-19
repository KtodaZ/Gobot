package gobotcore

type Move struct {
	from Location
	to   Location
}

func NewMove(from Location, to Location) Move {
	return Move{from: from, to: to}
}

func (move Move) ToString() string {
	return ToStringMultipleLocations(move.from, move.to)
}

func (move Move) IsContainedIn(moves []Move) bool {
	for _, curMove := range moves {
		if move.Equals(curMove) {
			return true
		}
	}
	return false
}

func (move1 Move) Equals(move2 Move) bool {
	return move1.from.Equals(move2.from) && move1.to.Equals(move2.to)
}
