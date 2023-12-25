package secret

import (
	"crypto/rand"
	"errors"
	"io"

	"github.com/jranson/example-set-repo-action-secrets/pkg/repokey"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/nacl/box"
)

const NonceSize = 24

type EncryptedSecret struct {
	Value string `json:"encrypted_value"`
	KeyID string `json:"key_id"`
}

var ErrInvalidEncrytpedSecret = errors.New("invalid encrytped secret")

func Encrypt(secretValue string, pk *repokey.PublicKey) ([]byte, error) {
	if secretValue == "" || pk == nil || len(pk.KeyBytes) == 0 {
		return nil, nil
	}
	keyCopy := new([repokey.PublicKeySize]byte)
	copy(keyCopy[:], pk.KeyBytes)
	pub, priv, err := box.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	nonce := new([NonceSize]byte)
	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		return nil, err
	}
	nhWriter, err := blake2b.New(NonceSize, nil)
	if err != nil {
		return nil, err
	}
	if _, err := nhWriter.Write(pub[:]); err != nil {
		return nil, err
	}
	if _, err := nhWriter.Write(keyCopy[:]); err != nil {
		return nil, err
	}
	copy(nonce[:], nhWriter.Sum(nil))
	return box.Seal(pub[:], []byte(secretValue), nonce, keyCopy, priv), nil
}
