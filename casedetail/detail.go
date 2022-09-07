package casedetail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Detail the detail model
type Detail struct {
	CaseNum int
	CaseDetail
}

type CaseDetail struct {
	ReturnObj struct {
		HDCase []struct {
			Company                       string      `json:"Company"`
			HDCaseNum                     int         `json:"HDCaseNum"`
			CustNum                       int         `json:"CustNum"`
			ShipToNum                     string      `json:"ShipToNum"`
			ShpConNum                     int         `json:"ShpConNum"`
			ParentCase                    int         `json:"ParentCase"`
			Description                   string      `json:"Description"`
			ResolutionText                string      `json:"ResolutionText"`
			PublishedText                 string      `json:"PublishedText"`
			PublishedSummary              string      `json:"PublishedSummary"`
			KBEntry                       bool        `json:"KBEntry"`
			PublishedItem                 bool        `json:"PublishedItem"`
			PartNum                       string      `json:"PartNum"`
			SerialNumber                  string      `json:"SerialNumber"`
			QuoteNum                      int         `json:"QuoteNum"`
			OrderNum                      int         `json:"OrderNum"`
			CallNum                       int         `json:"CallNum"`
			ContractNum                   int         `json:"ContractNum"`
			WarrantyCode                  string      `json:"WarrantyCode"`
			Priority                      int         `json:"Priority"`
			TaskSetID                     string      `json:"TaskSetID"`
			CurrentWFStageID              string      `json:"CurrentWFStageID"`
			ActiveTaskID                  string      `json:"ActiveTaskID"`
			LastTaskID                    string      `json:"LastTaskID"`
			CaseOwner                     string      `json:"CaseOwner"`
			WFGroupID                     string      `json:"WFGroupID"`
			WFComplete                    bool        `json:"WFComplete"`
			CreatedDate                   string      `json:"CreatedDate"`
			CreatedBy                     string      `json:"CreatedBy"`
			CreatedTime                   int         `json:"CreatedTime"`
			LastUpdatedBy                 string      `json:"LastUpdatedBy"`
			LastUpdatedDate               string      `json:"LastUpdatedDate"`
			LastUpdatedTime               int         `json:"LastUpdatedTime"`
			TopicID1                      string      `json:"TopicID1"`
			TopicID2                      string      `json:"TopicID2"`
			TopicID3                      string      `json:"TopicID3"`
			TopicID4                      string      `json:"TopicID4"`
			TopicID5                      string      `json:"TopicID5"`
			TopicID6                      string      `json:"TopicID6"`
			TopicID7                      string      `json:"TopicID7"`
			TopicID8                      string      `json:"TopicID8"`
			TopicID9                      string      `json:"TopicID9"`
			TopicID10                     string      `json:"TopicID10"`
			CaseTopics                    string      `json:"CaseTopics"`
			MktgCampaignID                string      `json:"MktgCampaignID"`
			MktgEvntSeq                   int         `json:"MktgEvntSeq"`
			RevisionNum                   string      `json:"RevisionNum"`
			PartDescription               string      `json:"PartDescription"`
			Quantity                      float64     `json:"Quantity"`
			QuantityUOM                   string      `json:"QuantityUOM"`
			OrderLine                     int         `json:"OrderLine"`
			OrderRelNum                   int         `json:"OrderRelNum"`
			QuoteLine                     int         `json:"QuoteLine"`
			CallLine                      int         `json:"CallLine"`
			RMANum                        int         `json:"RMANum"`
			RMALine                       int         `json:"RMALine"`
			InvoiceNum                    int         `json:"InvoiceNum"`
			InvoiceLine                   int         `json:"InvoiceLine"`
			PrcConNum                     int         `json:"PrcConNum"`
			ProjectID                     string      `json:"ProjectID"`
			CustomerName                  string      `json:"CustomerName"`
			PackNum                       int         `json:"PackNum"`
			PackLine                      int         `json:"PackLine"`
			ChangedBy                     string      `json:"ChangedBy"`
			ChangeDate                    string      `json:"ChangeDate"`
			ChangeTime                    int         `json:"ChangeTime"`
			CompletedBy                   string      `json:"CompletedBy"`
			CompletionDate                interface{} `json:"CompletionDate"`
			CompletionTime                int         `json:"CompletionTime"`
			ShipToCustNum                 int         `json:"ShipToCustNum"`
			DropShipPackSlip              string      `json:"DropShipPackSlip"`
			DropShipPackLine              int         `json:"DropShipPackLine"`
			VendorNum                     int         `json:"VendorNum"`
			PurPoint                      string      `json:"PurPoint"`
			EquipID                       string      `json:"EquipID"`
			EmpID                         string      `json:"EmpID"`
			BuyerID                       string      `json:"BuyerID"`
			VendorNumCon                  int         `json:"VendorNumCon"`
			PurPointCon                   string      `json:"PurPointCon"`
			VenConNum                     int         `json:"VenConNum"`
			PurPointConNum                int         `json:"PurPointConNum"`
			UnitPrice                     float64     `json:"UnitPrice"`
			DocUnitPrice                  float64     `json:"DocUnitPrice"`
			Rpt1UnitPrice                 float64     `json:"Rpt1UnitPrice"`
			Rpt2UnitPrice                 float64     `json:"Rpt2UnitPrice"`
			Rpt3UnitPrice                 float64     `json:"Rpt3UnitPrice"`
			ExtPrice                      float64     `json:"ExtPrice"`
			DocExtPrice                   float64     `json:"DocExtPrice"`
			Rp1ExtPrice                   float64     `json:"Rp1ExtPrice"`
			Rp2ExtPrice                   float64     `json:"Rp2ExtPrice"`
			Rp3ExtPrice                   float64     `json:"Rp3ExtPrice"`
			CurrencyCode                  string      `json:"CurrencyCode"`
			RateGrpCode                   string      `json:"RateGrpCode"`
			LockRate                      bool        `json:"LockRate"`
			ExchangeRate                  float64     `json:"ExchangeRate"`
			CaseTypeID                    string      `json:"CaseTypeID"`
			PONum                         int         `json:"PONum"`
			TerritoryID                   string      `json:"TerritoryID"`
			POLine                        int         `json:"POLine"`
			WorkflowType                  string      `json:"WorkflowType"`
			POPackSlip                    string      `json:"POPackSlip"`
			POPackLine                    int         `json:"POPackLine"`
			SysRevID                      int         `json:"SysRevID"`
			SysRowID                      string      `json:"SysRowID"`
			HDCaseStatus                  string      `json:"HDCaseStatus"`
			ReqPerConID                   int         `json:"ReqPerConID"`
			PerConID                      int         `json:"PerConID"`
			WebCase                       bool        `json:"WebCase"`
			WebComment                    string      `json:"WebComment"`
			IDNum                         string      `json:"IDNum"`
			LocationNum                   int         `json:"LocationNum"`
			AllowMilestoneUpdate          bool        `json:"AllowMilestoneUpdate"`
			AttrCodeList                  string      `json:"AttrCodeList"`
			AvailablePrcConNum            string      `json:"AvailablePrcConNum"`
			AvailablePurPointConNum       string      `json:"AvailablePurPointConNum"`
			AvailableShpConNum            string      `json:"AvailableShpConNum"`
			AvailableTaskSets             string      `json:"AvailableTaskSets"`
			AvailableVenConNum            string      `json:"AvailableVenConNum"`
			BaseCurrencyID                string      `json:"BaseCurrencyID"`
			BaseCurrSymbol                string      `json:"BaseCurrSymbol"`
			CaseCode                      string      `json:"CaseCode"`
			CaseStatus                    string      `json:"CaseStatus"`
			ChildCases                    string      `json:"ChildCases"`
			CurrentMilestone              int         `json:"CurrentMilestone"`
			CurrentMilestoneDesc          string      `json:"CurrentMilestoneDesc"`
			CustCntCorpName               string      `json:"CustCntCorpName"`
			CustCntEMail                  string      `json:"CustCntEMail"`
			CustCntFaxNum                 string      `json:"CustCntFaxNum"`
			CustCntFirstName              string      `json:"CustCntFirstName"`
			CustCntLastName               string      `json:"CustCntLastName"`
			CustCntMiddleName             string      `json:"CustCntMiddleName"`
			CustCntName                   string      `json:"CustCntName"`
			CustCntPhoneNum               string      `json:"CustCntPhoneNum"`
			CustomerRequiresPO            bool        `json:"CustomerRequiresPO"`
			DispCreateTime                string      `json:"DispCreateTime"`
			DispLastUpdateTime            string      `json:"DispLastUpdateTime"`
			DropShip                      bool        `json:"DropShip"`
			EvaluationStatus              string      `json:"EvaluationStatus"`
			EvaluationStatusDesc          string      `json:"EvaluationStatusDesc"`
			HDCaseNumString               string      `json:"HDCaseNumString"`
			Inactive                      bool        `json:"Inactive"`
			IssueSummary                  string      `json:"IssueSummary"`
			IssueText                     string      `json:"IssueText"`
			NextReviewDate                interface{} `json:"NextReviewDate"`
			PartSalesUM                   string      `json:"PartSalesUM"`
			PPCntEmailAddress             string      `json:"PPCntEmailAddress"`
			PPCntFaxNum                   string      `json:"PPCntFaxNum"`
			PPCntName                     string      `json:"PPCntName"`
			PPCntPhoneNum                 string      `json:"PPCntPhoneNum"`
			PricePerCode                  string      `json:"PricePerCode"`
			PurPointConName               string      `json:"PurPointConName"`
			ReqContextLink                string      `json:"ReqContextLink"`
			ReqPerConLnkID1               string      `json:"ReqPerConLnkID1"`
			ReqPerConLnkID2               string      `json:"ReqPerConLnkID2"`
			ReqPerConLnkName              string      `json:"ReqPerConLnkName"`
			ReqPerConLnkRowID             string      `json:"ReqPerConLnkRowID"`
			ReqPerConName                 string      `json:"ReqPerConName"`
			ReqPrimary                    bool        `json:"ReqPrimary"`
			Rpt1ExtPrice                  float64     `json:"Rpt1ExtPrice"`
			Rpt2ExtPrice                  float64     `json:"Rpt2ExtPrice"`
			Rpt3ExtPrice                  float64     `json:"Rpt3ExtPrice"`
			ShipCntCorpName               string      `json:"ShipCntCorpName"`
			ShipCntEMail                  string      `json:"ShipCntEMail"`
			ShipCntFaxNum                 string      `json:"ShipCntFaxNum"`
			ShipCntFirstName              string      `json:"ShipCntFirstName"`
			ShipCntLastName               string      `json:"ShipCntLastName"`
			ShipCntMiddleName             string      `json:"ShipCntMiddleName"`
			ShipCntName                   string      `json:"ShipCntName"`
			ShipCntPhoneNum               string      `json:"ShipCntPhoneNum"`
			ShipToCustID                  string      `json:"ShipToCustID"`
			ShipToNumName                 string      `json:"ShipToNumName"`
			TargetUOM                     string      `json:"TargetUOM"`
			TaskCompletePasswordIsValid   bool        `json:"TaskCompletePasswordIsValid"`
			TaskCompletePasswordRequired  bool        `json:"TaskCompletePasswordRequired"`
			VendCntEmailAddress           string      `json:"VendCntEmailAddress"`
			VendCntFaxNum                 string      `json:"VendCntFaxNum"`
			VendCntName                   string      `json:"VendCntName"`
			VendCntPhoneNum               string      `json:"VendCntPhoneNum"`
			WebQuoteNum                   int         `json:"WebQuoteNum"`
			AvailableMilestones           string      `json:"AvailableMilestones"`
			BitFlag                       int         `json:"BitFlag"`
			ActiveTaskIDTaskDescription   string      `json:"ActiveTaskIDTaskDescription"`
			BuyerIDName                   string      `json:"BuyerIDName"`
			CaseOwnerName                 string      `json:"CaseOwnerName"`
			CurrencyCodeCurrName          string      `json:"CurrencyCodeCurrName"`
			CurrencyCodeCurrSymbol        string      `json:"CurrencyCodeCurrSymbol"`
			CurrencyCodeCurrDesc          string      `json:"CurrencyCodeCurrDesc"`
			CurrencyCodeDocumentDesc      string      `json:"CurrencyCodeDocumentDesc"`
			CurrencyCodeCurrencyID        string      `json:"CurrencyCodeCurrencyID"`
			CustNumBTName                 string      `json:"CustNumBTName"`
			CustNumCustID                 string      `json:"CustNumCustID"`
			CustNumName                   string      `json:"CustNumName"`
			CustNumAllowShipTo3           bool        `json:"CustNumAllowShipTo3"`
			DropShipDtlLineDesc           string      `json:"DropShipDtlLineDesc"`
			EmpIDName                     string      `json:"EmpIDName"`
			EquipIDDescription            string      `json:"EquipIDDescription"`
			LastTaskIDTaskDescription     string      `json:"LastTaskIDTaskDescription"`
			LocationInventoryLotNum       string      `json:"LocationInventoryLotNum"`
			LocationInventorySerialNumber string      `json:"LocationInventorySerialNumber"`
			LocationInventoryIDNum        string      `json:"LocationInventoryIDNum"`
			MktgCampaignIDCampDescription string      `json:"MktgCampaignIDCampDescription"`
			MktgEventEvntDescription      string      `json:"MktgEventEvntDescription"`
			PackLineLineDesc              string      `json:"PackLineLineDesc"`
			PartNumTrackDimension         bool        `json:"PartNumTrackDimension"`
			PartNumSalesUM                string      `json:"PartNumSalesUM"`
			PartNumPartDescription        string      `json:"PartNumPartDescription"`
			PartNumPricePerCode           string      `json:"PartNumPricePerCode"`
			PartNumSellingFactor          float64     `json:"PartNumSellingFactor"`
			PartNumTrackSerialNum         bool        `json:"PartNumTrackSerialNum"`
			PartNumTrackLots              bool        `json:"PartNumTrackLots"`
			PartNumIUM                    string      `json:"PartNumIUM"`
			ProjectIDDescription          string      `json:"ProjectIDDescription"`
			ShipToCustNumName             string      `json:"ShipToCustNumName"`
			ShipToCustNumCustID           string      `json:"ShipToCustNumCustID"`
			TaskSetIDWorkflowType         string      `json:"TaskSetIDWorkflowType"`
			TaskSetIDTaskSetDescription   string      `json:"TaskSetIDTaskSetDescription"`
			TerritoryIDTerritoryDesc      string      `json:"TerritoryIDTerritoryDesc"`
			VendorNumConVendorID          string      `json:"VendorNumConVendorID"`
			VendorNumConName              string      `json:"VendorNumConName"`
			WarrantyCodeWarrDescription   string      `json:"WarrantyCodeWarrDescription"`
			WFGroupIDDescription          string      `json:"WFGroupIDDescription"`
			WFStageIDDescription          string      `json:"WFStageIDDescription"`
			RowMod                        string      `json:"RowMod"`
			CaseStatusC                   string      `json:"CaseStatus_c"`
			PhaseOperC                    int         `json:"PhaseOper_c"`
			ProjectPhaseC                 string      `json:"ProjectPhase_c"`
			UDSysRevID                    string      `json:"UD_SysRevID"`
		} `json:"HDCase"`
		HDCaseAttch  []Attachment  `json:"HDCaseAttch"`
		HDCaseFSCall []interface{} `json:"HDCaseFSCall"`
		HDCaseJob    []interface{} `json:"HDCaseJob"`
		HDCaseLink   []interface{} `json:"HDCaseLink"`
		HDCaseOrder  []struct {
			Company    string `json:"Company"`
			OrderNum   int    `json:"OrderNum"`
			HDCaseNum  int    `json:"HDCaseNum"`
			SysRowID   string `json:"SysRowID"`
			BitFlag    int    `json:"BitFlag"`
			RowMod     string `json:"RowMod"`
			UDSysRevID string `json:"UD_SysRevID"`
			ValidatedC bool   `json:"Validated_c"`
		} `json:"HDCaseOrder"`
		HDCaseQuote []struct {
			Company   string `json:"Company"`
			QuoteNum  int    `json:"QuoteNum"`
			HDCaseNum int    `json:"HDCaseNum"`
			SysRowID  string `json:"SysRowID"`
			BitFlag   int    `json:"BitFlag"`
			RowMod    string `json:"RowMod"`
		} `json:"HDCaseQuote"`
		HDCaseRMA    []interface{} `json:"HDCaseRMA"`
		HDChildCases []interface{} `json:"HDChildCases"`
		HDContact    []struct {
			Company        string `json:"Company"`
			HDCaseNum      int    `json:"HDCaseNum"`
			PerConLnkRowID string `json:"PerConLnkRowID"`
			Primary        bool   `json:"Primary"`
			Comment        string `json:"Comment"`
			SysRevID       int    `json:"SysRevID"`
			SysRowID       string `json:"SysRowID"`
			Requestor      bool   `json:"Requestor"`
			Name           string `json:"Name"`
			City           string `json:"City"`
			State          string `json:"State"`
			Zip            string `json:"Zip"`
			Address1       string `json:"Address1"`
			Address2       string `json:"Address2"`
			Address3       string `json:"Address3"`
			ContextLink    string `json:"ContextLink"`
			BuyerID        string `json:"BuyerID"`
			BuyerName      string `json:"BuyerName"`
			CustID         string `json:"CustID"`
			CustName       string `json:"CustName"`
			EmpID          string `json:"EmpID"`
			EmpName        string `json:"EmpName"`
			PurPoint       string `json:"PurPoint"`
			PurPointName   string `json:"PurPointName"`
			SalesRepCode   string `json:"SalesRepCode"`
			SalesRepName   string `json:"SalesRepName"`
			ShipToName     string `json:"ShipToName"`
			ShipToNum      string `json:"ShipToNum"`
			VendorID       string `json:"VendorID"`
			VendorName     string `json:"VendorName"`
			Selected       bool   `json:"Selected"`
			BitFlag        int    `json:"BitFlag"`
			RowMod         string `json:"RowMod"`
		} `json:"HDContact"`
		HDCaseMaintReq  []interface{} `json:"HDCaseMaintReq"`
		ExtensionTables []interface{} `json:"ExtensionTables"`
	} `json:"returnObj"`
}

