package git

// InitArgs configure the behavior of the Init subcommand.
type InitArgs struct {
	Dir string
	Env []string
}

// FetchArgs configure the behavior of the Fetch subcommand.
type FetchArgs struct {
	Upstream string
	Dir      string
	Env      []string
}
