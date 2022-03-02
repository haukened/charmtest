package pager

import (
	style "charmtest/internal/styles"
	"charmtest/internal/types"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

const useHighPerformanceRenderer = false

type Bubble struct {
	title      string
	initial    *types.MenuEntry
	ready      bool
	styles     *style.Styles
	viewport   viewport.Model
	lastHeight int
	lastWidth  int
	err        error
}

func NewBubble(s *style.Styles, initial *types.MenuEntry, logFile *os.File) *Bubble {
	log.SetOutput(logFile)
	return &Bubble{
		styles:  s,
		initial: initial,
	}
}

func (m *Bubble) Init() tea.Cmd {
	return nil
}

func (m *Bubble) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		log.Printf("PAGER Height: %d, Width: %d", msg.Height, msg.Width)
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		verticalMarginHeight := headerHeight + footerHeight
		allowedHeight := msg.Height - verticalMarginHeight - (2 * m.styles.App.GetVerticalFrameSize()) - (2 * m.styles.Pager.GetVerticalFrameSize())
		if !m.ready {
			var content string
			content, m.err = m.initial.RenderBytes()
			if m.err != nil {
				cmds = append(cmds, m.sendErrorMessage)
				return m, tea.Batch(cmds...)
			}
			m.viewport = viewport.New(msg.Width, allowedHeight)
			m.viewport.YPosition = headerHeight
			m.viewport.HighPerformanceRendering = useHighPerformanceRenderer
			m.viewport.SetContent(content)
			m.title = m.initial.Name
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = allowedHeight
		}
		m.lastWidth = msg.Width
		m.lastHeight = verticalMarginHeight
		if useHighPerformanceRenderer {
			cmds = append(cmds, viewport.Sync(m.viewport))
		}
	case *types.MenuEntry:
		m.title = msg.Name
		var content string
		content, m.err = msg.RenderBytes()
		if m.err != nil {
			cmds = append(cmds, m.sendErrorMessage)
			return m, tea.Batch(cmds...)
		}
		m.viewport.SetContent(wordwrap.String(content, m.lastWidth-2))
		m.viewport.GotoTop()
	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m *Bubble) View() string {
	if !m.ready {
		return "\n Initializing..."
	}
	return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.viewport.View(), m.footerView())
}

func (m *Bubble) headerView() string {
	title := m.styles.PagerHeader.Render(m.title)
	line := strings.Repeat("-", max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m *Bubble) footerView() string {
	info := m.styles.PagerFooter.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := strings.Repeat("-", max(0, m.viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func (m *Bubble) sendErrorMessage() tea.Msg {
	var errs []interface{}
	errs = append(errs, m.err)
	return types.LogMessage{
		Format: "%s",
		A:      errs,
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
