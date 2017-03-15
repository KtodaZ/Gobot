package gobotcore_test

import (
	"testing"
	"github.com/ktodaz/gobot/gobotcore"
)

func TestPlayer_Opponent(t *testing.T) {
	if gobotcore.HUMAN.Opponent() != gobotcore.GOBOT {
		t.Error("Opponent should be Gobot")
	}
	if gobotcore.GOBOT.Opponent() != gobotcore.HUMAN {
		t.Error("Opponent should be Human")
	}
}
