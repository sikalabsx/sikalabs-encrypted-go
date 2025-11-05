package encrypted

import (
	"github.com/sikalabs/sikalabs-encrypted-go/internal/decrypt"
)

type S3Config struct {
	BucketName string
	Region     string
	AccessKey  string
	SecretKey  string
}

func GetConfigSikaLabsXEncryptedBucket1() S3Config {
	return S3Config{
		BucketName: "sikalabs-encrypted-bucket-1",
		Region:     "eu-central-1",
		AccessKey:  decrypt.DecryptOrDie(ENCRYPTED_S3_ACCESS_KEY),
		SecretKey:  decrypt.DecryptOrDie(ENCRYPTED_S3_SECRET_KEY),
	}
}
