package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
)

type deleteCmd struct {
	config *Config
	flags  *Flags
}

func (*deleteCmd) Name() string     { return "delete" }
func (*deleteCmd) Synopsis() string { return "delete new command" }
func (*deleteCmd) Usage() string {
	return `delete`
}
func (c *deleteCmd) SetFlags(f *flag.FlagSet) {
}
func (c *deleteCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	selected := selectWithPeco(c.config, c.flags)
	var urlsToDelete []string
	for _, selection := range selected {
		if len(selection) == 0 {
			continue
		}
		bookmark := decodeBookmark(selection)
		urlsToDelete = append(urlsToDelete, bookmark.Url)
	}
	var bookmarks []*Bookmark
	for _, bookmark := range c.config.Bookmarks {
		if shouldDelete(bookmark, urlsToDelete) {
			fmt.Printf("Deleting %s (%s)\n", bookmark.Title, bookmark.Url)
		} else {
			bookmarks = append(bookmarks, bookmark)
		}
	}
	c.config.Bookmarks = bookmarks
	c.config.writeBookmark(c.flags)
	return subcommands.ExitSuccess
}

func shouldDelete(bookmark *Bookmark, urlsToDelete []string) bool {
	for _, urlToDelete := range urlsToDelete {
		if urlToDelete == bookmark.Url {
			return true
		}
	}
	return false
}