func (a Attachment) Title() string { return a.DocTypeID }

// Description the case description to display in a list
func (a Attachment) Description() string { return fmt.Sprintf("%s", a.FileName) }

// FilterValue choose what field to use for filtering in a Bubbletea list component
func (a Attachment) FilterValue() string { return a.FileName }

// HelpDesk the functions for HelpDesk Details
type HelpDesk interface {
	GetCaseDetails(caseNum int) error
	GetCaseContacts(caseNum int) error
	SendEmail(caseNum int, recipients []string, messageBody string) (string, error)
	AddComment(caseNum int, comment string) error
	GetTimeEntries(caseNum int) error
}

type Comment struct {
	Company            string `json:"Company"`
	Comment            string `json:"Character01"`
	CaseNum            string `json:"Key1"`
	UserId             string `json:"ShortChar01"`
	CommentSequence    int    `json:"Number01"`
	CommentSequenceKey string `json:"Key5"`
	CommentDate        string `json:"Date01"`
}

type UD01TableSetResponse struct {
	Parameters struct {
		Ds struct {
			UD01 []struct {
				Company      string      `json:"Company"`
				Key1         string      `json:"Key1"`
				Key2         string      `json:"Key2"`
				Key3         string      `json:"Key3"`
				Key4         string      `json:"Key4"`
				Key5         string      `json:"Key5"`
				Character01  string      `json:"Character01"`
				Character02  string      `json:"Character02"`
				Character03  string      `json:"Character03"`
				Character04  string      `json:"Character04"`
				Character05  string      `json:"Character05"`
				Character06  string      `json:"Character06"`
				Character07  string      `json:"Character07"`
				Character08  string      `json:"Character08"`
				Character09  string      `json:"Character09"`
				Character10  string      `json:"Character10"`
				Number01     float64     `json:"Number01"`
				Number02     float64     `json:"Number02"`
				Number03     float64     `json:"Number03"`
				Number04     float64     `json:"Number04"`
				Number05     float64     `json:"Number05"`
				Number06     float64     `json:"Number06"`
				Number07     float64     `json:"Number07"`
				Number08     float64     `json:"Number08"`
				Number09     float64     `json:"Number09"`
				Number10     float64     `json:"Number10"`
				Number11     float64     `json:"Number11"`
				Number12     float64     `json:"Number12"`
				Number13     float64     `json:"Number13"`
				Number14     float64     `json:"Number14"`
				Number15     float64     `json:"Number15"`
				Number16     float64     `json:"Number16"`
				Number17     float64     `json:"Number17"`
				Number18     float64     `json:"Number18"`
				Number19     float64     `json:"Number19"`
				Number20     float64     `json:"Number20"`
				Date01       interface{} `json:"Date01"`
				Date02       interface{} `json:"Date02"`
				Date03       interface{} `json:"Date03"`
				Date04       interface{} `json:"Date04"`
				Date05       interface{} `json:"Date05"`
				Date06       interface{} `json:"Date06"`
				Date07       interface{} `json:"Date07"`
				Date08       interface{} `json:"Date08"`
				Date09       interface{} `json:"Date09"`
				Date10       interface{} `json:"Date10"`
				Date11       interface{} `json:"Date11"`
				Date12       interface{} `json:"Date12"`
				Date13       interface{} `json:"Date13"`
				Date14       interface{} `json:"Date14"`
				Date15       interface{} `json:"Date15"`
				Date16       interface{} `json:"Date16"`
				Date17       interface{} `json:"Date17"`
				Date18       interface{} `json:"Date18"`
				Date19       interface{} `json:"Date19"`
				Date20       interface{} `json:"Date20"`
				CheckBox01   bool        `json:"CheckBox01"`
				CheckBox02   bool        `json:"CheckBox02"`
				CheckBox03   bool        `json:"CheckBox03"`
				CheckBox04   bool        `json:"CheckBox04"`
				CheckBox05   bool        `json:"CheckBox05"`
				CheckBox06   bool        `json:"CheckBox06"`
				CheckBox07   bool        `json:"CheckBox07"`
				CheckBox08   bool        `json:"CheckBox08"`
				CheckBox09   bool        `json:"CheckBox09"`
				CheckBox10   bool        `json:"CheckBox10"`
				CheckBox11   bool        `json:"CheckBox11"`
				CheckBox12   bool        `json:"CheckBox12"`
				CheckBox13   bool        `json:"CheckBox13"`
				CheckBox14   bool        `json:"CheckBox14"`
				CheckBox15   bool        `json:"CheckBox15"`
				CheckBox16   bool        `json:"CheckBox16"`
				CheckBox17   bool        `json:"CheckBox17"`
				CheckBox18   bool        `json:"CheckBox18"`
				CheckBox19   bool        `json:"CheckBox19"`
				CheckBox20   bool        `json:"CheckBox20"`
				ShortChar01  string      `json:"ShortChar01"`
				ShortChar02  string      `json:"ShortChar02"`
				ShortChar03  string      `json:"ShortChar03"`
				ShortChar04  string      `json:"ShortChar04"`
				ShortChar05  string      `json:"ShortChar05"`
				ShortChar06  string      `json:"ShortChar06"`
				ShortChar07  string      `json:"ShortChar07"`
				ShortChar08  string      `json:"ShortChar08"`
				ShortChar09  string      `json:"ShortChar09"`
				ShortChar10  string      `json:"ShortChar10"`
				ShortChar11  string      `json:"ShortChar11"`
				ShortChar12  string      `json:"ShortChar12"`
				ShortChar13  string      `json:"ShortChar13"`
				ShortChar14  string      `json:"ShortChar14"`
				ShortChar15  string      `json:"ShortChar15"`
				ShortChar16  string      `json:"ShortChar16"`
				ShortChar17  string      `json:"ShortChar17"`
				ShortChar18  string      `json:"ShortChar18"`
				ShortChar19  string      `json:"ShortChar19"`
				ShortChar20  string      `json:"ShortChar20"`
				GlobalUD01   bool        `json:"GlobalUD01"`
				GlobalLock   bool        `json:"GlobalLock"`
				SysRevID     int         `json:"SysRevID"`
				SysRowID     string      `json:"SysRowID"`
				BitFlag      int         `json:"BitFlag"`
				RowMod       string      `json:"RowMod"`
				CommentTextC string      `json:"CommentText_c"`
			} `json:"UD01"`
			UD01Attch       []interface{} `json:"UD01Attch"`
			ExtensionTables []interface{} `json:"ExtensionTables"`
		} `json:"ds"`
	} `json:"parameters"`
}

