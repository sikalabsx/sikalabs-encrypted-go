package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/sikalabsx/sikalabs-encrypted-go/pkg/encrypted"
)

func s3PrintFile(cfg encrypted.S3Config, name string) {
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

	// Get the object from S3
	result, err := client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &cfg.BucketName,
		Key:    &name,
	})
	if err != nil {
		log.Fatalf("error getting file: %v", err)
	}
	defer result.Body.Close()

	// Read and print the content
	content, err := io.ReadAll(result.Body)
	if err != nil {
		log.Fatalf("error reading file content: %v", err)
	}

	fmt.Println(string(content))
}
