package model

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"math/rand"
	"pvm/internal/game"
	"pvm/internal/style"
)

type Model struct {
	Game *game.Game
	Err  error
}

func NewModel(game *game.Game) Model {
	return Model{
		Game: game,
		Err:  nil,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var err error = nil

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "esc":
			return m, tea.Quit

		case "up", "w":
			err = m.Game.MovePlayerToTop()

		case "right", "d":
			err = m.Game.MovePlayerToRight()

		case "down", "s":
			err = m.Game.MovePlayerToBottom()

		case "left", "a":
			err = m.Game.MovePlayerToLeft()

		case "e":
			err = m.Game.MovePlayerToRight()
			err = m.Game.MovePlayerToTop()

		case "q":
			err = m.Game.MovePlayerToLeft()
			err = m.Game.MovePlayerToTop()

		case "enter":
			x := rand.Intn(game.WIDTH)
			y := rand.Intn(game.HEIGHT)

			err = m.Game.AddCoin(game.NewCoin(x, y))
		}
	}

	m.Err = err

	return m, nil
}

func (m Model) View() string {
	s := style.NewStyles()

	var err = "¯\\_(ツ)_/¯"

	if m.Err != nil {
		err = s.GetErrStyle().Render(m.Err.Error())
	}

	g := m.Game.String()
	coins := s.GetCoinStyle().Render(fmt.Sprintf("%d %s", m.Game.Player.CoinBalance, game.COIN))

	return fmt.Sprintf(
		"%s\n\n* err: %v\n* Coins: %s",
		g,
		err,
		coins,
	)
}
