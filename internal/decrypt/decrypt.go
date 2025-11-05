package decrypt

import (
	"log"

	"github.com/sikalabs/sikalabs-crypt-go/pkg/sikalabs_crypt"
	"github.com/sikalabs/sikalabs-encrypted-go/internal/password"
)

func DecryptOrDie(encrypted string) string {
	decrypted, err := sikalabs_crypt.SikaLabsSymmetricDecryptV1(
		password.GetPasswordOrDie(),
		encrypted,
	)
	if err != nil {
		log.Fatalln("Failed to decrypt:", err)
	}
	return decrypted
}
