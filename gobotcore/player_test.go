package gobotcore_test

import (
	"github.com/ktodaz/gobot/gobotcore"
	"testing"
)

func TestPlayer_Opponent(t *testing.T) {
	human := gobotcore.Player(gobotcore.HUMAN)
	gobot := gobotcore.Player(gobotcore.GOBOT)
	humanOpponent := human.Opponent()
	gobotOpponent := gobot.Opponent()

	if *humanOpponent != gobot {
		t.Error("Opponent should be Gobot")
	}
	if *gobotOpponent != human {
		t.Error("Opponent should be Human")
	}
}
