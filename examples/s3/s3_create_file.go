package main

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/sikalabs/sikalabs-encrypted-go/pkg/encrypted"
)

func s3createFile(cfg encrypted.S3Config, name, content string) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Build an AWS config that uses the provided region + static credentials.
	awscfg, err := awsConfig.LoadDefaultConfig(
		ctx,
		awsConfig.WithRegion(cfg.Region),
		awsConfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(cfg.AccessKey, cfg.SecretKey, ""),
		),
	)
	if err != nil {
		log.Fatalln(err)
	}

	client := s3.NewFromConfig(awscfg)

	// Upload the file content to S3
	_, err = client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(cfg.BucketName),
		Key:    aws.String(name),
		Body:   strings.NewReader(content),
	})
	if err != nil {
		log.Fatalf("error uploading file: %v", err)
	}
}
