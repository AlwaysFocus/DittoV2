package _case

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	format string = "%d : %s\n"
)

// Case the case holds primary information for a case
type Case struct {
	Company            string      `json:"Company"`
	HDCaseNum          int         `json:"HDCaseNum"`
	ParentCase         int         `json:"ParentCase"`
	CaseDescription    string      `json:"Description"`
	ResolutionText     string      `json:"ResolutionText"`
	PublishedText      string      `json:"PublishedText"`
	PublishedSummary   string      `json:"PublishedSummary"`
	KBEntry            bool        `json:"KBEntry"`
	PublishedItem      bool        `json:"PublishedItem"`
	PartNum            string      `json:"PartNum"`
	SerialNumber       string      `json:"SerialNumber"`
	QuoteNum           int         `json:"QuoteNum"`
	OrderNum           int         `json:"OrderNum"`
	CaseOwner          string      `json:"CaseOwner"`
	CreatedDate        string      `json:"CreatedDate"`
	CreatedBy          string      `json:"CreatedBy"`
	CreatedTime        int         `json:"CreatedTime"`
	LastUpdatedBy      string      `json:"LastUpdatedBy"`
	LastUpdatedDate    string      `json:"LastUpdatedDate"`
	LastUpdatedTime    int         `json:"LastUpdatedTime"`
	TopicID1           string      `json:"TopicID1"`
	TopicID2           string      `json:"TopicID2"`
	TopicID3           string      `json:"TopicID3"`
	TopicID4           string      `json:"TopicID4"`
	TopicID5           string      `json:"TopicID5"`
	TopicID6           string      `json:"TopicID6"`
	TopicID7           string      `json:"TopicID7"`
	TopicID8           string      `json:"TopicID8"`
	TopicID9           string      `json:"TopicID9"`
	TopicID10          string      `json:"TopicID10"`
	CaseTopics         string      `json:"CaseTopics"`
	RevisionNum        string      `json:"RevisionNum"`
	PartDescription    string      `json:"PartDescription"`
	Quantity           float64     `json:"Quantity"`
	QuantityUOM        string      `json:"QuantityUOM"`
	QuoteLine          int         `json:"QuoteLine"`
	InvoiceNum         int         `json:"InvoiceNum"`
	InvoiceLine        int         `json:"InvoiceLine"`
	PackNum            int         `json:"PackNum"`
	PackLine           int         `json:"PackLine"`
	ChangedBy          string      `json:"ChangedBy"`
	ChangeDate         string      `json:"ChangeDate"`
	ChangeTime         int         `json:"ChangeTime"`
	CompletedBy        string      `json:"CompletedBy"`
	CompletionDate     string      `json:"CompletionDate"`
	CompletionTime     int         `json:"CompletionTime"`
	UnitPrice          float64     `json:"UnitPrice"`
	DocUnitPrice       float64     `json:"DocUnitPrice"`
	Rp1ExtPrice        float64     `json:"Rp1ExtPrice"`
	Rp2ExtPrice        float64     `json:"Rp2ExtPrice"`
	CaseTypeID         string      `json:"CaseTypeID"`
	SysRevID           int         `json:"SysRevID"`
	SysRowID           string      `json:"SysRowID"`
	AttrCodeList       string      `json:"AttrCodeList"`
	IssueSummary       string      `json:"IssueSummary"`
	IssueText          string      `json:"IssueText"`
	DispCreateTime     string      `json:"DispCreateTime"`
	DispLastUpdateTime string      `json:"DispLastUpdateTime"`
	CaseOwnerName      string      `json:"CaseOwnerName"`
	CaseCode           string      `json:"CaseCode"`
	NextReviewDate     interface{} `json:"NextReviewDate"`
	EvaluationStatus   string      `json:"EvaluationStatus"`
	RowMod             string      `json:"RowMod"`
	UDSysRevID         interface{} `json:"UD_SysRevID"`
	CaseStatusC        interface{} `json:"CaseStatus_c"`
	StartDateC         interface{} `json:"StartDate_c"`
	RequestDateC       interface{} `json:"RequestDate_c"`
	ProjectPhaseC      interface{} `json:"ProjectPhase_c"`
	PhaseOperC         interface{} `json:"PhaseOper_c"`
}

