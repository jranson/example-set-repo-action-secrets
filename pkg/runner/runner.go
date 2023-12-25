package runner

import (
	"fmt"

	"github.com/jranson/example-set-repo-action-secrets/pkg/ghapi"
	"github.com/jranson/example-set-repo-action-secrets/pkg/repokey"
)

// Run will apply the map of secrets to the provided org + repo
func Run(token, orgName, repoName string, secrets map[string]string) error {
	if len(secrets) == 0 || token == "" || orgName == "" || repoName == "" {
		fmt.Println("WARNING: invalid inputs", token, orgName, repoName, secrets)
		return nil
	}
	pk, err := ghapi.GetRepoPublicKey(token, orgName, repoName)
	if err != nil {
		return err
	}
	if pk == nil {
		return repokey.ErrInvalidRepoPublicKey
	}
	for name, value := range secrets {
		err = ghapi.SetRepositoryActionSecret(token, orgName, repoName,
			name, value, pk)
		if err != nil {
			return err
		}
	}
	return nil
}
