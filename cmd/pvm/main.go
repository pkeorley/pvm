package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"pvm/internal/game"
	"pvm/internal/model"
)

func main() {
	g := game.NewGame(game.NewPlayer(game.PLAYER, 1, 1))
	if err := g.MovePlayer(g.GetCenter()); err != nil {
		log.Fatal(err)
	}

	run(model.NewModel(g))
}

func run(model model.Model) {
	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
