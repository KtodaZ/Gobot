package gobotcore_test

import (
	"github.com/ktodaz/gobot/gobotcore"
	"testing"
)

func TestPiece_IsKing(t *testing.T) {
	if !gobotcore.KING_GOB.IsKing() {
		t.Error("King should be a King")
	}
	if gobotcore.ROOK_GOB.IsKing() {
		t.Error("Rook is not a King")
	}
	if gobotcore.EMPTY.IsKing() {
		t.Error("Empty is not a King")
	}

}

func TestPiece_Morph(t *testing.T) {
	if gobotcore.BISHOP_GOB.Morph() != gobotcore.KNIGHT_GOB {
		t.Error("Bishop should become knight")
	}
	if gobotcore.BISHOP_HUM.Morph() != gobotcore.KNIGHT_HUM {
		t.Error("Bishop should become knight")
	}
	if gobotcore.ROOK_GOB.Morph() != gobotcore.BISHOP_GOB {
		t.Error("Rook should become Bishop")
	}
	if gobotcore.ROOK_HUM.Morph() != gobotcore.BISHOP_HUM {
		t.Error("Rook should become Bishop")
	}
	if gobotcore.KNIGHT_GOB.Morph() != gobotcore.ROOK_GOB {
		t.Error("Knight should become Rook")
	}
	if gobotcore.KNIGHT_HUM.Morph() != gobotcore.ROOK_HUM {
		t.Error("Knight should become Rook")
	}
	if gobotcore.PAWN_GOB.Morph() != gobotcore.PAWN_GOB {
		t.Error("Pawn should stay same")
	}
	if gobotcore.PAWN_HUM.Morph() != gobotcore.PAWN_HUM {
		t.Error("Pawn should stay same")
	}
	if gobotcore.KING_GOB.Morph() != gobotcore.KING_GOB {
		t.Error("King should stay same")
	}
	if gobotcore.KING_HUM.Morph() != gobotcore.KING_HUM {
		t.Error("King should stay same")
	}
}
