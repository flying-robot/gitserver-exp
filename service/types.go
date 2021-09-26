package service

import (
	"io"
	"os"
)

// BaseArgs are provided to every subcommand and control the operating environment.
type BaseArgs struct {
	Dir    string
	Env    []string
	Stdout func(w io.Writer) io.Writer
}

// InitArgs configure the behavior of the Init subcommand.
type InitArgs struct {
	BaseArgs
}

// FetchArgs configure the behavior of the Fetch subcommand.
type FetchArgs struct {
	BaseArgs
	Upstream string
}

// MkdirAllArgs configure the behavior of the MkdirAll command.
type MkdirAllArgs struct {
	Path string
	Mode os.FileMode
}

type CloneRequest struct {
	Upstream string
	Local    string
}
