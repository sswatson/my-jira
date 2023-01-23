package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

func createVsxIssue(cmd *cobra.Command, args []string) {
	name := args[0]
	description := args[1]
	client := newClient()
	fmt.Printf("Creating issue with name \"%s\" and description \"%s\"...\n", name, description)
	createIssue(client, name, description)
}

func getUser(cmd *cobra.Command, args []string) {
	client := newClient()
	user, _, err := client.User.Find(context.Background(), "samuel.watson@relational.ai")
	if err != nil {
		panic(err)
	}
	prettyString, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(prettyString))
}

func getIssue(cmd *cobra.Command, args []string) {
	client := newClient()
	issueId := "RAI-" + args[0]
	issue, _, err := client.Issue.Get(context.Background(), issueId, nil)
	if err != nil {
		panic(err)
	}
	prettyString, err := json.MarshalIndent(issue, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(prettyString))
}
