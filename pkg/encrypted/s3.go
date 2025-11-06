package encrypted

import (
	"github.com/sikalabsx/sikalabs-encrypted-go/internal/decrypt"
)

type S3Config struct {
	BucketName string
	Region     string
	AccessKey  string
	SecretKey  string
}

func GetConfigSikaLabsEncryptedBucket1() (S3Config, error) {
	return getConfigSikaLabsEncryptedBucket(
		"eu-central-1",
		SIKALABS_ENCRYPTED_BUCKET_1_NAME,
		SIKALABS_ENCRYPTED_BUCKET_1_ACCESS_KEY,
		SIKALABS_ENCRYPTED_BUCKET_1_SECRET_KEY,
	)
}

func GetConfigSikaLabsEncryptedBucket2() (S3Config, error) {
	return getConfigSikaLabsEncryptedBucket(
		"eu-central-1",
		SIKALABS_ENCRYPTED_BUCKET_2_NAME,
		SIKALABS_ENCRYPTED_BUCKET_2_ACCESS_KEY,
		SIKALABS_ENCRYPTED_BUCKET_2_SECRET_KEY,
	)
}

func getConfigSikaLabsEncryptedBucket(
	region,
	encryptedBucketName,
	encryptedAccessKey,
	encryptedSecretKey string,
) (S3Config, error) {
	bucketName, err := decrypt.Decrypt(encryptedBucketName)
	if err != nil {
		return S3Config{}, err
	}

	accessKey, err := decrypt.Decrypt(encryptedAccessKey)
	if err != nil {
		return S3Config{}, err
	}

	secretKey, err := decrypt.Decrypt(encryptedSecretKey)
	if err != nil {
		return S3Config{}, err
	}

	return S3Config{
		Region:     region,
		BucketName: bucketName,
		AccessKey:  accessKey,
		SecretKey:  secretKey,
	}, nil
}
