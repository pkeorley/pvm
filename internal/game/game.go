package game

import (
	"fmt"
	"pvm/internal/style"
)

var (
	cantMovePlayer = fmt.Errorf("you can't move player to this coordinates")
	cantInsertCoin = fmt.Errorf("you can't insert coin to this place")
)

type Game struct {
	Board  [][]int
	Player *Player
	Coins  []Coin
}

func NewGame(player *Player) *Game {

	board := make([][]int, WIDTH)

	for i := 0; i < WIDTH; i++ {
		for j := 0; j < HEIGHT; j++ {
			board[i] = append(board[i], 0)
		}
	}

	return &Game{
		Board:  board,
		Player: player,
	}
}

func (g *Game) MovePlayer(posX, posY int) error {

	if posX <= 0 && posX > WIDTH {
		return cantMovePlayer
	} else if posY <= 0 && posY > HEIGHT {
		return cantMovePlayer
	}

	g.updatePlayerMoving(0)
	g.Player.PositionX = posX
	g.Player.PositionY = posY
	g.updatePlayerMoving(1)

	return nil
}

func (g *Game) MovePlayerToTop() error {

	if g.Player.PositionY <= 0 {
		return cantMovePlayer
	}

	g.updatePlayerMoving(0)
	g.Player.PositionY -= 1
	g.updatePlayerMoving(1)

	return nil
}

func (g *Game) MovePlayerToRight() error {

	if g.Player.PositionX >= WIDTH-1 {
		return cantMovePlayer
	}

	g.updatePlayerMoving(0)
	g.Player.PositionX += 1
	g.updatePlayerMoving(1)

	return nil
}

func (g *Game) MovePlayerToBottom() error {

	if g.Player.PositionY >= HEIGHT-1 {
		return cantMovePlayer
	}

	g.updatePlayerMoving(0)
	g.Player.PositionY += 1
	g.updatePlayerMoving(1)

	return nil
}

func (g *Game) MovePlayerToLeft() error {

	if g.Player.PositionX <= 0 {
		return cantMovePlayer
	}

	g.updatePlayerMoving(0)
	g.Player.PositionX -= 1
	g.updatePlayerMoving(1)

	return nil
}

func (g *Game) updatePlayerMoving(number int) {
	if g.Board[g.Player.PositionY][g.Player.PositionX] == 2 {
		g.RemoveCoinByCoordinates(g.Player.PositionX, g.Player.PositionY)
		g.Player.CoinBalance += 1
	}

	g.Board[g.Player.PositionY][g.Player.PositionX] = number
}

func (g *Game) String() string {
	out := ""

	for _, i := range g.Board {
		for _, j := range i {
			out += fmt.Sprintf(" %s ", g.convertNumberToCharacter(j))
		}
		out += "\n"
	}

	return out
}

func (g *Game) convertNumberToCharacter(number int) (character string) {

	s := style.NewStyles()

	switch number {
	case 0:
		character = s.GetBlankStyle().Render(BLANK)
	case g.Player.Integer():
		character = s.GetPlayerStyle().Render(g.Player.Character)
	case 2:
		character = s.GetCoinStyle().Render(COIN)
	}
	return
}

func (g *Game) GetCenter() (int, int) {
	return HEIGHT / 2, WIDTH / 2
}

func (g *Game) AddCoin(coin Coin) error {
	if coin.X <= 0 && coin.X > WIDTH {
		return cantMovePlayer
	} else if coin.Y <= 0 && coin.Y > HEIGHT {
		return cantInsertCoin
	}

	g.Coins = append(g.Coins, coin)
	g.Board[coin.Y][coin.X] = coin.Integer()

	return nil
}

func (g *Game) RemoveCoinByCoinId(coinId string) error {
	var findCoin *Coin
	var index int

	for i, coin := range g.Coins {
		if coin.Id == coinId {
			findCoin = &g.Coins[i]
			index = i
			break
		}
	}

	if findCoin == nil {
		return fmt.Errorf("coin with ID %s not found", coinId)
	}

	g.Board[findCoin.Y][findCoin.X] = 0
	g.Coins = append(g.Coins[:index], g.Coins[index+1:]...)

	return nil
}

func (g *Game) RemoveCoinByCoordinates(x, y int) {

	for index, coin := range g.Coins {
		if coin.X == x && coin.Y == y {
			g.Coins = append(g.Coins[:index], g.Coins[index+1:]...)
		}
	}

}
