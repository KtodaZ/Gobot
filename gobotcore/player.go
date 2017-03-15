package gobotcore

type Player int

const (
	GOBOT = Player(iota)
	HUMAN
)

func (player Player) Opponent() Player {
	switch player {
	case GOBOT:
		return HUMAN
	default:
		return GOBOT
	}
}
