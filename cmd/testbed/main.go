package main

import (
	"context"
	"log"

	"github.com/flying-robot/gitserver/application"
	"github.com/flying-robot/gitserver/proxy/fs"
	"github.com/flying-robot/gitserver/proxy/git"
)

func main() {
	// Clone a repository to the local machine.
	err := application.HandleCloneRepository(
		context.Background(),
		application.CloneRepository{
			Source:      "https://github.com/flying-robot/commit-sink.git",
			Destination: "/tmp/commit-sink.git",
			MkdirAll:    fs.MkdirAll,
			Init:        git.Init,
			Fetch:       git.Fetch,
		},
	)
	log.Printf("err=%v", err)
}