type Attachment struct {
	Company         string `json:"Company"`
	HDCaseNum       int    `json:"HDCaseNum"`
	DrawingSeq      int    `json:"DrawingSeq"`
	XFileRefNum     int    `json:"XFileRefNum"`
	SysRevID        int    `json:"SysRevID"`
	SysRowID        string `json:"SysRowID"`
	ForeignSysRowID string `json:"ForeignSysRowID"`
	DrawDesc        string `json:"DrawDesc"`
	FileName        string `json:"FileName"`
	PDMDocID        string `json:"PDMDocID"`
	DocTypeID       string `json:"DocTypeID"`
	RowMod          string `json:"RowMod"`
}

// HelpDeskDetail holds the details of a Help Desk case
type HelpDeskDetail struct {
	Contacts         []Contact
	Comments         []Comment
	Details          CaseDetail
	Attachments      []Attachment
	PrimaryContact   Contact
	CaseOwnerContact Contact
	InternalContact  Contact
}

type ContactType string

const (
	// Primary is the primary contact for the case
	Primary ContactType = "Primary"
	// Internal is an internal contact for the case
	Internal ContactType = "Internal"
	// CaseOwner is the owner of the case
	CaseOwner ContactType = "CaseOwner"
)

// Contact is a struct for the contact information
type Contact struct {
	Company        string `json:"Company"`
	HDCaseNum      int    `json:"HDCaseNum"`
	PerConLnkRowID string `json:"PerConLnkRowID"`
	Primary        bool   `json:"Primary"`
	Comment        string `json:"Comment"`
	SysRevID       int    `json:"SysRevID"`
	SysRowID       string `json:"SysRowID"`
	Requestor      bool   `json:"Requestor"`
	Name           string `json:"Name"`
	City           string `json:"City"`
	State          string `json:"State"`
	Zip            string `json:"Zip"`
	Address1       string `json:"Address1"`
	Address2       string `json:"Address2"`
	Address3       string `json:"Address3"`
	ContextLink    string `json:"ContextLink"`
	BuyerID        string `json:"BuyerID"`
	BuyerName      string `json:"BuyerName"`
	CustID         string `json:"CustID"`
	CustName       string `json:"CustName"`
	EmpID          string `json:"EmpID"`
	EmpName        string `json:"EmpName"`
	PurPoint       string `json:"PurPoint"`
	PurPointName   string `json:"PurPointName"`
	SalesRepCode   string `json:"SalesRepCode"`
	SalesRepName   string `json:"SalesRepName"`
	ShipToName     string `json:"ShipToName"`
	ShipToNum      string `json:"ShipToNum"`
	VendorID       string `json:"VendorID"`
	VendorName     string `json:"VendorName"`
	Selected       bool   `json:"Selected"`
	BitFlag        int    `json:"BitFlag"`
	RowMod         string `json:"RowMod"`
	Type           ContactType
	Email          string
}

