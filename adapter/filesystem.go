package adapter

import (
	"context"
	"os"

	"github.com/flying-robot/gitserver/service"
)

// A Filesystem adapter allows gitserver to execute shell commands that access or mutate
// an underlying filesystem.
type Filesystem struct{}

// MkdirAll creates a path on disk, including subdirectories, or returns an error.
func (f *Filesystem) MkdirAll(ctx context.Context, args service.MkdirAllArgs) error {
	return os.MkdirAll(args.Path, args.Mode)
}
