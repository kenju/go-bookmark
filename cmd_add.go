package main

import (
	"bufio"
	"context"
	"flag"
	"io"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/google/subcommands"
)

type addCmd struct {
	config *Config
	flags  *Flags
}

func (*addCmd) Name() string     { return "add" }
func (*addCmd) Synopsis() string { return "add new command" }
func (*addCmd) Usage() string {
	return `add`
}
func (c *addCmd) SetFlags(f *flag.FlagSet) {
}
func (c *addCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	addNewBookmark(os.Stdin, os.Stdout, c.config, c.flags)
	return subcommands.ExitSuccess
}

func addNewBookmark(in io.Reader, out io.Writer, config *Config, flags *Flags) {
	scanner := bufio.NewScanner(in)
	// bookmark
	io.WriteString(out, "URL to Bookmark: ")
	scanner.Scan()
	newUrlStr := scanner.Text()
	if newUrlStr == "" {
		log.Fatalf("Invalid URL\n")
	}
	// title
	io.WriteString(out, "Title: ")
	scanner.Scan()
	title := scanner.Text()
	if title == "" {
		log.Fatalf("Invalid Title\n")
	}
	// tags
	io.WriteString(out, "Tags (comma-separated): ")
	scanner.Scan()
	tagsStr := scanner.Text()
	tags := strings.Split(tagsStr, ",")
	// parse url
	newUrl, err := url.Parse(newUrlStr)
	if err != nil {
		log.Fatalf("Invalid URL: %s\n", newUrlStr)
	}
	if newUrl.Scheme == "" {
		log.Fatalln("URL Schema is empty")
	}
	// check uniqueness
	for _, bookmark := range config.Bookmarks {
		if bookmark.Url == newUrlStr {
			log.Fatalf("Already bookmarked\n")
		}
	}
	// write to file
	newBookmark := &Bookmark{
		Url:   newUrlStr,
		Title: title,
		Tags:  tags,
	}
	config.Bookmarks = append(config.Bookmarks, newBookmark)
	config.writeBookmark(flags)
}