// SalesRep is the structure related to work force entries
type SalesRep struct {
	InActive                bool    `json:"InActive"`
	Company                 string  `json:"Company"`
	SalesRepCode            string  `json:"SalesRepCode"`
	Name                    string  `json:"Name"`
	CommissionPercent       float64 `json:"CommissionPercent"`
	CommissionEarnedAt      bool    `json:"CommissionEarnedAt"`
	AlertFlag               bool    `json:"AlertFlag"`
	Address1                string  `json:"Address1"`
	Address2                string  `json:"Address2"`
	Address3                string  `json:"Address3"`
	City                    string  `json:"City"`
	State                   string  `json:"State"`
	Zip                     string  `json:"Zip"`
	Country                 string  `json:"Country"`
	CountryNum              int     `json:"CountryNum"`
	OfficePhoneNum          string  `json:"OfficePhoneNum"`
	FaxPhoneNum             string  `json:"FaxPhoneNum"`
	CellPhoneNum            string  `json:"CellPhoneNum"`
	PagerNum                string  `json:"PagerNum"`
	HomePhoneNum            string  `json:"HomePhoneNum"`
	EMailAddress            string  `json:"EMailAddress"`
	SalesRepTitle           string  `json:"SalesRepTitle"`
	RepReportsTo            string  `json:"RepReportsTo"`
	Comment                 string  `json:"Comment"`
	SalesMgrConfidence      int     `json:"SalesMgrConfidence"`
	RoleCode                string  `json:"RoleCode"`
	ViewAllTer              bool    `json:"ViewAllTer"`
	ViewCompPipe            bool    `json:"ViewCompPipe"`
	WebSaleGetsCommission   bool    `json:"WebSaleGetsCommission"`
	CnvEmpID                string  `json:"CnvEmpID"`
	PerConID                int     `json:"PerConID"`
	SyncNameToPerCon        bool    `json:"SyncNameToPerCon"`
	SyncAddressToPerCon     bool    `json:"SyncAddressToPerCon"`
	SyncPhoneToPerCon       bool    `json:"SyncPhoneToPerCon"`
	SyncEmailToPerCon       bool    `json:"SyncEmailToPerCon"`
	SyncLinksToPerCon       bool    `json:"SyncLinksToPerCon"`
	WebSite                 string  `json:"WebSite"`
	IM                      string  `json:"IM"`
	Twitter                 string  `json:"Twitter"`
	LinkedIn                string  `json:"LinkedIn"`
	FaceBook                string  `json:"FaceBook"`
	WebLink1                string  `json:"WebLink1"`
	WebLink2                string  `json:"WebLink2"`
	WebLink3                string  `json:"WebLink3"`
	WebLink4                string  `json:"WebLink4"`
	WebLink5                string  `json:"WebLink5"`
	MgrWorstCsPct           int     `json:"MgrWorstCsPct"`
	MgrBestCsPct            int     `json:"MgrBestCsPct"`
	SysRevID                int     `json:"SysRevID"`
	SysRowID                string  `json:"SysRowID"`
	WebSalesRep             bool    `json:"WebSalesRep"`
	ECCSalesRepCode         string  `json:"ECCSalesRepCode"`
	PerConName              string  `json:"PerConName"`
	RepReportsToName        string  `json:"RepReportsToName"`
	BitFlag                 int     `json:"BitFlag"`
	CountryNumDescription   string  `json:"CountryNumDescription"`
	RoleCodeRoleDescription string  `json:"RoleCodeRoleDescription"`
	RowMod                  string  `json:"RowMod"`
	SalesRepTypeC           string  `json:"SalesRepType_c"`
	UDSysRevID              string  `json:"UD_SysRevID"`
}

