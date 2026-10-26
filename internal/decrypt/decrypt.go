package decrypt

import (
	"github.com/sikalabs/sikalabs-crypt-go/pkg/sikalabs_crypt"
	"github.com/sikalabsx/sikalabs-encrypted-go/internal/password"
)

func DecryptWithSuffix(encrypted string, suffix string) (string, error) {
	password, err := password.GetPasswordWithSuffix(suffix)
	if err != nil {
		return "", err
	}

	decrypted, err := sikalabs_crypt.SikaLabsSymmetricDecryptV1(
		password,
		encrypted,
	)
	if err != nil {
		return "", err
	}

	return decrypted, nil
}

func Decrypt(encrypted string) (string, error) {
	return DecryptWithSuffix(encrypted, "")
}
