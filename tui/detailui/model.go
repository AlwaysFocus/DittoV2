package detailui

import (
	"DittoV2/casedetail"
	"DittoV2/tui/attachmentui"
	"DittoV2/tui/commentui"
	"DittoV2/tui/constants"
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"log"
	"strings"
)

// You generally won't need this unless you're processing stuff with
// complicated ANSI escape sequences. Turn it on if you notice flickering.
//
// Also keep in mind that high performance rendering only works for programs
// that use the full size of the terminal. We're enabling that below with
// tea.EnterAltScreen().
const useHighPerformanceRenderer = false

var (
	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return titleStyle.Copy().BorderStyle(b)
	}()
)

type sessionState int

// The different views that the detail screen can have
const (
	descriptionView sessionState = iota
	commentsView
	contactsView
	attachmentsView
)

// BackMsg change state back to main menu view
type BackMsg bool

// ChangeViewMsg struct to handle view change messages
type ChangeViewMsg struct {
	view string
}

// Model implements tea.Model
type Model struct {
	viewport      viewport.Model
	hdd           *casedetail.HelpDeskDetail
	activeCaseNum uint
	p             *tea.Program
	error         string
	windowSize    tea.WindowSizeMsg
	paginator     paginator.Model
	detail        casedetail.CaseDetail
	ready         bool
	view          string
	attachments   tea.Model
	comments      tea.Model
	contactList   list.Model
	state         sessionState
}

// Init run any initial IO on program start
func (m Model) Init() tea.Cmd {
	return nil
}

// New initialize the casedetailui model for the program
func New(hdd *casedetail.HelpDeskDetail, activeCaseNum uint, p *tea.Program, windowSize tea.WindowSizeMsg) tea.Model {
	m := Model{hdd: hdd, activeCaseNum: activeCaseNum, windowSize: windowSize}
	m.p = p
	m.viewport = viewport.New(windowSize.Width, calculateHeight(windowSize.Height))
	m.viewport.Style = lipgloss.NewStyle().
		Align(lipgloss.Bottom)

	err := m.hdd.GetCaseDetails(m.activeCaseNum)
	m.comments = commentui.New(m.hdd)

	if err != nil {
		return nil
	}

	m.state = descriptionView
	m.view = "description"
	m.setViewportContent(m.hdd.Details.ReturnObj.HDCase[0].Description)

	return m
}

func (m *Model) setViewportContent(contentToRender string) {
	str, err := glamour.Render(contentToRender, "dracula")
	if err != nil {
		m.error = "could not render content with glamour"
	}
	m.viewport.SetContent(str)
}

// Update handle IO and commands
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		verticalMarginHeight := headerHeight + footerHeight

		if !m.ready {
			// Since this program is using the full size of the viewport we
			// need to wait until we've received the window dimensions before
			// we can initialize the viewport. The initial dimensions come in
			// quickly, though asynchronously, which is why we wait for them
			// here.
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			m.viewport.HighPerformanceRendering = useHighPerformanceRenderer

			m.ready = true

			// This is only necessary for high performance rendering, which in
			// most cases you won't need.
			//
			// Render the viewport one line below the header.
			m.viewport.YPosition = headerHeight + 1

			m.viewport, cmd = m.viewport.Update(msg)
			cmds = append(cmds, cmd)
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMarginHeight

		}

	case errMsg:
		m.error = msg.Error()

	case attachmentui.BackMsg:
		m.state = descriptionView
	case commentui.BackMsg:
		m.state = descriptionView
	case tea.KeyMsg:
		// Check to see if the comment input model is focused
		if m.comments.(commentui.Model).Comment.Focused() {
			newComment, newCmd := m.comments.Update(msg)
			commentModel, ok := newComment.(commentui.Model)
			if !ok {
				log.Fatal("comment model is not a commentui.Model")
			}
			m.comments = commentModel
			cmds = append(cmds, newCmd)
			return m, tea.Batch(cmds...)
		}
		switch {
		case key.Matches(msg, constants.Keymap.Attachments):
			m.attachments = attachmentui.New(m.hdd.Details.ReturnObj.HDCaseAttch)
			m.state = attachmentsView
			//cmd = changeViewCmd("attachments")
		case key.Matches(msg, constants.Keymap.Contacts):
			m.view = "contacts"
			//cmd = changeViewCmd("contacts")
		case key.Matches(msg, constants.Keymap.Description):
			m.view = "description"
			//cmd = changeViewCmd("description")
		case key.Matches(msg, constants.Keymap.Comments):
			m.comments = commentui.New(m.hdd)
			m.state = commentsView
			//cmd = changeViewCmd("comments")

		case key.Matches(msg, constants.Keymap.Back):
			if m.state == attachmentsView || m.state == commentsView {
				m.state = descriptionView
				return m, nil
			} else {
				return m, func() tea.Msg {
					return BackMsg(true)
				}
			}

		case msg.String() == "ctrl+c":
			return m, tea.Quit
		case msg.String() == "q":
			return m, tea.Quit
		default:
			m.viewport, cmd = m.viewport.Update(msg)
			cmds = append(cmds, cmd)
		}
	}

	// Depending on which state our detail ui is in
	// We will route the updates to the right model
	switch m.state {
	case attachmentsView:
		newAttachList, newCmd := m.attachments.Update(msg)
		attachmentModel, ok := newAttachList.(attachmentui.Model)
		if !ok {
			log.Fatal("attachment model is not an attachmentui.Model")
		}
		m.attachments = attachmentModel
		cmds = append(cmds, newCmd)
	case commentsView:
		newComment, newCmd := m.comments.Update(msg)
		commentModel, ok := newComment.(commentui.Model)
		if !ok {
			log.Fatal("comment model is not a commentui.Model")
		}
		m.comments = commentModel
		cmds = append(cmds, newCmd)
	}

	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) headerView() string {
	title := titleStyle.Render("Header")
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m Model) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (m Model) helpView() string {
	return constants.HelpStyle("\n ↑/↓: navigate  • esc: back • c: contacts • C: add comment• d: description • a: attachments • q: quit\n")
}

func (m Model) errorView() string {
	return constants.ErrStyle(m.error)
}

// View return the text UI to be output to the terminal
func (m Model) View() string {
	switch m.state {
	case descriptionView:
		content := strings.Join(strings.Fields(m.hdd.Details.ReturnObj.HDCase[0].Description), " ")
		m.setViewportContent(content)
	case attachmentsView:
		return m.attachments.View()

	case commentsView:
		return m.comments.View()
	default:
		content := strings.Join(strings.Fields(m.hdd.Details.ReturnObj.HDCase[0].Description), " ")
		m.setViewportContent(content)

	}

	var content string

	switch m.view {
	case "description":
		// Set the content to the description of the case
		content = strings.Join(strings.Fields(m.hdd.Details.ReturnObj.HDCase[0].Description), " ")
		m.setViewportContent(content)
	case "contacts":
		// Set the content to the contacts of the case
		content = casedetail.FormatCaseContacts(m.hdd.CaseOwnerContact, m.hdd.PrimaryContact, m.hdd.InternalContact)
		m.setViewportContent(content)
	}

	formatted := lipgloss.JoinVertical(lipgloss.Left, "\n", m.viewport.View(), m.helpView(), m.errorView())
	return constants.DocStyle.Render(formatted)
}

/* helpers */

func calculateHeight(height int) int {
	return height - height/7
}
