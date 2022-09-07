package attachmentui

import (
	"DittoV2/casedetail"
	"DittoV2/tui/constants"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// define our keymap
// this is used to tell the program what to do when certain keys are pressed
type keymap = struct {
	download key.Binding
}

// BackMsg used to navigate back to the details view
type BackMsg bool

// DownloadAttachmentMsg the message to download the attachment
type DownloadAttachmentMsg struct {
	Attachment casedetail.Attachment
}

// Model the Attachment List model Definition
type Model struct {
	attachmentList list.Model
	keymap         keymap
	help           help.Model
}

// Init run and initial IO on program start
func (m Model) Init() tea.Cmd {
	m.attachmentList.SetShowFilter(true)
	m.attachmentList.SetFilteringEnabled(true)
	m.attachmentList.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			constants.Keymap.Download,
			constants.Keymap.Back,
		}
	}
	return nil
}

// Update handle IO and commands
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		top, right, bottom, left := constants.DocStyle.GetMargin()
		m.attachmentList.SetSize(msg.Width-left-right, msg.Height-top-bottom-1)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, constants.Keymap.Download):
			cmd = downloadAttachmentCmd(m.getSelectedAttachment())
			cmds = append(cmds, cmd)

		case key.Matches(msg, constants.Keymap.Back):
			return m, func() tea.Msg {
				return BackMsg(true)
			}
		}

	// If we received the Download Msg then download the attachment to the current dir
	case DownloadAttachmentMsg:
		// Download the attachment to the current dir
		sendDownloadFile(msg.Attachment)

	}
	m.attachmentList, cmd = m.attachmentList.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

// View the Attachment List view Definition
func (m Model) View() string {
	return m.attachmentList.View()
}

func (m Model) getSelectedAttachment() casedetail.Attachment {
	var item = m.attachmentList.SelectedItem()
	return item.(casedetail.Attachment)
}

// New initialize the attachment ui model for the program
func New(attachmentList []casedetail.Attachment) tea.Model {
	items := attachmentsToItems(attachmentList)

	const defaultWidth = 140
	m := Model{
		attachmentList: list.New(items, list.NewDefaultDelegate(), defaultWidth, 30),
		keymap: keymap{
			download: key.NewBinding(key.WithKeys("d"), key.WithHelp("d", "download")),
		},
	}

	m.attachmentList.Title = "Attachments"
	m.attachmentList.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			m.keymap.download,
			constants.Keymap.Back,
		}
	}

	return m
}

// Convert our attachments to items for the list model
func attachmentsToItems(attachmentList []casedetail.Attachment) []list.Item {
	var items []list.Item
	for _, attachment := range attachmentList {
		items = append(items, attachment)
	}
	return items
}

// DownloadFile contains the struct for the download file message
type DownloadFile struct {
	ReturnObj string `json:"returnObj"`
}

// sendDownloadFile send the download file request to the server
func sendDownloadFile(attachment casedetail.Attachment) {
	// Get the necessary ENV variables
	apiKey := os.Getenv("EPICOR_API_KEY")
	authorization := os.Getenv("BASIC_AUTH")
	baseUrl := os.Getenv("BASE_EPICOR_URL")

	jsonRequest := []byte(`{"xFileRefNum": ` + fmt.Sprintf("%d", attachment.XFileRefNum) + `}`)
	body := bytes.NewBuffer(jsonRequest)

	// Create client
	client := &http.Client{}

	url := baseUrl + "Ice.BO.AttachmentSvc/DownloadFile"
	// Create request
	req, err := http.NewRequest("POST", url, body)

	// Headers
	// TODO: Move to env variables
	req.Header.Add("X-API-Key", apiKey)
	req.Header.Add("Authorization", authorization)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
		log.Println(parseFormErr)
	}

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := io.ReadAll(resp.Body)

	// unmarshal the json_request response
	var downloadFile DownloadFile

	jsonErr := json.Unmarshal(respBody, &downloadFile)
	if jsonErr != nil {
		log.Println(jsonErr)
	}

	// Decode the base64 file content
	decodedFileContent, decodeErr := base64.StdEncoding.DecodeString(downloadFile.ReturnObj)
	if decodeErr != nil {
		log.Fatal(decodeErr)
	}

	// Get the file name
	fileName := strings.Split(attachment.FileName, "\\")[len(strings.Split(attachment.FileName, "\\"))-1]

	log.Println("File Name: ", fileName)

	// write the file
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	_, err = file.Write(decodedFileContent)
	if err != nil {
		log.Fatal(err)
	}

}
