package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

const (
	fieldSplit = " | "
)

type Config struct {
	Version        string
	DefaultBrowser string `yaml:"default_browser"`
	Browsers       []*Browser
	Bookmarks      []*Bookmark
	// MultiSources assume that you have multiple bookmark.yaml config
	MultiSources []string `yaml:"multi_sources"`
}

type Browser struct {
	Key     string
	Command string
	Args    []string
}

type Bookmark struct {
	Title string
	Url   string
	Tags  []string
}

func (c *Config) validate() error {
	if c.Version != "v1" {
		msg := fmt.Sprintf("version=%s is not supported", c.Version)
		return errors.New(msg)
	}
	if c.DefaultBrowser == "" {
		return errors.New("default_browser is required")
	}
	if len(c.Browsers) == 0 {
		return errors.New("browsers should not be empty")
	}
	return nil
}

func (c *Config) buildCandidates(flags *Flags) string {
	var candidates []string
	for i, b := range c.Bookmarks {
		if flags.Tags == "" || isTagMatched(flags, b.Tags) {
			line := encodeBookmark(b, i)
			candidates = append(candidates, line)
		}
	}
	return strings.Join(candidates, "\n")
}

func (c *Config) selectBrowser(browser string) *Browser {
	if b := c.selectBrowserFromFlag(browser); b != nil {
		return b
	}
	return c.selectDefaultBrowser()
}

func (c *Config) selectBrowserFromFlag(browser string) *Browser {
	if browser != "" {
		for _, b := range c.Browsers {
			if b.Key == browser {
				return b
			}
		}
	}
	return nil
}

func (c *Config) selectDefaultBrowser() *Browser {
	for _, b := range c.Browsers {
		if b.Key == c.DefaultBrowser {
			return b
		}
	}
	return nil
}

func (c *Config) selectBookmarks(lines []string) []*Bookmark {
	var bookmarks []*Bookmark
	for _, selection := range lines {
		if len(selection) == 0 {
			continue
		}
		fields := strings.Split(selection, fieldSplit)
		idx, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Fatal(err)
		}
		bookmarks = append(bookmarks, c.Bookmarks[idx])
	}
	return bookmarks
}

func encodeBookmark(bk *Bookmark, idx int) string {
	tags := strings.Join(bk.Tags, ",#")
	lines := []string{
		fmt.Sprintf("%d", idx),
		fmt.Sprintf("%.30s", bk.Title),
		fmt.Sprintf("%.50s", tags),
		bk.Url,
	}
	return strings.Join(lines, fieldSplit)
}

func decodeBookmark(line string) *Bookmark {
	fields := strings.Split(line, fieldSplit)
	return &Bookmark{
		Title: fields[1],
		Tags:  strings.Split(fields[2], ",#"),
		Url:   fields[3],
	}
}

func (c *Config) writeBookmark(flags *Flags) {
	confFilePath := flags.confFilePath()
	c.writeBookmarkTo(confFilePath)
	c.writeBookmarkToMultiSources()
}

func (c *Config) writeBookmarkTo(filePath string) {
	os.Truncate(filePath, 0)
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	d, err := yaml.Marshal(c)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = f.WriteString(string(d))
	if err != nil {
		log.Fatalln(err)
	}
}

func (c *Config) writeBookmarkToMultiSources() {
	for _, source := range c.MultiSources {
		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		path := path.Join(userHomeDir, source)
		c.writeBookmarkTo(path)
	}
}

func isTagMatched(f *Flags, tags []string) bool {
	splitTags := strings.Split(f.Tags, ",")
	sets := make(map[string]bool)
	for _, t := range splitTags {
		sets[t] = false
	}
	for _, t := range tags {
		if _, ok := sets[t]; ok {
			sets[t] = true
		}
	}
	for _, matched := range sets {
		if !matched {
			return false
		}
	}
	return true
}
