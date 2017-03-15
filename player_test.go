package gobot_test

import (
	"testing"
	"180"
)

func TestPlayer_Opponent(t *testing.T) {
	if gobot.HUMAN.Opponent() != gobot.GOBOT {
		t.Error("Opponent should be Gobot")
	}
	if gobot.GOBOT.Opponent() != gobot.HUMAN {
		t.Error("Opponent should be Human")
	}
}
