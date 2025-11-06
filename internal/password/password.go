package password

import (
	"errors"
	"os"
	"strings"
)

func GetPassword() (string, error) {
	// Try to read password from .SIKALABS_ENCRYPTED_PASSWORD
	if data, err := os.ReadFile(".SIKALABS_ENCRYPTED_PASSWORD"); err == nil {
		return strings.TrimSpace(string(data)), nil
	}
	// Try to read password from ~/.SIKALABS_ENCRYPTED_PASSWORD
	if data, err := os.ReadFile(os.ExpandEnv("$HOME/.SIKALABS_ENCRYPTED_PASSWORD")); err == nil {
		return strings.TrimSpace(string(data)), nil
	}
	// Try to read password from /etc/SIKALABS_ENCRYPTED_PASSWORD
	if data, err := os.ReadFile("/etc/SIKALABS_ENCRYPTED_PASSWORD"); err == nil {
		return strings.TrimSpace(string(data)), nil
	}
	// Fall back to environment variable
	if os.Getenv("SIKALABS_ENCRYPTED_PASSWORD") != "" {
		return os.Getenv("SIKALABS_ENCRYPTED_PASSWORD"), nil
	}

	return "", errors.New("SIKALABS_ENCRYPTED_PASSWORD not found in /etc/SIKALABS_ENCRYPTED_PASSWORD or environment variable")
}
