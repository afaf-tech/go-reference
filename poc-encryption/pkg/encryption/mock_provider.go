package encryption

import (
	"crypto/sha256"
	"time"
)

type MockKeyProvider struct{}

func NewMockKeyProvider() *MockKeyProvider {
	return &MockKeyProvider{}
}

func (m *MockKeyProvider) FetchKey() (*Key, error) {
	// Mock key for the example, replace with actual key-fetching logic
	key := make([]byte, 32)

	// Derive a 32-byte key from the string using SHA-256
	hash := sha256.Sum256([]byte(key))

	return &Key{
		Value:     hash[:], // Slice the full 32 bytes
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	}, nil
}
