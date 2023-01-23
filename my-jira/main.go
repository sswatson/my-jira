package main

import (
	"github.com/spf13/cobra"
)

func main() {
	var root = &cobra.Command{Use: "jira"}
	addCommands(root)
	root.Execute()
}

func addCommands(root *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "vsx <name> <description>",
		Short: "Create a new VS Code issue",
		Args:  cobra.ExactArgs(2),
		Run:   createVsxIssue,
	}
	root.AddCommand(cmd)
	cmd = &cobra.Command{
		Use:   "whoami",
		Short: "Get your user information",
		Args:  cobra.ExactArgs(0),
		Run:   getUser,
	}
	root.AddCommand(cmd)
	cmd = &cobra.Command{
		Use:   "issue <issue-number>",
		Short: "Get issue information",
		Args:  cobra.ExactArgs(1),
		Run:   getIssue,
	}
	root.AddCommand(cmd)
}
