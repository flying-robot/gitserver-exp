package fs

import (
	"context"
	"os"
)

// MkdirAll creates a path on disk, including subdirectories, or returns an error.
func MkdirAll(ctx context.Context, args ...string) ([]byte, error) {
	return nil, os.MkdirAll(args[0], os.ModePerm)
}
