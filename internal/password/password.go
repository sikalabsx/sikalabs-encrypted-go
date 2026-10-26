package password

import (
	"errors"
	"os"
	"strings"
)

func GetPasswordWithSuffix(suffix string) (string, error) {
	// Try to read password from .SIKALABS_ENCRYPTED_PASSWORD with suffix
	if data, err := os.ReadFile(".SIKALABS_ENCRYPTED_PASSWORD" + suffix); err == nil {
		return strings.TrimSpace(string(data)), nil
	}
	// Try to read password from ~/.SIKALABS_ENCRYPTED_PASSWORD with suffix
	if data, err := os.ReadFile(os.ExpandEnv("$HOME/.SIKALABS_ENCRYPTED_PASSWORD" + suffix)); err == nil {
		return strings.TrimSpace(string(data)), nil
	}
	// Try to read password from /etc/SIKALABS_ENCRYPTED_PASSWORD with suffix
	if data, err := os.ReadFile("/etc/SIKALABS_ENCRYPTED_PASSWORD" + suffix); err == nil {
		return strings.TrimSpace(string(data)), nil
	}
	// Fall back to environment variable
	if os.Getenv("SIKALABS_ENCRYPTED_PASSWORD"+suffix) != "" {
		return os.Getenv("SIKALABS_ENCRYPTED_PASSWORD" + suffix), nil
	}

	return "", errors.New("SIKALABS_ENCRYPTED_PASSWORD" + suffix + " not found in /etc/SIKALABS_ENCRYPTED_PASSWORD" + suffix + " or environment variable")
}

func GetPassword() (string, error) {
	return GetPasswordWithSuffix("")
}
