package main

import (
	"errors"
	"testing"
)

func TestNormalizeUrl(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		expected      string
		expectedError error
	}{
		{
			name:          "remove scheme",
			inputURL:      "https://blog.boot.dev/path",
			expected:      "blog.boot.dev/path",
			expectedError: nil,
		},
		{
			name:          "empty url",
			inputURL:      "",
			expected:      "",
			expectedError: errors.New("url is empty"),
		},
		{
			name:          "invalid url with empty host",
			inputURL:      "http://",
			expected:      "",
			expectedError: errors.New("invalid url"),
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeUrl(tc.inputURL)
			if err != nil {
				if tc.expectedError.Error() == err.Error() {
					return
				}
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
			}
			if actual != tc.expected {
				t.Errorf("Test %v - '%s' FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
