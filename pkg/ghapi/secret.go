package ghapi

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/jranson/example-set-repo-action-secrets/pkg/repokey"
	"github.com/jranson/example-set-repo-action-secrets/pkg/secret"
)

func SetRepositoryActionSecret(token, orgName, repoName, secretName, val string,
	pk *repokey.PublicKey) error {

	b, err := secret.Encrypt(val, pk)
	if err != nil {
		return err
	}
	es := secret.EncryptedSecret{
		KeyID: pk.KeyID,
		Value: base64.StdEncoding.EncodeToString(b),
	}
	b, err = json.Marshal(es)
	if err != nil {
		return err
	}
	_, _, err = Put(token, fmt.Sprintf(
		"/repos/%s/%s/actions/secrets/%s", orgName, repoName, secretName), b)
	if err != nil {
		return err
	}
	return nil
}
