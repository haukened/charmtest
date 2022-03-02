package main

import (
	"charmtest/internal/mainview"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	v := mainview.NewMainView(entries)
	p := tea.NewProgram(v, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		log.Println(err)
	}
}
