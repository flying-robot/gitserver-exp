package application

import (
	"context"
	"errors"
	"testing"
)

func TestCloneRepository(t *testing.T) {
	err := HandleCloneRepository(
		context.Background(),
		CloneRepository{
			Source:      "https://github.com/flying-robot/commit-sink.git",
			Destination: "/tmp/commit-sink.git",
			MkdirAll:    func(ctx context.Context, args ...string) ([]byte, error) { return nil, nil },
			Init:        func(ctx context.Context, args ...string) ([]byte, error) { return nil, nil },
			Fetch:       func(ctx context.Context, args ...string) ([]byte, error) { return nil, nil },
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
			Source:      "https://github.com/flying-robot/commit-sink.git",
			Destination: "/tmp/commit-sink.git",
			MkdirAll:    func(ctx context.Context, args ...string) ([]byte, error) { return nil, errors.New("test") },
			Init:        func(ctx context.Context, args ...string) ([]byte, error) { return nil, nil },
			Fetch:       func(ctx context.Context, args ...string) ([]byte, error) { return nil, nil },
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
			Source:      "https://github.com/flying-robot/commit-sink.git",
			Destination: "/tmp/commit-sink.git",
			MkdirAll:    func(ctx context.Context, args ...string) ([]byte, error) { return nil, nil },
			Init:        func(ctx context.Context, args ...string) ([]byte, error) { return nil, errors.New("test") },
			Fetch:       func(ctx context.Context, args ...string) ([]byte, error) { return nil, nil },
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
			Source:      "https://github.com/flying-robot/commit-sink.git",
			Destination: "/tmp/commit-sink.git",
			MkdirAll:    func(ctx context.Context, args ...string) ([]byte, error) { return nil, nil },
			Init:        func(ctx context.Context, args ...string) ([]byte, error) { return nil, nil },
			Fetch:       func(ctx context.Context, args ...string) ([]byte, error) { return nil, errors.New("test") },
		},
	)
	if err == nil {
		t.Fatal("expected error, got none")
	}
}
