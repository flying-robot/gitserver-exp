package application

import (
	"context"

	"github.com/cockroachdb/errors"
)

// CloneRepository configures a cloning operation and its dependencies.
type CloneRepository struct {
	Source      string
	Destination string
	MkdirAll    Command
	Init        Command
	Fetch       Command
}

// HandleCloneRepository clones a repository to the local machine.
func HandleCloneRepository(ctx context.Context, cmd CloneRepository) error {
	if _, err := cmd.MkdirAll(ctx, cmd.Destination); err != nil {
		return errors.Wrap(err, "mkdirall")
	}
	if _, err := cmd.Init(ctx, cmd.Destination); err != nil {
		return errors.Wrap(err, "init")
	}
	if _, err := cmd.Fetch(ctx, cmd.Source, cmd.Destination); err != nil {
		return errors.Wrap(err, "fetch")
	}
	return nil
}