type SalesRepTableset struct {
	SalesRep      []SalesRep    `json:"SalesRep"`
	SalesRepAttch []interface{} `json:"SalesRepAttch"`
	SaleAuth      []struct {
		Company          string `json:"Company"`
		SalesRepCode     string `json:"SalesRepCode"`
		DcdUserID        string `json:"DcdUserID"`
		SysRevID         int    `json:"SysRevID"`
		SysRowID         string `json:"SysRowID"`
		DefaultUser      bool   `json:"DefaultUser"`
		BitFlag          int    `json:"BitFlag"`
		SalesRepCodeName string `json:"SalesRepCodeName"`
		UserIDName       string `json:"UserIDName"`
		RowMod           string `json:"RowMod"`
	} `json:"SaleAuth"`
	ExtensionTables []interface{} `json:"ExtensionTables"`
}

type SalesRepResponse struct {
	SalesRepTableset SalesRepTableset `json:"returnObj"`
}

type PerCon struct {
	Company                 string      `json:"Company"`
	PerConID                int         `json:"PerConID"`
	Name                    string      `json:"Name"`
	FirstName               string      `json:"FirstName"`
	MiddleName              string      `json:"MiddleName"`
	LastName                string      `json:"LastName"`
	PRName                  string      `json:"PRName"`
	Func                    string      `json:"Func"`
	FaxNum                  string      `json:"FaxNum"`
	PhoneNum                string      `json:"PhoneNum"`
	EMailAddress            string      `json:"EMailAddress"`
	CellPhoneNum            string      `json:"CellPhoneNum"`
	PagerNum                string      `json:"PagerNum"`
	HomeNum                 string      `json:"HomeNum"`
	AltNum                  string      `json:"AltNum"`
	Prefix                  string      `json:"Prefix"`
	Suffix                  string      `json:"Suffix"`
	Initials                string      `json:"Initials"`
	WebSite                 string      `json:"WebSite"`
	IM                      string      `json:"IM"`
	Twitter                 string      `json:"Twitter"`
	LinkedIn                string      `json:"LinkedIn"`
	FaceBook                string      `json:"FaceBook"`
	WebLink1                string      `json:"WebLink1"`
	WebLink2                string      `json:"WebLink2"`
	WebLink3                string      `json:"WebLink3"`
	WebLink4                string      `json:"WebLink4"`
	WebLink5                string      `json:"WebLink5"`
	Address1                string      `json:"Address1"`
	Address2                string      `json:"Address2"`
	Address3                string      `json:"Address3"`
	City                    string      `json:"City"`
	State                   string      `json:"State"`
	Zip                     string      `json:"Zip"`
	Country                 string      `json:"Country"`
	CountryNum              int         `json:"CountryNum"`
	CorpName                string      `json:"CorpName"`
	RoleCode                string      `json:"RoleCode"`
	Comment                 string      `json:"Comment"`
	ContactTitle            string      `json:"ContactTitle"`
	ReportsTo               string      `json:"ReportsTo"`
	PRLastName              string      `json:"PRLastName"`
	PRFirstName             string      `json:"PRFirstName"`
	PRMiddleName            string      `json:"PRMiddleName"`
	DcdUserID               string      `json:"DcdUserID"`
	PhotoFile               string      `json:"PhotoFile"`
	GlobalPerCon            bool        `json:"GlobalPerCon"`
	GlobalLock              bool        `json:"GlobalLock"`
	HCMLinked               bool        `json:"HCMLinked"`
	SysRevID                int         `json:"SysRevID"`
	SysRowID                string      `json:"SysRowID"`
	IssuerPrsnIDCode        string      `json:"IssuerPrsnIDCode"`
	IssuerIDType            string      `json:"IssuerIDType"`
	IssuerName              string      `json:"IssuerName"`
	IssuerSurname           string      `json:"IssuerSurname"`
	IssuerSerialNum         string      `json:"IssuerSerialNum"`
	IssuerIDIssDate         interface{} `json:"IssuerIDIssDate"`
	ImageID                 string      `json:"ImageID"`
	CrtBuyerLnk             bool        `json:"CrtBuyerLnk"`
	CrtEmpLnk               bool        `json:"CrtEmpLnk"`
	CrtPREmpLnk             bool        `json:"CrtPREmpLnk"`
	CrtWFLnk                bool        `json:"CrtWFLnk"`
	EnableGlbLock           bool        `json:"EnableGlbLock"`
	EnableGlbPerCon         bool        `json:"EnableGlbPerCon"`
	GlbFlag                 bool        `json:"GlbFlag"`
	GlbLink                 string      `json:"GlbLink"`
	PhotoFilePath           string      `json:"PhotoFilePath"`
	BitFlag                 int         `json:"BitFlag"`
	RoleCodeRoleDescription string      `json:"RoleCodeRoleDescription"`
	RowMod                  string      `json:"RowMod"`
}

