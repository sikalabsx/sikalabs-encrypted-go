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

func GetConfigSikaLabsEncryptedBucket1() (S3Config, error) {
	accessKey, err := decrypt.Decrypt(SIKALABS_ENCRYPTED_BUCKET_1_ACCESS_KEY)
	if err != nil {
		return S3Config{}, err
	}

	secretKey, err := decrypt.Decrypt(SIKALABS_ENCRYPTED_BUCKET_1_SECRET_KEY)
	if err != nil {
		return S3Config{}, err
	}

	return S3Config{
		BucketName: "sikalabs-encrypted-bucket-1",
		Region:     "eu-central-1",
		AccessKey:  accessKey,
		SecretKey:  secretKey,
	}, nil
}
