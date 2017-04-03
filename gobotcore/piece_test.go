package gobotcore_test

import (
	"github.com/ktodaz/gobot/gobotcore"
	"testing"
)

func TestPiece_IsEmpty(t *testing.T) {
	king := gobotcore.KING_GOB
	empty := gobotcore.EMPTY
	if king.IsEmpty() {
		t.Error("King Gob is not empty")
	}
	if !empty.IsEmpty() {
		t.Error("Empty is empty")
	}
}

func TestPiece_IsKing(t *testing.T) {
	king := gobotcore.KING_GOB
	rook := gobotcore.ROOK_GOB
	empty := gobotcore.EMPTY
	if !king.IsKing() {
		t.Error("King should be a King")
	}
	if rook.IsKing() {
		t.Error("Rook is not a King")
	}
	if empty.IsKing() {
		t.Error("Empty is not a King")
	}

}

func TestGetPieceByName(t *testing.T) {
	if gobotcore.GetPieceByName("N") != gobotcore.KNIGHT_GOB {
		t.Error("Knight should be returned")
	}
}
