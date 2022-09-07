package caseui

import (
	_case "DittoV2/case"
	"DittoV2/tui/constants"
	"log"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// SelectMsg the message to change the view to the selected entry
type SelectMsg struct {
	ActiveCaseNum uint
}

type mode int

const (
	nav mode = iota
)

// Model the Case List model Definition
type Model struct {
	mode        mode
	List        list.Model
	hd          *_case.HelpDesk
	isFiltering bool
}

// New initialize the caseui model for the program
func New(hd *_case.HelpDesk) tea.Model {

	items := newCaseList(hd)
	m := Model{mode: nav, List: list.New(items, list.NewDefaultDelegate(), 0, 0), hd: hd}
	m.List.Title = "Cases"
	m.List.SetShowFilter(true)
	m.List.SetFilteringEnabled(true)
	m.List.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			constants.Keymap.Enter,
			constants.Keymap.Back,
		}
	}
	return m
}

func newCaseList(hd *_case.HelpDesk) []list.Item {
	cases, err := hd.GetAllCases()
	if err != nil {
		log.Fatal(err)
	}
	return casesToItems(cases)
}

// Init run any intial IO on program start
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handle IO and commands
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Set the current filter state
	// We use this to override the select message in the main model
	// If the FilterState is "filtering" then set to true
	// If the FilterState is "not filtering" then set to false
	if m.List.FilterState() == 1 {
		m.isFiltering = true
	} else {
		m.isFiltering = false
	}

	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		top, right, bottom, left := constants.DocStyle.GetMargin()
		m.List.SetSize(msg.Width-left-right, msg.Height-top-bottom-1)
	case updateCaseListMsg:
		cases, err := m.hd.GetAllCases()
		if err != nil {
			log.Fatal(err)
		}
		items := casesToItems(cases)
		m.List.SetItems(items)
		m.mode = nav

	case tea.KeyMsg:

		switch {

		case msg.String() == "ctrl+c":
			return m, tea.Quit
		case key.Matches(msg, constants.Keymap.Enter):
			cmd = selectCaseCmd(m.getActiveCaseNum())
		}
		cmds = append(cmds, cmd)
	}

	m.List, cmd = m.List.Update(msg)

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

// View return the text UI to be output to the terminal
func (m Model) View() string {
	return constants.DocStyle.Render(m.List.View() + "\n")
}

// TODO: use generics
// casesToItems convert []model.Case to []list.Item
func casesToItems(cases []_case.Case) []list.Item {
	items := make([]list.Item, len(cases))
	for i, cs := range cases {
		items[i] = list.Item(cs)
	}
	return items
}

// Get the active CaseNum
func (m Model) getActiveCaseNum() uint {

	if m.List.FilterState() == 1 || m.List.FilterState() == 2 {
		var item = m.List.SelectedItem()
		return uint(item.(_case.Case).HDCaseNum)

	} else {
		items := m.List.Items()
		activeItem := items[m.List.Index()]

		return uint(activeItem.(_case.Case).HDCaseNum)
	}
}
