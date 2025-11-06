package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	sikalabs_s3_go "github.com/sikalabs-go/sikalabs-s3-go"
	"github.com/sikalabsx/sikalabs-encrypted-go/pkg/encrypted"
)

func main() {
	s3Config, err := encrypted.GetConfigSikaLabsEncryptedBucket1()
	handleError(err)

	s3ConfigSL := sikalabs_s3_go.S3Config{
		AccessKey:  s3Config.AccessKey,
		SecretKey:  s3Config.SecretKey,
		Region:     s3Config.Region,
		BucketName: s3Config.BucketName,
	}

	printStructAsJSON(s3Config)

	sikalabs_s3_go.PutObject(
		s3ConfigSL,
		"sikalabs_encrypted_go_test.txt",
		[]byte("sikalabs-encrypted-go "+time.Now().Format("2006-01-02 15:04:05")),
	)
	s3ListFiles(s3Config)
	content, err := sikalabs_s3_go.GetObject(
		s3ConfigSL,
		"sikalabs_encrypted_go_test.txt",
	)
	handleError(err)
	fmt.Println("Downloaded content:", string(content))

	s3Config2, err := encrypted.GetConfigSikaLabsEncryptedBucket2()
	handleError(err)

	printStructAsJSON(s3Config2)

	s3ListFiles(s3Config2)

	s3Config3, err := encrypted.GetConfigSikaLabsEncryptedBucket3()
	handleError(err)

	printStructAsJSON(s3Config3)

	s3ListFiles(s3Config3)
}

func printStructAsJSON(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(b))
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
