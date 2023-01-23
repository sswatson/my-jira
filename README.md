
# My Jira

This repo provides starter code for you to create your own command-line application for interacting with your Jira account via the Jira API. The idea is that you clone this repo and modify the Go files to suit your needs. This is a very simple project; it's basically just saving you the time it takes to figure out how get a setup working with the [Go Jira](https://github.com/andygrunwald/go-jira) library.


## Setup instructions

1. Clone this repo: `git clone https://github.com/sswatson/my-jira.git`
2. Install [Go](https://golang.org/doc/install) on your machine (if you haven't already).
3. Go into your Jira account and create an API key: https://id.atlassian.com/manage-profile/security -> "Create and Manage API Tokens".
4. Add the following line to your shell init file:

    ```bash
    export JIRA_API_TOKEN="your-api-token-from-step-3"
    ```

5. Run `./make` (or `go build -o bin/my-jira my-jira/*`, if you have some problem with the `make` script) to build an executable called `my-jira`.
6. Run `./bin/my-jira whoami` and copy the resulting account ID value.
7. Add the following line to your shell init file:

    ```bash
    export JIRA_ACCOUNT_ID="your-account-id-from-step-6"
    ```

8. Run `./bin/my-jira` to see the list of available commands.

## Installation

You can move the `my-jira` executable to a directory in your `$PATH` to make it available system-wide.

## Files

The source code is in `./my-jira/`.

The `go.mod` and `go.sum` files are used by Go to manage dependencies.