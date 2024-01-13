package game

import (
	"github.com/google/uuid"
)

type Coin struct {
	Id   string
	X, Y int
}

func NewCoin(posX, posY int) Coin {
	return Coin{
		Id: uuid.New().String(),
		X:  posX,
		Y:  posY,
	}
}

func (c Coin) String() string {
	return COIN
}

func (c Coin) Integer() int {
	return 2
}
