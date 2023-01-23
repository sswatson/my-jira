package main

import (
	"context"
	"fmt"
	"os"

	jira "github.com/andygrunwald/go-jira/v2/cloud"
)

func newClient() *jira.Client {
	jiraURL := "https://relationalai.atlassian.net/"

	// Jira docs: https://support.atlassian.com/atlassian-account/docs/manage-api-tokens-for-your-atlassian-account/
	// Create a new API token: https://id.atlassian.com/manage-profile/security/api-tokens
	tp := jira.BasicAuthTransport{
		Username: "samuel.watson@relational.ai",
		APIToken: os.Getenv("JIRA_API_TOKEN"),
	}

	client, err := jira.NewClient(jiraURL, tp.Client())
	if err != nil {
		panic(err)
	}

	return client
}

func getAccountId() string {
	accountID := os.Getenv("JIRA_ACCOUNT_ID")
	if accountID == "" {
		panic("JIRA_ACCOUNT_ID not set")
	}
	return accountID
}

// functon to create an issue; takes a jira.Client as an argument:
func createIssue(client *jira.Client, summary string, description string) {
	i := jira.Issue{
		Fields: &jira.IssueFields{
			Assignee: &jira.User{
				AccountID: getAccountId(), // you can get this with ./jira whoami
			},
			Description: description,
			Type: jira.IssueType{
				Name: "Task",
			},
			Components: []*jira.Component{
				{
					Name: "VSCode",
				},
			},
			Project: jira.Project{
				Key: "RAI",
			},
			Summary: summary,
		},
	}
	issue, _, err := client.Issue.Create(context.Background(), &i)
	// defer response.Body.Close()
	// body, _ := ioutil.ReadAll(response.Body)
	// fmt.Println(string(body))

	if err != nil {
		panic(err)
	}

	fmt.Println("Created issue", issue.Key, "successfully!")
}
