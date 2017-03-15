package gobotcore_test

import (
	"testing"
	"180"
)

func TestPiece_IsKing(t *testing.T) {
	if !main.KING_GOB.IsKing() {
		t.Error("King should be a King")
	}
	if main.ROOK_GOB.IsKing() {
		t.Error("Rook is not a King")
	}
	if main.EMPTY.IsKing() {
		t.Error("Empty is not a King")
	}

}

func TestPiece_Morph(t *testing.T) {
	if main.BISHOP_GOB.Morph() != main.KNIGHT_GOB {
		t.Error("Bishop should become knight")
	}
	if main.BISHOP_HUM.Morph() != main.KNIGHT_HUM {
		t.Error("Bishop should become knight")
	}
	if main.ROOK_GOB.Morph() != main.BISHOP_GOB {
		t.Error("Rook should become Bishop")
	}
	if main.ROOK_HUM.Morph() != main.BISHOP_HUM {
		t.Error("Rook should become Bishop")
	}
	if main.KNIGHT_GOB.Morph() != main.ROOK_GOB {
		t.Error("Knight should become Rook")
	}
	if main.KNIGHT_HUM.Morph() != main.ROOK_HUM {
		t.Error("Knight should become Rook")
	}
	if main.PAWN_GOB.Morph() != main.PAWN_GOB {
		t.Error("Pawn should stay same")
	}
	if main.PAWN_HUM.Morph() != main.PAWN_HUM {
		t.Error("Pawn should stay same")
	}
	if main.KING_GOB.Morph() != main.KING_GOB {
		t.Error("King should stay same")
	}
	if main.KING_HUM.Morph() != main.KING_HUM {
		t.Error("King should stay same")
	}
}
