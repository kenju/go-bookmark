package main

import (
	"fmt"
	"sort"
)

type Stats struct {
	TotalCount int
	TotalTags  map[string]int
}

func NewStats(config *Config) *Stats {
	st := &Stats{
		TotalCount: 0,
		TotalTags:  make(map[string]int),
	}
	st.calcTotalTags(config)
	return st
}

func (st *Stats) calcTotalTags(config *Config) {
	for _, bookmark := range config.Bookmarks {
		st.TotalCount += 1
		for _, tag := range bookmark.Tags {
			st.TotalTags[tag] += 1
		}
	}
}

func (st *Stats) Show() {
	fmt.Printf("Total Count: %d\n", st.TotalCount)
	var sorted []string
	for tag, cnt := range st.TotalTags {
		line := fmt.Sprintf("%s: %d\n", tag, cnt)
		sorted = append(sorted, line)
	}
	sort.Strings(sorted)
	for _, line := range sorted {
		fmt.Print(line)
	}
}
