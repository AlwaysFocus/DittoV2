package contactui

import "github.com/charmbracelet/bubbles/key"

// define our keymap
// this is used to tell the program what to do when certain keys are pressed
type keymap = struct {
	primary, owner, internal, sendEmail, emailAll key.Binding
}
