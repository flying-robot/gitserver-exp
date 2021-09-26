package filesystem

import (
	"context"
	"os"

	"github.com/flying-robot/gitserver/service"
)

// A FilesystemAdapter allows gitserver to execute shell commands that access
// or mutate an underlying filesystem.
type FilesystemAdapter struct{}

// MkdirAll creates a path on disk, including subdirectories, or returns an error.
func (f *FilesystemAdapter) MkdirAll(ctx context.Context, args service.MkdirAllArgs) ([]byte, error) {
	return nil, os.MkdirAll(args.Path, args.Mode)
}
