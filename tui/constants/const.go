package constants

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/lipgloss"
)

// DocStyle styling for viewports
var DocStyle = lipgloss.NewStyle().Margin(0, 2)

// HelpStyle styling for help context menu
var HelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render

// ErrStyle provides styling for error messages
var ErrStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#bd534b")).Render

// AlertStyle provides styling for alert messages
var AlertStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("62")).Render

type keymap struct {
	Create         key.Binding
	Enter          key.Binding
	Rename         key.Binding
	Delete         key.Binding
	Back           key.Binding
	Description    key.Binding
	Contacts       key.Binding
	Comments       key.Binding
	Attachments    key.Binding
	Download       key.Binding
	ViewAttachment key.Binding
}

// Keymap reusable key mappings shared across models
var Keymap = keymap{
	Contacts: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "contacts"),
	),

	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	),
	Attachments: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "attachments"),
	),
	Description: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "description"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back"),
	),
	Comments: key.NewBinding(
		key.WithKeys("C"),
		key.WithHelp("C", "comments"),
	),
	Download: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "download attachment"),
	),
	ViewAttachment: key.NewBinding(
		key.WithKeys("v"),
		key.WithHelp("v", "view attachment"),
	),
}
