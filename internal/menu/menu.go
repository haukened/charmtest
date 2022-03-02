package menu

import (
	style "charmtest/internal/styles"
	"charmtest/internal/types"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/truncate"
)

type Bubble struct {
	Items  []*types.MenuEntry
	Cursor int
	styles *style.Styles
}

func NewBubble(items []*types.MenuEntry, styles *style.Styles, logFile *os.File) *Bubble {
	log.SetOutput(logFile)
	return &Bubble{
		Items:  items,
		styles: styles,
	}
}

func (b *Bubble) Init() tea.Cmd {
	return nil
}

func (b *Bubble) View() string {
	s := strings.Builder{}
	menuNameMaxWidth := b.styles.Menu.GetWidth() - // menu width
		b.styles.Menu.GetHorizontalPadding() - // menu padding
		lipgloss.Width(b.styles.MenuCursor.String()) - // cursor
		b.styles.MenuItem.GetHorizontalFrameSize() // menu item gaps
	for i, item := range b.Items {
		item := truncate.StringWithTail(item.Name, uint(menuNameMaxWidth), "…")
		if i == b.Cursor {
			s.WriteString(b.styles.MenuCursor.String())
			s.WriteString(b.styles.SelectedMenuItem.Render(item))
		} else {
			s.WriteString(b.styles.MenuItem.Render(item))
		}
		if i < len(b.Items)-1 {
			s.WriteRune('\n')
		}
	}
	return s.String()
}

func (b *Bubble) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "k", "up":
			if b.Cursor > 0 {
				b.Cursor--
				cmds = append(cmds, b.sendSelectedMessage)
			}
		case "j", "down":
			if b.Cursor < len(b.Items)-1 {
				b.Cursor++
				cmds = append(cmds, b.sendSelectedMessage)
			}
		}
	}
	return b, tea.Batch(cmds...)
}

func (b *Bubble) Help() []types.HelpMessage {
	return []types.HelpMessage{
		{Key: "↑/↓", Value: "Navigate"},
	}
}

func (b *Bubble) sendSelectedMessage() tea.Msg {
	if b.Cursor >= 0 && b.Cursor < len(b.Items) {
		return b.Items[b.Cursor]
	}
	return nil
}
