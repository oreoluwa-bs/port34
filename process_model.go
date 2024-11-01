package main

import (
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type ProcesstModel struct {
	table table.Model
}

type updateProcessesMsg struct {
	process []Process
}

func NewProcessModel() *ProcesstModel {
	pm := ProcesstModel{}

	processes := GetProcesses()

	columns := []table.Column{
		{Title: "S/N", Width: 8},
		{Title: "Port", Width: 40},
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
		pm.table.SetWidth(msg.Width)
		pm.table.SetHeight(msg.Height)

	case updateProcessesMsg:
		r := []table.Row{}
		for i, p := range msg.process {
			n := strconv.Itoa(i + 1)
			r = append(r, []string{n, p.Port, p.PID, p.Application})
		}
		pm.table.SetRows(r)

		return pm, nil

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

		case "k":
			r := pm.table.SelectedRow()
			p := Process{
				Port:        r[1],
				PID:         r[2],
				Application: r[3],
			}
			p.kill()

			return pm, tea.Batch(refetchProcesses)

		case "enter":
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
	return pm.table.View() + "\n  " + pm.table.HelpView() + "\n"
}
