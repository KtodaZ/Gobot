package gobotcore

type Piece int

const (
	EMPTY = Piece(iota)
	BISHOP_GOB
	ROOK_GOB
	KNIGHT_GOB
	PAWN_GOB
	KING_GOB
	BISHOP_HUM
	ROOK_HUM
	KNIGHT_HUM
	PAWN_HUM
	KING_HUM
)

func (piece Piece) GetName() string {

	switch piece {
	case EMPTY:
		return "-"
	case BISHOP_GOB:
		return "B" //TODO - make these unicode for the competition?
	case BISHOP_HUM:
		return "b"
	case ROOK_GOB:
		return "R"
	case ROOK_HUM:
		return "r"
	case KNIGHT_GOB:
		return "N"
	case KNIGHT_HUM:
		return "n"
	case PAWN_GOB:
		return "P"
	case PAWN_HUM:
		return "p"
	case KING_GOB:
		return "K"
	case KING_HUM:
		return "k"
	}
	panic("Unknown piece")
}

func GetPieceByName(name string) Piece {
	switch name {
	case "-":
		return EMPTY
	case "B":
		return BISHOP_GOB
	case "b":
		return BISHOP_HUM
	case "R":
		return ROOK_GOB
	case "r":
		return ROOK_HUM
	case "N":
		return KNIGHT_GOB
	case "n":
		return KNIGHT_HUM
	case "P":
		return PAWN_GOB
	case "p":
		return PAWN_HUM
	case "K":
		return KING_GOB
	case "k":
		return KING_HUM
	}
	panic("Unknown name")
}

func (piece Piece) IsOwnedBy(player Player) bool {
	switch player {
	case GOBOT:
		switch piece {
		case BISHOP_GOB:
			return true
		case ROOK_GOB:
			return true
		case KNIGHT_GOB:
			return true
		case PAWN_GOB:
			return true
		case KING_GOB:
			return true
		default:
			return false
		}
	default:
		switch piece {
		case BISHOP_HUM:
			return true
		case ROOK_HUM:
			return true
		case KNIGHT_HUM:
			return true
		case PAWN_HUM:
			return true
		case KING_HUM:
			return true
		default:
			return false
		}

	}
}

func (piece Piece) IsEmpty() bool {
	return piece == EMPTY
}

func (piece Piece) IsKing() bool {
	switch piece {
	case KING_GOB:
		return true
	case KING_HUM:
		return true
	default:
		return false
	}

}

func (piece Piece) Morph() Piece {
	switch piece {
	case BISHOP_GOB:
		return KNIGHT_GOB
	case BISHOP_HUM:
		return KNIGHT_HUM
	case ROOK_GOB:
		return BISHOP_GOB
	case ROOK_HUM:
		return BISHOP_HUM
	case KNIGHT_GOB:
		return ROOK_GOB
	case KNIGHT_HUM:
		return ROOK_HUM
	case PAWN_GOB:
		return PAWN_GOB // Pawns do not evolve
	case PAWN_HUM:
		return PAWN_HUM
	case KING_GOB:
		return KING_GOB // Kings do not evolve
	case KING_HUM:
		return KING_HUM
	}
	panic("Unknown piece")
}

func (piece Piece) UnMorph() Piece {
	switch piece {
	case KNIGHT_GOB:
		return BISHOP_GOB
	case KNIGHT_HUM:
		return BISHOP_HUM
	case BISHOP_GOB:
		return ROOK_GOB
	case BISHOP_HUM:
		return ROOK_HUM
	case ROOK_GOB:
		return KNIGHT_GOB
	case ROOK_HUM:
		return KNIGHT_HUM
	case PAWN_GOB:
		return PAWN_GOB // Pawns do not evolve
	case PAWN_HUM:
		return PAWN_HUM
	case KING_GOB:
		return KING_GOB // Kings do not evolve
	case KING_HUM:
		return KING_HUM
	}
	panic("Unknown piece")
}

func (piece Piece) Weight() float64 {
	switch piece {
	case BISHOP_GOB:
		return 6.0
	case BISHOP_HUM:
		return 6.0
	case ROOK_GOB:
		return 6.0
	case ROOK_HUM:
		return 6.0
	case KNIGHT_GOB:
		return 6.0
	case KNIGHT_HUM:
		return 6.0
	case PAWN_GOB:
		return 1.0
	case PAWN_HUM:
		return 1.0
	case KING_GOB:
		return 1000
	case KING_HUM:
		return 1000
	case EMPTY:
		return 0.0
	}
	panic("Unknown piece")
}

// Different weight from above - for sorting moves
func (piece Piece) MoveWeight() int {
	switch piece {
	case BISHOP_GOB:
		return 1
	case BISHOP_HUM:
		return 1
	case ROOK_GOB:
		return 1
	case ROOK_HUM:
		return 1
	case KNIGHT_GOB:
		return 1
	case KNIGHT_HUM:
		return 1
	case PAWN_GOB:
		return 1
	case PAWN_HUM:
		return 1
	case KING_GOB:
		return 0
	case KING_HUM:
		return 0
	case EMPTY:
		return 2
	default:
		return -1
	}
}