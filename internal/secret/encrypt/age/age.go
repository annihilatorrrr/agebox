package age

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"filippo.io/age"

	internalerrors "github.com/slok/agebox/internal/errors"
	keyage "github.com/slok/agebox/internal/key/age"
	"github.com/slok/agebox/internal/model"
	"github.com/slok/agebox/internal/secret/encrypt"
)

type encrypter bool

// Encrypter is the secret.Encrypter implementation for age encrypt/decrypt algorithm.
// This encrypter is coupled to age keys, so it should be used with age based keys.
const Encrypter = encrypter(true)

var _ encrypt.Encrypter = Encrypter

func (encrypter) Encrypt(ctx context.Context, secret model.Secret, keys []model.PublicKey) (*model.Secret, error) {
	if secret.DecryptedData == nil {
		return nil, internalerrors.ErrNotDecrypted
	}

	ageRecipients := make([]age.Recipient, 0, len(keys))
	for _, k := range keys {
		// Get age compatible key.
		var aKey keyage.PublicKey
		switch v := k.(type) {
		case *keyage.PublicKey:
			aKey = *v
		case keyage.PublicKey:
			aKey = v
		default:
			return nil, fmt.Errorf("invalid public key: %w", internalerrors.ErrNotAgeKey)
		}
		ageRecipients = append(ageRecipients, aKey.AgeRecipient())
	}

	// Encrypt data.
	var b bytes.Buffer
	cryptedW, err := age.Encrypt(&b, ageRecipients...)
	if err != nil {
		return nil, fmt.Errorf("age could not prepare secret encrypt: %w", err)
	}

	_, err = cryptedW.Write(secret.DecryptedData)
	if err != nil {
		return nil, fmt.Errorf("could not to encrypt secret: %w", err)
	}
	err = cryptedW.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close encrypted blob: %w", err)
	}

	secret.EncryptedData = b.Bytes()

	return &secret, nil
}
func (encrypter) Decrypt(ctx context.Context, secret model.Secret, keys []model.PrivateKey) (*model.Secret, error) {
	if secret.EncryptedData == nil {
		return nil, internalerrors.ErrNotEncrypted
	}

	ageIdentities := make([]age.Identity, 0, len(keys))
	for _, k := range keys {
		// Get age compatible key.
		var aKey keyage.PrivateKey
		switch v := k.(type) {
		case *keyage.PrivateKey:
			aKey = *v
		case keyage.PrivateKey:
			aKey = v
		default:
			return nil, fmt.Errorf("invalid private key: %w", internalerrors.ErrNotAgeKey)
		}

		ageIdentities = append(ageIdentities, aKey.AgeIdentity())
	}

	// Decrypt data.
	r, err := age.Decrypt(bytes.NewReader(secret.EncryptedData), ageIdentities...)
	if err != nil {
		return nil, fmt.Errorf("age could not decrypt the secret: %w", err)
	}
	secret.DecryptedData, err = io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve decrypted data: %w", err)
	}

	return &secret, nil
}
