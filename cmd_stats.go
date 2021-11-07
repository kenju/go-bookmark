package main

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

type statsCmd struct {
	config *Config
}

func (*statsCmd) Name() string     { return "stats" }
func (*statsCmd) Synopsis() string { return "show version" }
func (*statsCmd) Usage() string {
	return `stats`
}
func (c *statsCmd) SetFlags(f *flag.FlagSet) {
}
func (c *statsCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	showStats(c.config)
	return subcommands.ExitSuccess
}

func showStats(config *Config) {
	stats := NewStats(config)
	stats.Show()
}
