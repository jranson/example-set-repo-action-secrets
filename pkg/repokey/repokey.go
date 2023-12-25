package repokey

import (
	"errors"
)

type PublicKey struct {
	KeyID    string `json:"key_id"`
	Key      string `json:"key"`
	KeyBytes []byte `json:"-"`
}

const PublicKeySize = 32

var ErrInvalidRepoPublicKey = errors.New("invalid public key for repo")
