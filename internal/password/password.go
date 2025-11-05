package password

import (
	"log"
	"os"
	"strings"
)

func GetPasswordOrDie() string {
	pwd := ""
	// Try to read password from .SIKALABS_ENCRYPTED_PASSWORD
	if data, err := os.ReadFile(".SIKALABS_ENCRYPTED_PASSWORD"); err == nil {
		return strings.TrimSpace(string(data))
	}
	// Try to read password from ~/.SIKALABS_ENCRYPTED_PASSWORD
	if data, err := os.ReadFile(os.ExpandEnv("$HOME/.SIKALABS_ENCRYPTED_PASSWORD")); err == nil {
		return strings.TrimSpace(string(data))
	}
	// Try to read password from /etc/SIKALABS_ENCRYPTED_PASSWORD
	if data, err := os.ReadFile("/etc/SIKALABS_ENCRYPTED_PASSWORD"); err == nil {
		return strings.TrimSpace(string(data))
	}
	// Fall back to environment variable
	if pwd == "" {
		pwd = os.Getenv("SIKALABS_ENCRYPTED_PASSWORD")
	}
	// Fatal if password is still empty
	if pwd == "" {
		log.Fatalln("SIKALABS_ENCRYPTED_PASSWORD not found in /etc/SIKALABS_ENCRYPTED_PASSWORD or environment variable")
	}
	return pwd
}
