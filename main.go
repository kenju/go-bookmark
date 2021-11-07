package main

import (
	"context"
	"log"
	"os"

	"github.com/google/subcommands"
	"gopkg.in/yaml.v2"
)

const (
	appVersion = "v0.2.0"
)

func main() {
	flags := NewFlags()
	config := loadConfig(flags)
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(&versionCmd{config}, "")
	subcommands.Register(&statsCmd{config}, "")
	subcommands.Register(&addCmd{config, flags}, "")
	subcommands.Register(&deleteCmd{config, flags}, "")
	subcommands.Register(&openCmd{config: config, flags: flags}, "")
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}

func loadConfig(flags *Flags) *Config {
	confFilePath := flags.confFilePath()
	confFile, err := os.ReadFile(confFilePath)
	if err != nil {
		log.Fatal(err)
	}
	config := unmarshalConfig(confFile)
	if err := config.validate(); err != nil {
		log.Fatal(err)
	}
	return config
}

func unmarshalConfig(conf []byte) *Config {
	config := &Config{}
	err := yaml.Unmarshal([]byte(conf), &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
