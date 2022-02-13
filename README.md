## Introduction

[![GitHub release](https://img.shields.io/github/release/elliotxx/mdfmt.svg)](https://github.com/elliotxx/mdfmt/releases)
[![Github All Releases](https://img.shields.io/github/downloads/elliotxx/mdfmt/total.svg)](https://github.com/elliotxx/mdfmt/releases)
[![license](https://img.shields.io/github/license/elliotxx/mdfmt.svg)](https://github.com/elliotxx/mdfmt/blob/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/elliotxx/mdfmt.svg)](https://pkg.go.dev/github.com/elliotxx/mdfmt)
[![Coverage Status](https://coveralls.io/repos/github/elliotxx/mdfmt/badge.svg)](https://coveralls.io/github/elliotxx/mdfmt)

> A Markdown formatter that follow the CommonMark. Like gofmt, but for Markdown.

## Intallation

### Homebrew

The `elliotxx/tap` has macOS and GNU/Linux pre-built binaries available:

```
brew install elliotxx/tap/mdfmt
```

### Build from Source

Starting with Go 1.17, you can install `mdfmt` from source using go install:

```
go install github.com/elliotxx/mdfmt@latest
```

*NOTE*: This will install `mdfmt` based on the latest available code base. Even though the goal is that the latest commit on the main branch should always be a stable and usable version, this is not the recommended way to install and use `mdfmt`. The version output will show `mdfmt` version (default-version) for go install based builds.

## Usage

```
$ mdfmt -h
A Markdown formatter that follow the CommonMark. Like gofmt, but for Markdown.

Usage:
  mdfmt [flags] [path ...]

Examples:
  # Format markdown file, and write to stdout
  mdfmt README.md
  
  # Format and rewrite markdown file
  mdfmt -w README.md
  
  # Format and rewrite markdown file and directory
  mdfmt -w README.md testdir/
  
  # Format stdin to stdout
  cat README.md | mdfmt
  
  # Show version info
  mdfmt -V

Flags:
  -h, --help      help for mdfmt
  -V, --version   show version info
  -w, --write     write result to (source) file instead of stdout
```

## Thanks

* CLI Template from [elliotxx/go-cli-prototype](https://github.com/elliotxx/go-cli-prototype)
* The specification follows [GFM](https://github.github.com/gfm/)/[CommonMark](https://commonmark.org/)
* The Markdown engine uses [lute](https://github.com/88250/lute), cool!
