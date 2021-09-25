package service

import (
	"context"
)

// A GitCommand is used to invoke a Git operation.
type GitCommand func(ctx context.Context, dir string, env []string, args ...string) ([]byte, error)

// A FilesystemCommand is used to invoke disk operations.
type FilesystemCommand func(ctx context.Context, args ...string) ([]byte, error)
