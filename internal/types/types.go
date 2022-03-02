package types

import "github.com/charmbracelet/glamour"

type MenuEntry struct {
	Name  string
	Value []byte
}

func (m *MenuEntry) RenderBytes() (string, error) {
	b, err := glamour.RenderBytes(m.Value, "dark")
	if err != nil {
		return string(m.Value), err
	}
	return string(b), nil
}

type SelectedMessage struct {
	Name  string
	Index int
}

type SetPagerContentMessage struct {
	Title   string
	Content string
}

type BubbleHelper interface {
	Help() []HelpMessage
}

type HelpMessage struct {
	Key   string
	Value string
}

type LogMessage struct {
	Format string
	A      []interface{}
}
