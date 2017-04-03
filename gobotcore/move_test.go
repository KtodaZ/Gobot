package gobotcore

import "testing"

func TestMove_GetReverse(t *testing.T) {
	move := Move{from: Location{1, 2}, to: Location{2, 1}}
	expected := Move{from: Location{2, 1}, to: Location{1, 2}}
	move = *move.GetReverse()
	if !move.Equals(&expected) {
		t.Error("Should be equal")
	}
}
