package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
)

type versionCmd struct {
	config *Config
}

func (*versionCmd) Name() string     { return "version" }
func (*versionCmd) Synopsis() string { return "show bookmark version" }
func (*versionCmd) Usage() string {
	return `version`
}
func (c *versionCmd) SetFlags(f *flag.FlagSet) {
}
func (c *versionCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	showVersion()
	return subcommands.ExitSuccess
}

func showVersion() {
	fmt.Printf("go-bookmark: version=%s\n", appVersion)
}
