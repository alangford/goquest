package createIssue

import (
	"os"
	"fmt"
	b64 "encoding/base64"
	"net/http"
	"bytes"
)


func Create(d, t, pro, s string) (error) {
	fmt.Println(d + " " + t + " " + pro + " " + s)
	url := os.Getenv("JIRA_URI")
	u := os.Getenv("JIRA_USER")
	p := os.Getenv("JIRA_PASS")
	fmt.Println(url + " " + u + " " + p)
	var data = []byte(fmt.Sprintf(`{"fields": {"project": { "key": %s}, "summary": "%s", "description": "%s", "issuetype": {"name": "%s"}}}`, pro, s, d, t))

	auth := b64.StdEncoding.EncodeToString([]byte(u + ":" + p))
	req, err := http.NewRequest("POST", url + "rest/api/2/issue/", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Basic " + auth)
	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}