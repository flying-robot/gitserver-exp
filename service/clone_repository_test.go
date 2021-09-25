package service

import (
	"context"
	"errors"
	"testing"
)

func TestCloneRepository(t *testing.T) {
	err := HandleCloneRepository(
		context.Background(),
		CloneRepository{
			Source:   "https://github.com/flying-robot/commit-sink.git",
			WorkDir:  "/tmp/commit-sink.git",
			MkdirAll: okFileSystemCommand,
			Init:     okGitCommand,
			Fetch:    okGitCommand,
		},
	)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestCloneRepository_MkdirAllFailed(t *testing.T) {
	err := HandleCloneRepository(
		context.Background(),
		CloneRepository{
			Source:   "https://github.com/flying-robot/commit-sink.git",
			WorkDir:  "/tmp/commit-sink.git",
			Env:      []string{},
			MkdirAll: errFilesystemCommand,
			Init:     okGitCommand,
			Fetch:    okGitCommand,
		},
	)
	if err == nil {
		t.Fatal("expected error, got none")
	}
}

func TestCloneRepository_InitFailed(t *testing.T) {
	err := HandleCloneRepository(
		context.Background(),
		CloneRepository{
			Source:   "https://github.com/flying-robot/commit-sink.git",
			WorkDir:  "/tmp/commit-sink.git",
			Env:      []string{},
			MkdirAll: okFileSystemCommand,
			Init:     errGitCommand,
			Fetch:    okGitCommand,
		},
	)
	if err == nil {
		t.Fatal("expected error, got none")
	}
}

func TestCloneRepository_FetchFailed(t *testing.T) {
	err := HandleCloneRepository(
		context.Background(),
		CloneRepository{
			Source:   "https://github.com/flying-robot/commit-sink.git",
			WorkDir:  "/tmp/commit-sink.git",
			Env:      []string{},
			MkdirAll: okFileSystemCommand,
			Init:     okGitCommand,
			Fetch:    errGitCommand,
		},
	)
	if err == nil {
		t.Fatal("expected error, got none")
	}
}

// These dummy commands are used to force the various handler scenarios to behave
// in different ways.
var (
	okFileSystemCommand = func(_ context.Context, _ ...string) ([]byte, error) {
		return nil, nil
	}
	errFilesystemCommand = func(_ context.Context, _ ...string) ([]byte, error) {
		return nil, errors.New("test")
	}
	okGitCommand = func(_ context.Context, _ string, _ []string, _ ...string) ([]byte, error) {
		return nil, nil
	}
	errGitCommand = func(_ context.Context, _ string, _ []string, _ ...string) ([]byte, error) {
		return nil, errors.New("test")
	}
)
