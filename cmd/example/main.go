package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jranson/example-set-repo-action-secrets/pkg/runner"
)

// Usage:
// - Set the GH_TOKEN env to a Github Personal Access Token or other credential
// - Update the exampleSecrets map below with the desired name/values
//   (or implement some other way to read in a map of secrets)
// - run: go run cmd/example.main.go <orgName> <repoName>

var gitHubToken = os.Getenv("GH_TOKEN")

var exampleSecrets = map[string]string{
	"SECRET_NAME_1": "secret value 1",
	"SECRET_NAME_2": "secret value 2",
}

func main() {
	if gitHubToken == "" {
		panic("GH_TOKEN env not set")
	}

	usage := func() {
		fmt.Println("Usage:")
		fmt.Println("\n  example-set-repo-action-secrets <orgName> <repoName1,repoName2>")
		fmt.Println("\n\nExample:")
		fmt.Println("\n  example-set-repo-action-secrets example-org example-repo")
		fmt.Println()
		fmt.Println()
	}

	if len(os.Args) < 3 {
		usage()
		return
	}

	repos := strings.Split(os.Args[2], ",")
	for _, repo := range repos {
		if repo == "" {
			continue
		}
		fmt.Println("Working on repo", os.Args[1], "/", repo)
		err := runner.Run(gitHubToken, os.Args[1], repo, exampleSecrets)
		if err != nil {
			panic(err)
		}
	}
}
