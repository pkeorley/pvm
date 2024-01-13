package game

type Player struct {
	Character   string
	PositionX   int
	PositionY   int
	CoinBalance int
}

func NewPlayer(character string, positionX, positionY int) *Player {
	return &Player{
		Character: character,
		PositionX: positionX,
		PositionY: positionY,
	}
}

func (p Player) String() string {
	return PLAYER
}

func (p Player) Integer() int {
	return 1
}
