package main

import (
	"charmtest/internal/mainview"
	"charmtest/internal/types"
	"io/fs"
	"log"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	var entries []*types.MenuEntry
	for _, file := range find("./entries", ".md") {
		entry, err := mainview.NewMenuEntryFromMarkdown(file)
		if err != nil {
			log.Fatal(err)
		}
		entries = append(entries, entry)
	}
	if len(entries) == 0 {
		log.Fatal("no markdown files found")
	}
	v := mainview.NewMainView(entries)
	p := tea.NewProgram(v, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		log.Println(err)
	}
}

func find(root, ext string) []string {
	var results []string
	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(d.Name()) == ext {
			results = append(results, path)
		}
		return nil
	})
	return results
}
