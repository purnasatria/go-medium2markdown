package md2

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCallMediumAPI(t *testing.T) {
	tests := []struct {
		name           string
		responseBody   string
		expectedResult []byte
		expectError    bool
	}{
		{
			name:           "Valid Medium response",
			responseBody:   mediumJsonPrefix + `{"key": "value"}`,
			expectedResult: []byte(`{"key": "value"}`),
			expectError:    false,
		},
		{
			name:           "Invalid Medium response",
			responseBody:   `{"key": "value"}`,
			expectedResult: nil,
			expectError:    true,
		},
		{
			name:           "Empty response",
			responseBody:   "",
			expectedResult: nil,
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(tt.responseBody))
			}))
			defer server.Close()

			result, err := callMediumAPI(server.URL)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected an error, but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if !bytes.Equal(result, tt.expectedResult) {
					t.Errorf("Expected %s, but got %s", tt.expectedResult, result)
				}
			}
		})
	}
}

func TestDownloadFile(t *testing.T) {
	tests := []struct {
		name           string
		responseBody   string
		expectedResult []byte
		statusCode     int
		expectError    bool
	}{
		{
			name:           "Successful download",
			responseBody:   "file content",
			statusCode:     http.StatusOK,
			expectedResult: []byte("file content"),
			expectError:    false,
		},
		{
			name:           "Empty file",
			responseBody:   "",
			statusCode:     http.StatusOK,
			expectedResult: nil,
			expectError:    true,
		},
		{
			name:           "Bad status code",
			responseBody:   "error content",
			statusCode:     http.StatusNotFound,
			expectedResult: nil,
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
				w.Write([]byte(tt.responseBody))
			}))
			defer server.Close()

			result, err := downloadFile(server.URL)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected an error, but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if !bytes.Equal(result, tt.expectedResult) {
					t.Errorf("Expected %s, but got %s", tt.expectedResult, result)
				}
			}
		})
	}
}
