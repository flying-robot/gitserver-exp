package main

import (
	"context"
	"fmt"
	"log"

	"github.com/flying-robot/gitserver/adapter/filesystem"
	"github.com/flying-robot/gitserver/adapter/git"
	"github.com/flying-robot/gitserver/service"
)

func main() {
	log.Println("cloning with default adapters")
	{
		fmt.Println("=================================================")
		cloneRepoService := &service.CloneRepositoryService{
			Filesystem: &filesystem.FilesystemAdapter{},
			Git:        &git.GitAdapter{},
		}
		err := cloneRepoService.Clone(context.Background(), service.CloneRequest{
			Upstream: "https://github.com/flying-robot/commit-sink.git",
			Local:    "/tmp/commit-sink.git",
		})
		fmt.Println(err)
		fmt.Println()
	}

	log.Println("cloning with tracing adapters")
	{
		fmt.Println("=================================================")
		cloneRepoService := &service.CloneRepositoryService{
			Filesystem: &filesystem.FilesystemAdapter{},
			Git: &git.TracingGitAdapter{
				Trace:            true,
				TracePackAccess:  true,
				TracePacket:      true,
				TracePerformance: true,
				TraceSetup:       true,
			},
		}
		err := cloneRepoService.Clone(context.Background(), service.CloneRequest{
			Upstream: "https://github.com/flying-robot/commit-sink.git",
			Local:    "/tmp/commit-sink.git",
		})
		fmt.Println(err)
		fmt.Println()
	}
}
