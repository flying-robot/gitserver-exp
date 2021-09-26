package git

import (
	"context"
	"os/exec"

	"github.com/flying-robot/gitserver/service"
)

// A GitAdapter allows gitserver to execute Git subcommands that modify either
// local or remote repositories.
type GitAdapter struct{}

// Init creates an empty Git repository or reinitializes an existing one.
func (g *GitAdapter) Init(ctx context.Context, args service.InitArgs) ([]byte, error) {
	cmd := exec.CommandContext(ctx, "git", "init", "--bare", ".")
	cmd.Dir = args.Dir
	cmd.Env = args.Env
	if args.Stdout != nil {
		cmd.Stdout = args.Stdout(cmd.Stdout)
	}
	return cmd.CombinedOutput()
}

// Fetch downloads objects and refs from another repository.
func (g *GitAdapter) Fetch(ctx context.Context, args service.FetchArgs) ([]byte, error) {
	cmd := exec.CommandContext(ctx, "git", "fetch", "--progress", "--prune", args.Upstream)
	cmd.Dir = args.Dir
	cmd.Env = args.Env
	if args.Stdout != nil {
		cmd.Stdout = args.Stdout(cmd.Stdout)
	}
	return cmd.CombinedOutput()
}
