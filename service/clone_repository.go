package service

import (
	"context"

	"github.com/cockroachdb/errors"
)

// CloneRepository configures a cloning operation and its dependencies.
type CloneRepository struct {
	Source   string
	WorkDir  string
	Env      []string
	MkdirAll FilesystemCommand
	Init     GitCommand
	Fetch    GitCommand
}

// HandleCloneRepository clones a repository to the local machine.
func HandleCloneRepository(ctx context.Context, cmd CloneRepository) error {
	if _, err := cmd.MkdirAll(ctx, cmd.WorkDir); err != nil {
		return errors.Wrap(err, "mkdirall")
	}
	if _, err := cmd.Init(ctx, cmd.WorkDir, cmd.Env); err != nil {
		return errors.Wrap(err, "init")
	}
	if _, err := cmd.Fetch(ctx, cmd.WorkDir, cmd.Env, cmd.Source); err != nil {
		return errors.Wrap(err, "fetch")
	}
	return nil
}
