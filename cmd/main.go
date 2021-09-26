package main

import (
	"context"
	"fmt"

	"github.com/flying-robot/gitserver/adapter/filesystem"
	"github.com/flying-robot/gitserver/adapter/git"
	"github.com/flying-robot/gitserver/service"
)

func main() {
	cloneRepoService := &service.CloneRepositoryService{
		Filesystem: &filesystem.FilesystemAdapter{},
		Git:        &git.GitAdapter{},
	}
	out, err := cloneRepoService.Clone(context.Background(), service.CloneRequest{
		Upstream: "https://github.com/flying-robot/commit-sink.git",
		Local:    "/tmp/commit-sink.git",
	})
	fmt.Println(string(out))
	fmt.Println(err)
}
