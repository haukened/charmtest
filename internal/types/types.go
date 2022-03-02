package types

type MenuEntry struct {
	Name  string
	Value string
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
