# go-bookmark

Mange your browser bookmarks with CLI.

## Install

https://github.com/kenju/go-bookmark/releases

## Usage

### Open Bookmark

Add `~/.config/go-bookmark/bookmark.yaml` as follows:

```yaml
version: v1
default_browser: firefox
browsers:
- key: chrome
  command: /Applications/Google Chrome.app/Contents/MacOS/Google Chrome
  args:
  - --args
  - --kiosk
- key: firefox
  command: /Applications/Firefox.app/Contents/MacOS/firefox
  args:
  - -new-tab
bookmarks:
- title: Example Domain
  url: https://example.com
  tags:
  - test
- title: go-bookmark
  url: https://github.com/kenju/go-bookmark
  tags:
  - develop
- title: GitHub
  url: https://github.com
  tags:
  - develop
```

Run `go-bookmark open` and select bookmark(s) to open.

[![asciicast](https://asciinema.org/a/dS6owVBhyjOjRJXvHzaqF42Pr.svg)](https://asciinema.org/a/dS6owVBhyjOjRJXvHzaqF42Pr)

### Add Bookmark

Run `go-bookmark add` to add a new bookmark.

> Pro Tips:
>
> You can also update the bookmark conf file directly instead.

### Delete Bookmark

Run `go-bookmark delete` to delete bookmark(s).

> Pro Tips:
>
> You can also delete the bookmark items from the conf file directly instead.

## Commands

```
$ bookmark help
Usage: bookmark <flags> <subcommand> <subcommand args>

Subcommands:
        add              add new command
        delete           delete new command
        flags            describe all known top-level flags
        help             describe subcommands and their syntax
        open             open new command
        stats            show version
        version          show bookmark version


Use "bookmark flags" for a list of top-level flags
```

## Requirements

- `peco` is installed and in your $PATH
    - https://github.com/peco/peco

## Development

### Unit Tests

Run `make test` locally.

[GitHub Actions](https://github.com/kenju/go-bookmark/actions/workflows/ci-test.yml) runs when commits pushed.

### Release

`git tag` and push to the `master` branch.

[`goreleaser`](https://goreleaser.com/) is triggered via [GitHub Actions](https://github.com/kenju/go-bookmark/actions/workflows/release.yml).
