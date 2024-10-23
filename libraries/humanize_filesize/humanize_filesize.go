package humanize_filesize

import "fmt"

// GetHumanizedFilesize takes size_in_bytes as an int32 pointer and returns the size in megabytes.
func GetHumanizedFilesize(size_in_bytes *int32) string {
	if size_in_bytes != nil {
		size_in_megabytes := float64(*size_in_bytes) / (1024 * 1024)
		return fmt.Sprintf("%.4f MB", size_in_megabytes)
	}
	return "0 MB"
}
