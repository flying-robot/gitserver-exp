package service

import (
	"context"
	"os"

	"github.com/flying-robot/gitserver/adapter/writer"
	"github.com/pkg/errors"
)

// A CloneRepositoryService allows clients to clone a repository to the local filesystem.
type CloneRepositoryService struct {
	Filesystem interface {
		MkdirAll(ctx context.Context, args MkdirAllArgs) error
	}
	Git interface {
		Init(ctx context.Context, args InitArgs) error
		Fetch(ctx context.Context, args FetchArgs) error
	}
}

// Clone materializes or reinitializes a repository on disk.
func (c *CloneRepositoryService) Clone(ctx context.Context, req CloneRequest) error {
	var (
		baseArgs     = BaseArgs{Dir: req.Local, Stdout: writer.FlowrateWriter}
		mkdirAllArgs = MkdirAllArgs{Path: req.Local, Mode: os.ModePerm}
		initArgs     = InitArgs{BaseArgs: baseArgs}
		fetchArgs    = FetchArgs{BaseArgs: baseArgs, Upstream: req.Upstream}
	)

	// First we need to set up the location for the repository, along with any
	// intermediate directories on the way to that destination.
	if err := c.Filesystem.MkdirAll(ctx, mkdirAllArgs); err != nil {
		return errors.Wrap(err, "filesystem.mkdirall")
	}

	// We can now initialize the repository in the given directory. The repository
	// will be configured as "bare", without a working directory of its own.
	if err := c.Git.Init(ctx, initArgs); err != nil {
		return errors.Wrap(err, "git.init")
	}

	// Now we can retrieve objects and refs from the upstream into the repository.
	if err := c.Git.Fetch(ctx, fetchArgs); err != nil {
		return errors.Wrap(err, "git.fetch")
	}

	return nil
}
