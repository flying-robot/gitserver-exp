package filesystem

import (
	"context"
	"os"

	"github.com/flying-robot/gitserver/service"
)

// An Adapter allows gitserver to execute shell commands that access or mutate
// an underlying filesystem.
type Adapter struct{}

// MkdirAll creates a path on disk, including subdirectories, or returns an error.
func (a *Adapter) MkdirAll(ctx context.Context, args service.MkdirAllArgs) error {
	return os.MkdirAll(args.Path, args.Mode)
}
