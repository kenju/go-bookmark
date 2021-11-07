# go-bookmark

Mange your browser bookmarks with CLI.

## Usage

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