type PerConTableSet struct {
	PerCon      []PerCon      `json:"PerCon"`
	PerConAttch []interface{} `json:"PerConAttch"`
	PerConLnk   []struct {
		Company             string `json:"Company"`
		PerConID            int    `json:"PerConID"`
		ContextLink         string `json:"ContextLink"`
		LinkSysRowID        string `json:"LinkSysRowID"`
		PrimaryContext      bool   `json:"PrimaryContext"`
		FaxNum              string `json:"FaxNum"`
		PhoneNum            string `json:"PhoneNum"`
		EMailAddress        string `json:"EMailAddress"`
		CellPhoneNum        string `json:"CellPhoneNum"`
		PagerNum            string `json:"PagerNum"`
		HomeNum             string `json:"HomeNum"`
		AltNum              string `json:"AltNum"`
		WebSite             string `json:"WebSite"`
		IM                  string `json:"IM"`
		Twitter             string `json:"Twitter"`
		LinkedIn            string `json:"LinkedIn"`
		FaceBook            string `json:"FaceBook"`
		WebLink1            string `json:"WebLink1"`
		WebLink2            string `json:"WebLink2"`
		WebLink3            string `json:"WebLink3"`
		WebLink4            string `json:"WebLink4"`
		WebLink5            string `json:"WebLink5"`
		Address1            string `json:"Address1"`
		Address2            string `json:"Address2"`
		Address3            string `json:"Address3"`
		City                string `json:"City"`
		State               string `json:"State"`
		Zip                 string `json:"Zip"`
		Country             string `json:"Country"`
		CountryNum          int    `json:"CountryNum"`
		CorpName            string `json:"CorpName"`
		SysRevID            int    `json:"SysRevID"`
		SysRowID            string `json:"SysRowID"`
		CustID              string `json:"CustID"`
		CustName            string `json:"CustName"`
		ShipToNum           string `json:"ShipToNum"`
		ShipToName          string `json:"ShipToName"`
		CustContactName     string `json:"CustContactName"`
		VendorID            string `json:"VendorID"`
		VendorName          string `json:"VendorName"`
		PurPoint            string `json:"PurPoint"`
		PurPointName        string `json:"PurPointName"`
		VendorContactName   string `json:"VendorContactName"`
		SalesRepCode        string `json:"SalesRepCode"`
		SalesRepName        string `json:"SalesRepName"`
		BuyerID             string `json:"BuyerID"`
		BuyerName           string `json:"BuyerName"`
		EmpID               string `json:"EmpID"`
		EmpName             string `json:"EmpName"`
		PREmpID             string `json:"PREmpID"`
		PREmpName           string `json:"PREmpName"`
		SyncNameToPerCon    bool   `json:"SyncNameToPerCon"`
		SyncAddressToPerCon bool   `json:"SyncAddressToPerCon"`
		SyncPhoneToPerCon   bool   `json:"SyncPhoneToPerCon"`
		SyncEmailToPerCon   bool   `json:"SyncEmailToPerCon"`
		SyncLinksToPerCon   bool   `json:"SyncLinksToPerCon"`
		BitFlag             int    `json:"BitFlag"`
		RowMod              string `json:"RowMod"`
	} `json:"PerConLnk"`
	ExtensionTables []interface{} `json:"ExtensionTables"`
}

