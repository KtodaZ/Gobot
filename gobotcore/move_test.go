package gobotcore

import "testing"

func TestMove_GetReverse(t *testing.T) {
	move := Move{Location{1, 2}, Location{2, 1}}
	expected := Move{Location{2, 1}, Location{1, 2}}
	move = move.GetReverse()
	if !move.Equals(expected) {
		t.Error("Should be equal")
	}
}
