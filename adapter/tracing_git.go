package adapter

import (
	"context"

	"github.com/flying-robot/gitserver/service"
)

// A TracingGit adapter delegates to a normal Git Adapter, but with Git's debugging
// environment variables configured for more verbose output.
type TracingGit struct {
	Trace            bool
	TracePackAccess  bool
	TracePacket      bool
	TracePerformance bool
	TraceSetup       bool

	Git
}

// Init creates an empty Git repository or reinitializes an existing one.
func (t *TracingGit) Init(ctx context.Context, args service.InitArgs) error {
	args.Env = append(args.Env, t.env()...)
	return t.Git.Init(ctx, args)
}

// Fetch downloads objects and refs from another repository.
func (t *TracingGit) Fetch(ctx context.Context, args service.FetchArgs) error {
	args.Env = append(args.Env, t.env()...)
	return t.Git.Fetch(ctx, args)
}

// env returns a slice of environment variables that configure Git to produce
// more verbose debugging output.
func (t *TracingGit) env() []string {
	var env []string
	if t.Trace {
		env = append(env, "GIT_TRACE=true")
	}
	if t.TracePackAccess {
		env = append(env, "GIT_TRACE_PACK_ACCESS=true")
	}
	if t.TracePacket {
		env = append(env, "GIT_TRACE_PACKET=true")
	}
	if t.TracePerformance {
		env = append(env, "GIT_TRACE_PERFORMANCE=true")
	}
	if t.TraceSetup {
		env = append(env, "GIT_TRACE_SETUP=true")
	}
	return env
}
