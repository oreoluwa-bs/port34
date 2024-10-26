package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type ProcesstModel struct {
	table table.Model
}

func NewProcessModel() *ProcesstModel {
	pm := ProcesstModel{}

	columns := []table.Column{
		{Title: "S/N", Width: 4},
		{Title: "Port", Width: 10},
		{Title: "PID", Width: 10},
		{Title: "Application", Width: 10},
	}

	rows := []table.Row{
		{"1", "3000", "93940", "Nodejs runtime"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	pm.table = t

	return &pm
}

func (pm ProcesstModel) Init() tea.Cmd {
	return nil
}

func (pm ProcesstModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		fmt.Println(msg.Height)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return pm, tea.Quit

		case "esc":
			if pm.table.Focused() {
				pm.table.Blur()
			} else {
				pm.table.Focus()
			}

		case "enter":
			return pm, tea.Batch(
				tea.Printf("Let's go to %s!", pm.table.SelectedRow()[1]),
			)
		}
	}

	pm.table, cmd = pm.table.Update(msg)
	return pm, cmd
}

func (pm ProcesstModel) View() string {
	return pm.table.View() + "\n  " + pm.table.HelpView() + "\n"
}
