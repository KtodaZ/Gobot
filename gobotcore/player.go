package gobotcore

type Player int8

const (
	GOBOT = Player(iota)
	HUMAN
)

func (player *Player) Opponent() *Player {
	newPlayer := Player(GOBOT)
	if *player == GOBOT {
		newPlayer = HUMAN
	}
	return &newPlayer
}
