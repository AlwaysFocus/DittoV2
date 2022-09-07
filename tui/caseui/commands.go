package caseui

import tea "github.com/charmbracelet/bubbletea"

func selectCaseCmd(ActiveCaseNum uint) tea.Cmd {
	return func() tea.Msg {
		return SelectMsg{ActiveCaseNum: ActiveCaseNum}
	}
}