type PerConResponse struct {
	PerConTableSet PerConTableSet `json:"returnObj"`
}

// GetCaseDetails gets the case details from the help desk
func (h *HelpDeskDetail) GetCaseDetails(caseNum uint) error {
	// Get the necessary ENV variables
	apiKey := os.Getenv("EPICOR_API_KEY")
	authorization := os.Getenv("BASIC_AUTH")
	baseUrl := os.Getenv("BASE_EPICOR_URL")

	baseURL := fmt.Sprintf("%s/Erp.BO.HelpDeskSvc/GetByID?hdCaseNum=%d", baseUrl, caseNum)

	jsonResponse := []byte(`{}`)
	body := bytes.NewBuffer(jsonResponse)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", baseURL, body)

	// Headers
	req.Header.Add("X-API-Key", apiKey)
	req.Header.Add("Authorization", authorization)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

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

	var result CaseDetail

	err = json.Unmarshal(respBody, &result)
	if err != nil {
		log.Println(err)
	}

	h.Details = result

	err = h.GetCaseContacts()
	if err != nil {
		return err
	}

	return nil
}

// GetCaseContacts gets the contacts for the case
func (h *HelpDeskDetail) GetCaseContacts() error {
	// Retrieve the CaseOwner
	caseOwner, err := h.GetSalesRepByID(h.Details.ReturnObj.HDCase[0].CaseOwner)
	if err != nil {
		log.Println(err)
	}
	caseOwner.Type = CaseOwner

	// Retrieve the Primary Contact
	primaryContact, err := h.GetPerConByID(int32(h.Details.ReturnObj.HDCase[0].ReqPerConID))
	if err != nil {
		log.Println(err)
	}
	primaryContact.Type = Primary

	// Get Case Owner Contact
	h.CaseOwnerContact = caseOwner
	// Get Internal Contact
	h.InternalContact = Contact{Type: Internal, Email: "Implement Email Retrieval", Name: "Field not in Live yet"}
	// Get Primary Contact
	h.PrimaryContact = primaryContact

	return nil
}

