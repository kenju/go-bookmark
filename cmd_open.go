package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/google/subcommands"
)

type openCmd struct {
	config  *Config
	flags   *Flags
	browser string
}

func (*openCmd) Name() string     { return "open" }
func (*openCmd) Synopsis() string { return "open new command" }
func (*openCmd) Usage() string {
	return `open`
}
func (c *openCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.browser, "browser",
		"", "a browser key to overwrite 'default_browser'")
}
func (c *openCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	browser := c.config.selectBrowser(c.browser)
	if browser == nil {
		log.Fatalf("browser is not selected\n")
	}
	urls := buildUrls(c.config, c.flags)
	runBrowser(browser, urls)
	return subcommands.ExitSuccess
}

func buildUrls(config *Config, flags *Flags) []string {
	selectedLines := selectWithPeco(config, flags)
	bookmarks := config.selectBookmarks(selectedLines)
	if len(bookmarks) == 0 {
		fmt.Println("No bookmark was selected.")
		os.Exit(0)
	}
	var urls []string
	for _, b := range bookmarks {
		urls = append(urls, b.Url)
	}
	return urls
}

func runBrowser(browser *Browser, urls []string) {
	var args []string
	if len(browser.Args) > 0 {
		for _, arg := range browser.Args {
			args = append(args, arg)
		}
	}
	for _, url := range urls {
		args = append(args, url)
	}
	cmd := exec.Command(
		browser.Command,
		args...,
	)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}
