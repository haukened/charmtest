package style

import (
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	ActiveBorderColor   lipgloss.Color
	InactiveBorderColor lipgloss.Color
	App                 lipgloss.Style
	Header              lipgloss.Style
	Footer              lipgloss.Style
	HelpKey             lipgloss.Style
	HelpValue           lipgloss.Style
	HelpDivider         lipgloss.Style
	Menu                lipgloss.Style
	MenuCursor          lipgloss.Style
	MenuItem            lipgloss.Style
	SelectedMenuItem    lipgloss.Style
	Pager               lipgloss.Style
	PagerHeader         lipgloss.Style
	PagerFooter         lipgloss.Style
}

func DefaultStyles() *Styles {
	s := new(Styles)

	s.ActiveBorderColor = lipgloss.Color("62")
	s.InactiveBorderColor = lipgloss.Color("236")

	s.App = lipgloss.NewStyle().Margin(1, 2)

	s.Header = lipgloss.NewStyle().
		Foreground(lipgloss.Color("62")).
		Align(lipgloss.Right).
		Bold(true)

	s.Footer = lipgloss.NewStyle().
		MarginTop(1)

	s.HelpKey = lipgloss.NewStyle().
		Foreground(lipgloss.Color("241"))

	s.HelpValue = lipgloss.NewStyle().
		Foreground(lipgloss.Color("239"))

	s.HelpDivider = lipgloss.NewStyle().
		Foreground(lipgloss.Color("237")).
		SetString(" • ")

	s.Menu = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		Padding(1, 2).
		MarginRight(1).
		Width(24)

	s.MenuCursor = lipgloss.NewStyle().
		Foreground(lipgloss.Color("213")).
		SetString(">")

	s.MenuItem = lipgloss.NewStyle().
		PaddingLeft(2)

	s.SelectedMenuItem = lipgloss.NewStyle().
		Foreground(lipgloss.Color("207")).
		PaddingLeft(1)

	s.Pager = lipgloss.NewStyle().
		Align(lipgloss.Left).
		BorderStyle(lipgloss.NormalBorder())

	fhbs := lipgloss.NormalBorder()
	fhbs.Right = "├"

	s.PagerHeader = lipgloss.NewStyle().
		BorderStyle(fhbs).
		Padding(0, 1)

	ffbs := lipgloss.NormalBorder()
	ffbs.Left = "┤"

	s.PagerFooter = lipgloss.NewStyle().
		BorderStyle(ffbs)

	return s
}
