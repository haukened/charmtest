package mainview

import (
	"charmtest/internal/menu"
	"charmtest/internal/pager"
	style "charmtest/internal/styles"
	"charmtest/internal/types"
	"fmt"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MainView struct {
	styles         *style.Styles
	width          int
	height         int
	menuSelections []*types.MenuEntry
	boxes          []tea.Model
	activeBox      int
}

func NewMainView(selections []*types.MenuEntry) *MainView {
	logFile, err := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
	s := style.DefaultStyles()
	menu := menu.NewBubble(selections, s, logFile)
	pager := pager.NewBubble(s, selections[0], logFile)
	return &MainView{
		styles:         s,
		menuSelections: selections,
		boxes:          []tea.Model{menu, pager},
	}
}

func (m *MainView) Init() tea.Cmd {
	return nil
}

func (m *MainView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab", "shift+tab":
			m.activeBox = (m.activeBox + 1) % 2
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		log.Printf("MAIN: Height: %d, Width: %d", msg.Height, msg.Width)
		m.width = msg.Width
		m.height = msg.Height
		maxBoxHeight := msg.Height - lipgloss.Height(m.footerView())
		maxBoxWidth := msg.Width -
			m.styles.Menu.GetWidth() -
			m.styles.Menu.GetHorizontalFrameSize() -
			m.styles.App.GetMarginRight()
		boxMsg := tea.WindowSizeMsg{
			Width:  maxBoxWidth,
			Height: maxBoxHeight,
		}
		for i, bx := range m.boxes {
			bm, cmd := bx.Update(boxMsg)
			m.boxes[i] = bm
			if cmd != nil {
				cmds = append(cmds, cmd)
			}
		}
	case *types.MenuEntry:
		bm, cmd := m.boxes[1].Update(msg)
		m.boxes[1] = bm
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
	}
	ab, cmd := m.boxes[m.activeBox].Update(msg)
	m.boxes[m.activeBox] = ab
	if cmd != nil {
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}

func (m *MainView) View() string {
	s := strings.Builder{}
	lb := m.viewForBox(0)
	rb := m.viewForBox(1)
	s.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, lb, rb))
	s.WriteRune('\n')
	s.WriteString(m.footerView())
	return m.styles.App.Render(s.String())
}

func (m *MainView) viewForBox(i int) string {
	isActive := i == m.activeBox
	switch box := m.boxes[i].(type) {
	case *menu.Bubble:
		s := m.styles.Menu
		if isActive {
			s = s.Copy().BorderForeground(m.styles.ActiveBorderColor)
		}
		return s.Render(box.View())
	case *pager.Bubble:
		s := m.styles.Pager
		if isActive {
			s = s.Copy().BorderForeground(m.styles.ActiveBorderColor)
		}
		return s.Render(box.View())
	default:
		panic(fmt.Sprintf("unknown box type %T", box))
	}
}

func (m *MainView) footerView() string {
	w := &strings.Builder{}
	var h []types.HelpMessage = []types.HelpMessage{
		{Key: "tab", Value: "Change Active Pane"},
	}
	if box, ok := m.boxes[m.activeBox].(types.BubbleHelper); ok {
		h = append(h, box.Help()...)
	}
	for i, v := range h {
		fmt.Fprint(w, helpEntryRender(v, m.styles))
		if i != len(h)-1 {
			fmt.Fprint(w, m.styles.HelpDivider)
		}
	}
	help := w.String()
	return m.styles.Footer.Render(help)
}

func helpEntryRender(h types.HelpMessage, s *style.Styles) string {
	return fmt.Sprintf("%s %s", s.HelpKey.Render(h.Key), s.HelpValue.Render(h.Value))
}
