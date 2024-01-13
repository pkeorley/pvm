package style

import (
	"github.com/charmbracelet/lipgloss"
	"math/rand"
)

type Styles struct {
}

func NewStyles() *Styles {
	return &Styles{}
}

func (s Styles) GetErrStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#F04848"))
}

func (s Styles) GetPlayerStyle() lipgloss.Style {
	colors := []string{"#F0F391", "#DDE165", "#D2D834", "#A7AC20", "#878B1C"}
	return lipgloss.NewStyle().Foreground(lipgloss.Color(s.GetRandomColorOf(colors)))
}

func (s Styles) GetCoinStyle() lipgloss.Style {
	colors := []string{"#FF9C00", "#DC8700", "#AD7318", "#F0B14F", "#9F7027"}
	return lipgloss.NewStyle().Foreground(lipgloss.Color(s.GetRandomColorOf(colors)))
}

func (s Styles) GetBlankStyle() lipgloss.Style {
	colors := []string{"#4DCB30", "#31A925", "#15820A", "#1DD20A", "#70CD66"}
	return lipgloss.NewStyle().Foreground(lipgloss.Color(s.GetRandomColorOf(colors)))
}

func (s Styles) GetRandomColorOf(colors []string) string {
	return colors[rand.Intn(len(colors))]
}
