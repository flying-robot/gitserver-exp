package filesystem

import "os"

// MkdirAllArgs configure the behavior of the MkdirAll command.
type MkdirAllArgs struct {
	Path string
	Mode os.FileMode
}
