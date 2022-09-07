package detailui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func changeViewCmd(view string) tea.Cmd {
	return func() tea.Msg {
		return ChangeViewMsg{view}
	}
}
