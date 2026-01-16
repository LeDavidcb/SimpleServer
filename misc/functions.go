package misc

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// This function generates a ramdom 256 bytes hash for a refreshToken
func GenerateRamdomToken() (string, error) {
	// Generate a slice of random bytes (e.g., 32 bytes for a good seed)
	randomBytes := make([]byte, 32)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", fmt.Errorf("failed to read random bytes: %w", err)
	}

	// Calculate the SHA256 hash of the random bytes
	hashBytes := sha256.Sum256(randomBytes)

	// Convert the byte slice to a hexadecimal string
	hashString := hex.EncodeToString(hashBytes[:])

	return hashString, nil
}
