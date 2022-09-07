package attachmentui

import (
	"DittoV2/casedetail"
	tea "github.com/charmbracelet/bubbletea"
)

func downloadAttachmentCmd(attachment casedetail.Attachment) tea.Cmd {
	return func() tea.Msg {
		return DownloadAttachmentMsg{Attachment: attachment}
	}
}
