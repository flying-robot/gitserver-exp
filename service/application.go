package service

import (
	"context"
)

// A Command is generally an invocation of another program on the local machine,
// but can be a mock or fake for testing purposes.
type Command func(ctx context.Context, args ...string) ([]byte, error)
