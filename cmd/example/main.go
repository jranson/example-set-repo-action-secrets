package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/jranson/example-set-repo-action-secrets/pkg/runner"
)

// Usage:
// - Set the GH_TOKEN env to a Github Personal Access Token or other credential
// - Create secrets.json file providing secret names as keys and their values.
//   All keys and values must be strings.
// - run: go run cmd/example/main.go <orgName> <repoName>

var gitHubToken = os.Getenv("GH_TOKEN")
var errInvalidSecrets = errors.New("secrets file values must all be strings")

type secretsLookup map[string]string

func main() {
	if gitHubToken == "" {
		panic("GH_TOKEN env not set")
	}

	secrets, err := loadSecrets()
	if err != nil {
		panic(err)
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
		err := runner.Run(gitHubToken, os.Args[1], repo, secrets)
		if err != nil {
			panic(err)
		}
	}
}

func loadSecrets() (secretsLookup, error) {
	b, err := os.ReadFile("secrets.json")
	if err != nil {
		return nil, err
	}

	lkp := make(map[string]any)
	out := make(secretsLookup)
	err = json.Unmarshal(b, &lkp)
	if err != nil {
		return nil, err
	}
	for k, v := range lkp {
		if s, ok := v.(string); ok {
			out[k] = s
		} else {
			return nil, errInvalidSecrets
		}
	}
	return out, nil
}