// GetSalesRepByID gets a given sales rep given a sales rep ID
func (h *HelpDeskDetail) GetSalesRepByID(salesRepID string) (Contact, error) {
	// Get the necessary ENV variables
	apiKey := os.Getenv("EPICOR_API_KEY")
	authorization := os.Getenv("BASIC_AUTH")
	baseUrl := os.Getenv("BASE_EPICOR_URL")

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/Erp.BO.SalesRepSvc/GetByID?salesRepCode=%s", baseUrl, salesRepID), nil)
	// Headers
	req.Header.Add("X-API-Key", apiKey)
	req.Header.Add("Authorization", authorization)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

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

	var result SalesRepResponse
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		log.Println(err)
	}

	// Create a contact for the sales rep and return it
	return Contact{Type: "", Email: result.SalesRepTableset.SalesRep[0].EMailAddress, Name: result.SalesRepTableset.SalesRep[0].Name}, nil

}

// GetPerConByID gets a given contact given a perconID
func (h *HelpDeskDetail) GetPerConByID(perConID int32) (Contact, error) {

	// Get the necessary ENV variables
	apiKey := os.Getenv("EPICOR_API_KEY")
	authorization := os.Getenv("BASIC_AUTH")
	baseUrl := os.Getenv("BASE_EPICOR_URL")

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/Erp.BO.PerConSvc/GetByID?perConID=%d", baseUrl, perConID), nil)
	// Headers
	req.Header.Add("X-API-Key", apiKey)
	req.Header.Add("Authorization", authorization)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

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

	var result PerConResponse
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		log.Println(err)
	}

	// Create a contact for the percon and return it
	return Contact{Type: "", Email: result.PerConTableSet.PerCon[0].EMailAddress, Name: result.PerConTableSet.PerCon[0].Name}, nil

}

// AddComment adds a comment to the case
func (h *HelpDeskDetail) AddComment(caseNum int32, comment string) error {
	// Get the necessary ENV variables
	apiKey := os.Getenv("EPICOR_API_KEY")
	authorization := os.Getenv("BASIC_AUTH")
	baseUrl := os.Getenv("BASE_EPICOR_URL")
	userId := os.Getenv("EPICOR_USER_ID")

	// Get our next comment sequence number
	nextCommentSeq := GetNextCommentSeq(int(caseNum))

	// Create a new Comment
	newComment := Comment{
		CommentSequence:    nextCommentSeq,
		CaseNum:            strconv.Itoa(int(caseNum)),
		Comment:            comment,
		CommentDate:        time.Now().Format("2006-01-02 15:04:05"),
		CommentSequenceKey: strconv.Itoa(nextCommentSeq),
		Company:            "100",
		UserId:             userId,
	}

	// convert comment to json
	commentJson, err := json.Marshal(newComment)
	if err != nil {
		log.Println(err)
	}

	// Create the request body with the comment
	jsonBody := []byte(`{"rollbackParentOnChildError": true,"ds": {"UD01": [` + string(commentJson) + `]},"continueProcessingOnError": true}`)
	body := bytes.NewBuffer(jsonBody)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("POST", baseUrl+"Ice.BO.UD01Svc/UpdateExt", body)

	// Headers
	req.Header.Add("X-API-Key", apiKey)
	req.Header.Add("Authorization", authorization)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Failure Creating UD01: ", err)
	}

	// Read Response Body
	respBody, _ := io.ReadAll(resp.Body)
	log.Println("Response from adding comment:", string(respBody))
	var result UD01TableSetResponse
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		log.Println(err)
	}

	log.Println("Added Comment to Case: ", caseNum)

	return nil
}

type CommentSequenceResponse struct {
	OdataContext string `json:"@odata.context"`
	Value        []struct {
		Number01 float64 `json:"Number01"`
	} `json:"value"`
}

// GetNextCommentSeq gets the next comment sequence number for a given case
func GetNextCommentSeq(caseNum int) int {
	// Get the necessary ENV variables
	apiKey := os.Getenv("EPICOR_API_KEY")
	authorization := os.Getenv("BASIC_AUTH")
	baseUrl := os.Getenv("BASE_EPICOR_URL")

	// Create client
	client := &http.Client{}

	// Create our url
	url := baseUrl + "Ice.BO.UD01Svc/UD01s?$select=Number01&$filter=Key1%20eq%20%27" + strconv.Itoa(caseNum) + "%27&$orderby=Number01%20desc&$top=1"
	// Create request
	req, err := http.NewRequest("GET", url, nil)
	// Headers
	req.Header.Add("X-API-Key", apiKey)
	req.Header.Add("Authorization", authorization)

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

	// Unmarhsal the response into the comment sequence response
	var result CommentSequenceResponse
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		log.Println(err)
	}
	log.Println("New Comment Sequence: ", result.Value[0].Number01+1)
	// Return the next comment sequence number
	return int(result.Value[0].Number01) + 1
}
