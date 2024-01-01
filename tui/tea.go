package tui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	version string
}

func newModel(version string) model {
	return model{version}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	return fmt.Sprintf("Welcome to the TUI for Lobe Version : %s", m.version)
}

func NewTUI() *tea.Program {
	m := newModel("0.0.5")
	p := tea.NewProgram(m)
	return p
}