// CaseList is used to unmarshal the case list from the HelpDesk endpoint
type CaseList struct {
	OdataContext string `json:"@odata.context"`
	Value        []Case `json:"value"`
}

// HelpDesk Primary type for this module
type HelpDesk struct {
	cases        []Case
	selectedCase Case
}

// NewHelpDesk create a new HelpDesk instance
func NewHelpDesk() *HelpDesk {

	// TODO: Not sure if there is a better way to do this
	hd := &HelpDesk{}
	// Get all the cases
	cases, err := hd.GetAllCases()
	if err != nil {
		log.Fatal(err)
	}

	return &HelpDesk{
		cases: cases,
	}
}

// Itos helper function to convert int to string
func Itos(i int) string {
	return fmt.Sprintf("%d", i)
}

// Implement list.Item for Bubbletea TUI

// Title the case title to display in a list
func (c Case) Title() string { return Itos(c.HDCaseNum) }

// Description the case description to display in a list
func (c Case) Description() string { return fmt.Sprintf("%s", c.CaseDescription) }

// FilterValue choose what field to use for filtering in a Bubbletea list component
func (c Case) FilterValue() string { return Itos(c.HDCaseNum) }

type (
	// CaseManagement Operations for Cases
	CaseManagement interface {
		PrintCases()
		HasContacts() bool
		GetCaseByID(CaseId uint) (Case, error)
		GetAllCases() ([]Case, error)
	}
)

// GetCaseByID get a case by HDCaseNum
func (c *HelpDesk) GetCaseByID(hdCaseNum uint) (Case, error) {

	// Get the necessary ENV variables
	apiKey := os.Getenv("EPICOR_API_KEY")
	authorization := os.Getenv("BASIC_AUTH")
	baseUrl := os.Getenv("BASE_EPICOR_URL")

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", baseUrl+"/Erp.BO.HelpDeskSvc/List?$filter=HDCaseNum%20eq%20"+Itos(int(hdCaseNum)), nil)
	// Headers
	req.Header.Add("X-API-Key", apiKey)
	req.Header.Add("Authorization", authorization)

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
		fmt.Println(parseFormErr)
	}

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := io.ReadAll(resp.Body)

	// Unmarshall the response body into a CaseList struct
	var caseList CaseList
	if err := json.Unmarshal(respBody, &caseList); err != nil {
		fmt.Println("Unable to parse CaseList struc", err)
	}

	// Log that we have found the case
	log.Println("Retrieved case: ", caseList.Value[0].HDCaseNum)

	// Get the first case in the list, we will only ever have one case since we specified the HDCaseNum
	var hdCase Case
	hdCase = caseList.Value[0]

	return hdCase, nil
}

// PrintCases print all cases to the console
func (c *HelpDesk) PrintCases() {
	cases, err := c.GetAllCases()
	if err != nil {
		log.Fatal(err)
	}

	for _, caseItem := range cases {
		fmt.Println(caseItem.HDCaseNum)
	}
}

// GetAllCases get all cases
func (c *HelpDesk) GetAllCases() ([]Case, error) {
	// Get the necessary ENV variables
	apiKey := os.Getenv("EPICOR_API_KEY")
	authorization := os.Getenv("BASIC_AUTH")
	baseUrl := os.Getenv("BASE_EPICOR_URL")

	// Url for retrieving all the cases
	url := baseUrl + "/Erp.BO.HelpDeskSvc/List?$top=99999"
	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", url, nil)

	// Headers
	req.Header.Add("X-API-Key", apiKey)
	req.Header.Add("Authorization", authorization)

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := io.ReadAll(resp.Body)

	// create a variable that will hold the unmarshalled json
	var caseList CaseList
	// unmarshal the json into the variable
	if err := json.Unmarshal(respBody, &caseList); err != nil {
		log.Println("Failure unmarshalling json: ", err)
	}

	var cases = caseList.Value

	return cases, nil
}
