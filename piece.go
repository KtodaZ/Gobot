package gobot

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

func (piece Piece) getName() string {

	switch piece {
	case EMPTY:
		return "-"
	case BISHOP_GOB:
		return "B"
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

func (piece Piece) OwnedBy(player Player) bool {
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
