package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/sikalabs/sikalabs-encrypted-go/pkg/encrypted"
)

func main() {
	s3Config, err := encrypted.GetConfigSikaLabsEncryptedBucket1()
	handleError(err)

	printStructAsJSON(s3Config)

	s3createFile(
		s3Config,
		"sikalabs_encrypted_go_test.txt",
		"sikalabs-encrypted-go "+time.Now().Format("2006-01-02 15:04:05"),
	)
	s3ListFiles(s3Config)
	s3PrintFile(s3Config, "sikalabs_encrypted_go_test.txt")
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
