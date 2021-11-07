package main

import (
	"reflect"
	"testing"
)

func TestIsTagMatched(t *testing.T) {
	var tests = []struct {
		testName string
		flags    *Flags
		tags     []string
		expected bool
	}{
		{
			testName: "empty",
			flags: &Flags{
				Tags: "",
			},
			tags:     []string{""},
			expected: true,
		},
		{
			testName: "empty flags tags",
			flags: &Flags{
				Tags: "",
			},
			tags:     []string{"foo"},
			expected: false,
		},
		{
			testName: "empty tags",
			flags: &Flags{
				Tags: "foo",
			},
			tags:     []string{""},
			expected: false,
		},
		{
			testName: "single match",
			flags: &Flags{
				Tags: "foo",
			},
			tags:     []string{"foo"},
			expected: true,
		},
		{
			testName: "single unmatch",
			flags: &Flags{
				Tags: "foo",
			},
			tags:     []string{"bar"},
			expected: false,
		},
		{
			testName: "match with multi flags tags",
			flags: &Flags{
				Tags: "foo,bar",
			},
			tags:     []string{"foo", "bar"},
			expected: true,
		},
		{
			testName: "unmatch with multi flags tags",
			flags: &Flags{
				Tags: "foo,bar",
			},
			tags:     []string{"foo"},
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			actual := isTagMatched(tt.flags, tt.tags)
			if tt.expected != actual {
				t.Errorf("expected=%t, actual=%t\n",
					tt.expected, actual)
			}
		})
	}
}
func TestSelectBrowser(t *testing.T) {
	var tests = []struct {
		testName string
		browser  string
		config   *Config
		expected *Browser
	}{
		{
			testName: "select from default browser",
			config: &Config{
				DefaultBrowser: "chrome",
				Browsers: []*Browser{
					{
						Key:     "chrome",
						Command: "google-chrome",
						Args:    []string{},
					},
					{
						Key:     "firefox",
						Command: "firefox",
						Args:    []string{"-new-tab"},
					},
				},
			},
			expected: &Browser{
				Key:     "chrome",
				Command: "google-chrome",
				Args:    []string{},
			},
		},
		{
			testName: "select from -browser flag",
			browser:  "firefox",
			config: &Config{
				DefaultBrowser: "chrome",
				Browsers: []*Browser{
					{
						Key:     "chrome",
						Command: "google-chrome",
						Args:    []string{},
					},
					{
						Key:     "firefox",
						Command: "firefox",
						Args:    []string{"-new-tab"},
					},
				},
			},
			expected: &Browser{
				Key:     "firefox",
				Command: "firefox",
				Args:    []string{"-new-tab"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			actual := tt.config.selectBrowser(tt.browser)
			if !reflect.DeepEqual(tt.expected, actual) {
				t.Errorf("expected=%v, actual=%v\n",
					tt.expected, actual)
			}
		})
	}
}

func TestSelectBookmarks(t *testing.T) {
	var tests = []struct {
		testName string
		config   *Config
		lines    []string
		expected []*Bookmark
	}{
		{
			testName: "select one bookmark",
			config: &Config{
				Bookmarks: []*Bookmark{
					{
						Title: "foo",
						Url:   "https://foo.example.com",
						Tags:  []string{"foo"},
					},
					{
						Title: "bar",
						Url:   "https://bar.example.com",
						Tags:  []string{"bar"},
					},
					{
						Title: "buz",
						Url:   "https://buz.example.com",
						Tags:  []string{"buz"},
					},
				},
			},
			lines: []string{
				"0 | foo | (#foo) | https://foo.example.com",
			},
			expected: []*Bookmark{
				{
					Title: "foo",
					Url:   "https://foo.example.com",
					Tags:  []string{"foo"},
				},
			},
		},
		{
			testName: "select two bookmarks",
			config: &Config{
				Bookmarks: []*Bookmark{
					{
						Title: "foo",
						Url:   "foo.example.com",
						Tags:  []string{"foo"},
					},
					{
						Title: "bar",
						Url:   "bar.example.com",
						Tags:  []string{"bar"},
					},
					{
						Title: "buz",
						Url:   "buz.example.com",
						Tags:  []string{"buz"},
					},
				},
			},
			lines: []string{
				"0 | foo | (#foo) | foo.example.com",
				"2 | buz | (#buz) | buz.example.com",
			},
			expected: []*Bookmark{
				{
					Title: "foo",
					Url:   "foo.example.com",
					Tags:  []string{"foo"},
				},
				{
					Title: "buz",
					Url:   "buz.example.com",
					Tags:  []string{"buz"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			actual := tt.config.selectBookmarks(tt.lines)
			if !reflect.DeepEqual(tt.expected, actual) {
				t.Errorf("expected=%+v, actual=%+v\n",
					tt.expected, actual)
			}
		})
	}
}
