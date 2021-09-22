package git

import (
	"context"
	"os/exec"
)

// Init creates an empty Git repository or reinitializes an existing one.
func Init(ctx context.Context, args ...string) ([]byte, error) {
	cmd := exec.CommandContext(ctx, "git", "init", "--bare", ".")
	cmd.Dir = args[0]
	return cmd.CombinedOutput()
}
