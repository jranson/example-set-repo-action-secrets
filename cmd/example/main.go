package main

import (
	"fmt"
	"os"

	"github.com/jranson/example-set-repo-action-secrets/pkg/runner"
)

// Usage:
// - Set the GH_TOKEN env to a Github Personal Access Token or other credential
// - Update the exampleSecrets map below with the desired name/values
//   (or implement some other way to read in a map of secrets)
// - run: go run cmd/example/main.go <orgName> <repoName>

var gitHubToken = os.Getenv("GH_TOKEN")

func main() {
	if gitHubToken == "" {
		panic("GH_TOKEN env not set")
	}

	usage := func() {
		fmt.Println("Usage:")
		fmt.Println("\n  example-set-repo-action-secrets <orgName> <repoName>")
		fmt.Println("\n\nExample:")
		fmt.Println("\n  example-set-repo-action-secrets example-org example-repo")
		fmt.Println()
		fmt.Println()
	}

	if len(os.Args) < 3 {
		usage()
		return
	}

	// you can pass these in from the command line, or read from a json file, etc.
	exampleSecrets := map[string]string{
		"SECRET_NAME_1": "secret value 1",
		"SECRET_NAME_2": "secret value 2",
	}

	fmt.Println("Working on repo", os.Args[1], "/", os.Args[2])
	err := runner.Run(gitHubToken, os.Args[1], os.Args[2], exampleSecrets)
	if err != nil {
		panic(err)
	}
}
