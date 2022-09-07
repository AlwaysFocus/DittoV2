package main

import (
	_case "DittoV2/case"
	"DittoV2/casedetail"
	"DittoV2/tui"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	// Load the .env file in the current directory
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create new HelpDesk instance
	hd := _case.NewHelpDesk()

	// Create new CaseDetail instance
	hdd := casedetail.HelpDeskDetail{}

	// Get all cases
	cases, err := hd.GetAllCases()

	if err != nil {
		log.Fatal("Unable to retrieve cases", err)
	}
	if len(cases) < 1 {

		log.Fatal("no cases found")

	} else {
		// If we didn't get an error, we can start the app
		tui.StartTea(*hd, hdd)
	}
}
