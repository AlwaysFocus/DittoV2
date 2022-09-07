package casedetail

import (
	"fmt"
)

const divider = "---"

// FormatCaseDescription return the description as a formatted string
func FormatCaseDescription(caseDescription string) string {
	return fmt.Sprintf("%s\n%s\n", caseDescription, divider)
}

// FormatCaseContacts return the contacts as a formatted string
func FormatCaseContacts(caseOwner Contact, casePrimaryContact Contact, caseInternalContact Contact) string {
	return fmt.Sprintf(
		"Case Owner: `%s`\nEmail: %s\n %s\nCase Primary Contact: `%s`\nEmail: %s\n%s\nCase Internal Contact: `%s`\nEmail:%s\n%s\n",
		caseOwner.Name,
		caseOwner.Email,
		divider,
		casePrimaryContact.Name,
		casePrimaryContact.Email,
		divider,
		caseInternalContact.Name,
		caseInternalContact.Email,
		divider,
	)
}
