package ghapi

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/jranson/example-set-repo-action-secrets/pkg/repokey"
)

func GetRepoPublicKey(token, orgName, repoName string) (*repokey.PublicKey,
	error) {
	code, b, err := Get(token,
		fmt.Sprintf("/repos/%s/%s/actions/secrets/public-key",
			orgName, repoName),
	)
	if err != nil {
		return nil, err
	}
	if code != 200 {
		return nil, repokey.ErrInvalidRepoPublicKey
	}
	var pk repokey.PublicKey
	err = json.Unmarshal(b, &pk)
	if err != nil {
		return nil, err
	}
	if pk.Key == "" || pk.KeyID == "" {
		return nil, repokey.ErrInvalidRepoPublicKey
	}
	b, err = base64.StdEncoding.DecodeString(pk.Key)
	if err != nil {
		return nil, err
	}
	if len(b) != repokey.PublicKeySize {
		return nil, repokey.ErrInvalidRepoPublicKey
	}
	pk.KeyBytes = b
	return &pk, nil
}
