package commentui

import (
	"DittoV2/casedetail"
	"DittoV2/tui/constants"
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

// BackMsg change state back to detail view
type BackMsg bool

// Model the data model for the Comment ui
type Model struct {
	hdd      *casedetail.HelpDeskDetail
	Comment  textinput.Model
	comments []casedetail.Comment
	err      error
}

func New(hdd *casedetail.HelpDeskDetail) Model {
	comment := textinput.New()
	comment.Placeholder = "Enter Comment here"
	comment.CharLimit = 1000
	comment.Width = 100

	return Model{
		Comment: comment,
		hdd:     hdd,
		err:     nil,
	}

}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, constants.Keymap.Comments):
			m.Comment.Focus()
			return m, nil
		case key.Matches(msg, constants.Keymap.Back):
			m.Comment.Blur()
			return m, func() tea.Msg {
				return BackMsg(true)
			}
		case msg.String() == "q":
			return m, tea.Quit

		}
		switch msg.Type {
		case tea.KeyEnter:
			// Add Comment
			err := m.hdd.AddComment(int32(m.hdd.Details.ReturnObj.HDCase[0].HDCaseNum), m.Comment.Value())
			if err != nil {
				m.err = err
				log.Fatal(err)
			}
			// Clear the comment and return to the main menu
			m.Comment.SetValue("")
			m.Comment.Blur()
			return m, func() tea.Msg {
				return BackMsg(true)
			}

		}

	}

	m.Comment, cmd = m.Comment.Update(msg)
	return m, cmd
}

func (m Model) helpView() string {
	return constants.HelpStyle("\n esc: back • enter: add comment to case • q: quit\n")
}

func (m Model) View() string {
	return fmt.Sprintf(
		"Please enter your Comment:\n\n%s\n\n%s",
		m.Comment.View(),
		m.helpView(),
	) + "\n"
}
