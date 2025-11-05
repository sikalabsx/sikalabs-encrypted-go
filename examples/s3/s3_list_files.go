package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/sikalabs/sikalabs-encrypted-go/pkg/encrypted"
)

func s3ListFiles(cfg encrypted.S3Config) {
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

	p := s3.NewListObjectsV2Paginator(client, &s3.ListObjectsV2Input{
		Bucket: aws.String(cfg.BucketName),
	})

	var keys []string
	for p.HasMorePages() {
		page, err := p.NextPage(ctx)
		if err != nil {
			log.Fatalf("error fetching page: %v", err)
		}
		for _, obj := range page.Contents {
			keys = append(keys, aws.ToString(obj.Key))
		}
	}

	if err != nil {
		log.Fatalf("error listing objects: %v", err)
	}

	for _, k := range keys {
		fmt.Println(k)
	}
}
