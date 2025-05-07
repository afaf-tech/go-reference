package connector

import (
	"errors"
	"testing"
)

func TestIsRetriableError(t *testing.T) {
	tests := []struct {
		err      error
		expected bool
	}{
		// Test network-related errors
		{errors.New("connect: connection refused"), true},
		{errors.New("i/o timeout"), true},
		{errors.New("connection reset by peer"), true},
		{errors.New(`Exception (501) Reason: "EOF"`), true},

		// Test non-retriable errors
		{errors.New("database error: syntax error"), false},
		{errors.New("file not found"), false},
		{errors.New("operation not allowed"), false},
	}

	for _, tt := range tests {
		t.Run(tt.err.Error(), func(t *testing.T) {
			got := IsRetriableError(tt.err)
			if got != tt.expected {
				t.Errorf("IsRetriableError(%v) = %v; want %v", tt.err, got, tt.expected)
			}
		})
	}
}
