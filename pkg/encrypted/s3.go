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

func GetConfigSikaLabsEncryptedBucket1() S3Config {
	return S3Config{
		BucketName: "sikalabs-encrypted-bucket-1",
		Region:     "eu-central-1",
		AccessKey:  decrypt.DecryptOrDie(SIKALABS_ENCRYPTED_BUCKET_1_ACCESS_KEY),
		SecretKey:  decrypt.DecryptOrDie(SIKALABS_ENCRYPTED_BUCKET_1_SECRET_KEY),
	}
}
