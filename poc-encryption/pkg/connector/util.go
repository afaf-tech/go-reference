package connector

import "strings"

// ?: separate from db and rabbit mq
// Function to check if the error is retriable
func IsRetriableError(err error) bool {
	// Check if the error is a network-related error or a temporary database issue
	if strings.Contains(err.Error(), "connect: connection refused") ||
		strings.Contains(err.Error(), "i/o timeout") ||
		strings.Contains(err.Error(), "connection reset by peer") || strings.Contains(err.Error(), `Exception (501) Reason: "EOF"`) {
		return true
	}
	return false
}
