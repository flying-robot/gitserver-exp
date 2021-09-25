package main

import (
	"context"
	"log"

	"github.com/flying-robot/gitserver/adapter/fs"
	"github.com/flying-robot/gitserver/adapter/git"
	"github.com/flying-robot/gitserver/service"
)

func main() {
	// Clone a repository to the local machine.
	err := service.HandleCloneRepository(
		context.Background(),
		service.CloneRepository{
			Source:   "https://github.com/flying-robot/commit-sink.git",
			WorkDir:  "/tmp/commit-sink.git",
			Env:      []string{},
			MkdirAll: fs.MkdirAll,
			Init:     git.Init,
			Fetch:    git.Fetch,
		},
	)
	log.Printf("err=%v", err)
}
