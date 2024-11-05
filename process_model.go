package main

import (
	"strconv"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ProcesstModel struct {
	table table.Model
	help  help.Model
	keys  keyMap
}

type updateProcessesMsg struct {
	process []Process
}

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

func NewProcessModel() *ProcesstModel {
	pm := ProcesstModel{
		keys: keys,
		help: help.New(),
	}

	processes := GetProcesses()

	columns := []table.Column{
		{Title: "S/N", Width: 8},
		{Title: "Port", Width: 60},
		{Title: "PID", Width: 10},
		{Title: "Application", Width: 40},
	}

	rows := []table.Row{}
	for i, p := range processes {
		n := strconv.Itoa(i + 1)
		rows = append(rows, []string{n, p.Port, p.PID, p.Application})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(20),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	pm.table = t

	return &pm
}

func (pm ProcesstModel) Init() tea.Cmd {
	return nil
}

func (pm ProcesstModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case updateProcessesMsg:
		r := []table.Row{}
		for i, p := range msg.process {
			n := strconv.Itoa(i + 1)
			r = append(r, []string{n, p.Port, p.PID, p.Application})
		}
		pm.table.SetRows(r)

		return pm, nil

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, pm.keys.Quit):
			return pm, tea.Quit

		case key.Matches(msg, pm.keys.Escape):
			if pm.table.Focused() {
				pm.table.Blur()
			} else {
				pm.table.Focus()
			}

		case key.Matches(msg, pm.keys.Kill):
			r := pm.table.SelectedRow()
			p := Process{
				Port:        r[1],
				PID:         r[2],
				Application: r[3],
			}
			p.kill()

			return pm, tea.Batch(refetchProcesses)

		case key.Matches(msg, pm.keys.Enter):
			return pm, tea.Batch(
				tea.Printf("Let's go to %s!", pm.table.SelectedRow()[1]),
			)
		}
	}

	pm.table, cmd = pm.table.Update(msg)
	return pm, cmd
}

func refetchProcesses() tea.Msg {
	newP := GetProcesses()

	return updateProcessesMsg{
		process: newP,
	}
}

func (pm ProcesstModel) View() string {
	helpView := pm.help.View(pm.keys)
	return baseStyle.Render(pm.table.View()) + "\n" + helpView + "\n"
}

type keyMap struct {
	Up     key.Binding
	Down   key.Binding
	Left   key.Binding
	Right  key.Binding
	Help   key.Binding
	Quit   key.Binding
	Kill   key.Binding
	Escape key.Binding
	Enter  key.Binding
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Left: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/h", "move left"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "move right"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	Kill: key.NewBinding(
		key.WithKeys("k"),
		key.WithHelp("k", "kill process"),
	),
	Escape: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "toggle focus"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select process"),
	),
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Kill, k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Kill}, // first column
		{k.Help, k.Quit},       // second column
	}
}
