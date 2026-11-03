package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
	spinner  spinner.Model
}

const (
	Reset      = "\033[0m"
	RedFont    = "\033[31m"
	GreenFont  = "\033[32m"
	YellowFont = "\033[33m"
	BlueFont   = "\033[34m"
)

func main() {
	p := tea.NewProgram(initiateModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("There is an error: %v", err)
		os.Exit(1)
	}
}

func initiateModel() model {
	s := spinner.New()
	s.Spinner = spinner.Dot

	return model{
		choices:  []string{"a", "b", "c"},
		selected: make(map[int]struct{}),
		spinner:  s,
	}
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor == 0 {
				m.cursor = len(m.choices) - 1
			} else {
				m.cursor--
			}

		case "down", "j":
			m.cursor++
			m.cursor %= len(m.choices)

		case "enter", " ":
			_, ok := m.selected[m.cursor]

			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}

		}
	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	s := fmt.Sprintf("\n   %s Choose one. \n\n", m.spinner.View())
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.\n"

	return s
}

func Red(s string) string {
	return RedFont + s + Reset
}
