package gobotcore

type Move struct {
	from   Location
	to     Location
	weight int8 // Weigh capturing moves higher than others for sorting
}

type ScoredMove struct {
	move  Move
	score float32
}

type Moves []Move

func NewMove(from Location, to Location) Move {
	return Move{from: from, to: to}
}

func NewMoveFromString(str string) Move {
	return NewMove(NewLocationsFromString(str))
}

func (move Move) ToString() string {
	return ToStringMultipleLocations(move.from, move.to)
}

func (move Move) ToStringFlipped() string {
	return ToStringMultipleLocationsFlipped(move.from, move.to)
}

func (move Move) IsContainedIn(moves *Moves) bool {
	for _, curMove := range *moves {
		if move.Equals(&curMove) {
			return true
		}
	}
	return false
}

func (move1 *Move) Equals(move2 *Move) bool {
	return move1.from.Equals(&move2.from) && move1.to.Equals(&move2.to)
}

func (move *Move) GetReverse() *Move {
	from := move.from
	move.from = move.to
	move.to = from
	return move
}

func (move *Move) From() *Location {
	return &move.from
}

func (move *Move) To() *Location {
	return &move.to
}

func (move ScoredMove) Move() *Move {
	return &move.move
}

func (move ScoredMove) Score() *float32 {
	return &move.score
}

// Implementing the sort interface
func (move Moves) Len() int {
	return len(move)
}

func (move Moves) Less(i, j int) bool {
	return move[i].weight < move[j].weight
}

func (move Moves) Swap(i, j int) {
	move[i], move[j] = move[j], move[i]
}
