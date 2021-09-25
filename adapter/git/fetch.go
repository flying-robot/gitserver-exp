package git

import (
	"context"
	"os/exec"
)

// Fetch downloads objects and refs from another repository.
func Fetch(ctx context.Context, args ...string) ([]byte, error) {
	cmd := exec.CommandContext(ctx, "git", "fetch", "--progress", "--prune", args[0])
	cmd.Dir = args[1]
	return cmd.CombinedOutput()
}
