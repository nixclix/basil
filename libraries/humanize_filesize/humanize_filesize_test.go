package humanize_filesize

import (
	"testing"
)

func TestHumanizeFilesize(t *testing.T) {
	tests := []struct {
		name          string
		size_in_bytes *int32
		expected      string
	}{
		{
			name:          "nil bytes",
			size_in_bytes: nil,
			expected:      "0 MB",
		},
		{
			name:          "2048 bytes",
			size_in_bytes: int32Ptr(2048),
			expected:      "0.0020 MB",
		},
		{
			name:          "0 bytes",
			size_in_bytes: int32Ptr(0),
			expected:      "0.0000 MB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetHumanizedFilesize(tt.size_in_bytes)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func int32Ptr(n int32) *int32 {
	return &n
}
