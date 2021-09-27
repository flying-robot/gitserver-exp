package adapter

import (
	"context"
	"os"
	"os/exec"

	"github.com/flying-robot/gitserver/service"
)

// A Git adapter allows gitserver to execute Git subcommands that modify either
// local or remote repositories.
type Git struct{}

// Init creates an empty Git repository or reinitializes an existing one.
func (g *Git) Init(ctx context.Context, args service.InitArgs) error {
	return g.CmdWithConfig(
		exec.CommandContext(ctx, "git", "init", "--bare", "."),
		args.BaseArgs,
	).Run()
}

// Fetch downloads objects and refs from another repository.
func (g *Git) Fetch(ctx context.Context, args service.FetchArgs) error {
	return g.CmdWithConfig(
		exec.CommandContext(ctx, "git", "fetch", "--progress", "--prune", args.Upstream),
		args.BaseArgs,
	).Run()
}

// CmdWithConfig accepts a base command and configuration arguments. The arguments
// are used to set up the command's operating environment.
func (g *Git) CmdWithConfig(cmd *exec.Cmd, args service.BaseArgs) *exec.Cmd {
	cmd.Dir = args.Dir
	cmd.Env = args.Env

	// We occasionally use flow rate limiters to restrict bandwidth usage. If one
	// is provided, it should wrap STDOUT.
	if args.Stdout != nil {
		cmd.Stdout = args.Stdout(os.Stdout)
	}

	// This makes Run() behave like CombinedOutput(), where STDOUT and STDERR
	// are collapsed into one stream.
	cmd.Stderr = cmd.Stdout
	return cmd
}
