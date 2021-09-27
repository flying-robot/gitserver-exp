package main

import (
	"context"
	"log"

	"github.com/flying-robot/gitserver/adapter"
	"github.com/flying-robot/gitserver/service"
)

func main() {
	request := service.CloneRequest{
		Upstream:        "https://github.com/flying-robot/commit-sink.git",
		Local:           "/tmp/commit-sink.git",
		FlowRateLimiter: adapter.FlowrateWriter,
	}

	log.Println("cloning with default adapters")
	cloneRepoService := &service.CloneRepositoryService{
		Filesystem: &adapter.Filesystem{},
		Git:        &adapter.Git{},
	}
	cloneRepoService.Clone(context.Background(), request)

	log.Println("cloning with tracing adapters")
	cloneRepoService.Git = &adapter.TracingGit{
		Trace:            true,
		TracePackAccess:  true,
		TracePacket:      true,
		TracePerformance: true,
		TraceSetup:       true,
	}
	cloneRepoService.Clone(context.Background(), request)
}
