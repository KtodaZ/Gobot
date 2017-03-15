package gobot_test

import (
	"testing"
	"180"
)

func TestPiece_IsKing(t *testing.T) {
	if !gobot.KING_GOB.IsKing() {
		t.Error("King should be a King")
	}
	if gobot.ROOK_GOB.IsKing() {
		t.Error("Rook is not a King")
	}
	if gobot.EMPTY.IsKing() {
		t.Error("Empty is not a King")
	}

}

func TestPiece_Morph(t *testing.T) {
	if gobot.BISHOP_GOB.Morph() != gobot.KNIGHT_GOB {
		t.Error("Bishop should become knight")
	}
	if gobot.BISHOP_HUM.Morph() != gobot.KNIGHT_HUM {
		t.Error("Bishop should become knight")
	}
	if gobot.ROOK_GOB.Morph() != gobot.BISHOP_GOB {
		t.Error("Rook should become Bishop")
	}
	if gobot.ROOK_HUM.Morph() != gobot.BISHOP_HUM {
		t.Error("Rook should become Bishop")
	}
	if gobot.KNIGHT_GOB.Morph() != gobot.ROOK_GOB {
		t.Error("Knight should become Rook")
	}
	if gobot.KNIGHT_HUM.Morph() != gobot.ROOK_HUM {
		t.Error("Knight should become Rook")
	}
	if gobot.PAWN_GOB.Morph() != gobot.PAWN_GOB {
		t.Error("Pawn should stay same")
	}
	if gobot.PAWN_HUM.Morph() != gobot.PAWN_HUM {
		t.Error("Pawn should stay same")
	}
	if gobot.KING_GOB.Morph() != gobot.KING_GOB {
		t.Error("King should stay same")
	}
	if gobot.KING_HUM.Morph() != gobot.KING_HUM {
		t.Error("King should stay same")
	}
}
