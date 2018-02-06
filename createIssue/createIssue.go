package createIssue

import (
	// "github.com/andygrunwald/go-jira"
	"encoding/json"
	"fmt"
	"net/http"
)


func Create() func(http.ResponseWriter, *http.Request) {
	fmt.Println("attempting to create issuer")
	return func (w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var issue Issue
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&issue)
		if err != nil {
			panic(err)
	}
	fmt.Println(issue.Name)
	}
}