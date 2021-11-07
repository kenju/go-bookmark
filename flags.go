package main

import (
	"flag"
	"log"
	"os"
	"path"
)

type Flags struct {
	ConfPath string
	Tags     string
}

func NewFlags() *Flags {
	flags := &Flags{}
	flag.StringVar(&flags.ConfPath, "conf-path",
		".config/go-bookmark/bookmark.yaml", "a file path to the config")
	flag.StringVar(&flags.Tags, "tags",
		"", "comma-separated tags to filter output")
	flag.Parse()
	return flags
}

func (f *Flags) confFilePath() string {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return path.Join(userHomeDir, f.ConfPath)
}
