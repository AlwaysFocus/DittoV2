package tui

import (
	"DittoV2/case"
	"DittoV2/casedetail"
	"DittoV2/tui/caseui"
	"DittoV2/tui/detailui"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"os"
)

var p *tea.Program

type sessionState int

const (
	mainMenuView sessionState = iota
	caseView
)

// MainModel the main model of the program; holds other models and bubbles
type MainModel struct {
	state          sessionState
	helpdesk       tea.Model
	helpdeskdetail tea.Model
	hd             *_case.HelpDesk
	hdd            *casedetail.HelpDeskDetail
	activeCaseNum  uint
	windowSize     tea.WindowSizeMsg
}

// StartTea the entry point for the UI. Initializes the model.
func StartTea(hd _case.HelpDesk, hdd casedetail.HelpDeskDetail) {
	if f, err := tea.LogToFile("debug.log", "help"); err != nil {
		fmt.Println("Couldn't open a file for logging:", err)
		os.Exit(1)
	} else {
		defer func() {
			err = f.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

	m := New(&hd, &hdd)

	p = tea.NewProgram(m, tea.WithAltScreen())

	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

// New initialize the main model for the program
func New(hd *_case.HelpDesk, hdd *casedetail.HelpDeskDetail) MainModel {
	return MainModel{
		state:    mainMenuView,
		helpdesk: caseui.New(hd),
		hdd:      hdd,
	}
}

// Init run any initial IO on program start
func (m MainModel) Init() tea.Cmd {
	return nil
}

// Update handle IO and commands
func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.windowSize = msg // pass this along to the entry view, so it uses the full window size when it's initialized
	case detailui.BackMsg:
		m.state = mainMenuView
	case caseui.SelectMsg:
		// Check the filter state to determine if we need to override the enter key
		if m.helpdesk.(caseui.Model).List.FilterState() == 1 || !m.helpdesk.(caseui.Model).List.FilteringEnabled() {
			return m, nil
		}
		if m.state == mainMenuView {
			m.activeCaseNum = msg.ActiveCaseNum
			m.helpdeskdetail = detailui.New(m.hdd, msg.ActiveCaseNum, p, m.windowSize)
			m.state = caseView
		}
	}

	switch m.state {
	case mainMenuView:
		newHelpDesk, newCmd := m.helpdesk.Update(msg)
		helpDeskModel, ok := newHelpDesk.(caseui.Model)
		if !ok {
			panic("could not perform assertion on caseui model")
		}
		m.helpdesk = helpDeskModel
		cmd = newCmd
	case caseView:
		newCaseDetail, newCmd := m.helpdeskdetail.Update(msg)
		caseDetailModel, ok := newCaseDetail.(detailui.Model)
		if !ok {
			panic("could not perform assertion on detailui model")
		}
		m.helpdeskdetail = caseDetailModel
		cmd = newCmd
	}
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

// View return the text UI to be output to the terminal
func (m MainModel) View() string {
	switch m.state {
	case caseView:
		return m.helpdeskdetail.View()
	default:
		return m.helpdesk.View()
	}
}
