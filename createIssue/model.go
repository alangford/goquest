package createIssue

type Issue struct {
	Name string `json:"name"`
	Description string `json:"description"`
	IssueType string `json:"issuetype"`
}