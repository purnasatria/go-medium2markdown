package md2

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestGetFunctionName(t *testing.T) {
	functionName := getFunctionName()
	if functionName != "tRunner" {
		t.Errorf("Expected 'tRunner', got '%s'", functionName)
	}
}

func TestNewError(t *testing.T) {
	tests := []struct {
		name   string
		format string
		args   []interface{}
	}{
		{
			name:   "Simple error",
			format: "simple error",
			args:   []interface{}{},
		},
		{
			name:   "Error with formatting",
			format: "error with %s",
			args:   []interface{}{"formatting"},
		},
		{
			name:   "Error with multiple args",
			format: "error with %d and %s",
			args:   []interface{}{42, "string"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := newError(tt.format, tt.args...)
			errString := err.Error()

			parts := strings.SplitN(errString, " - ", 2)
			if len(parts) != 2 {
				t.Errorf("Error message does not contain expected format: %s", errString)
				return
			}

			if parts[0] == "" {
				t.Error("Function name part of the error is empty")
			}

			expectedMessage := fmt.Sprintf(tt.format, tt.args...)
			if parts[1] != expectedMessage {
				t.Errorf("Expected error message '%s', got '%s'", expectedMessage, parts[1])
			}
		})
	}
}

func TestNewErrorWrapping(t *testing.T) {
	originalErr := errors.New("original error")
	wrappedErr := newError("wrapped: %v", originalErr)

	if !strings.Contains(wrappedErr.Error(), originalErr.Error()) {
		t.Errorf("Wrapped error '%v' does not contain original error '%v'", wrappedErr, originalErr)
	}
}
