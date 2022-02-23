## 简介

[![GitHub release](https://img.shields.io/github/release/elliotxx/mdfmt.svg)](https://github.com/elliotxx/mdfmt/releases)
[![Github All Releases](https://img.shields.io/github/downloads/elliotxx/mdfmt/total.svg)](https://github.com/elliotxx/mdfmt/releases)
[![license](https://img.shields.io/github/license/elliotxx/mdfmt.svg)](https://github.com/elliotxx/mdfmt/blob/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/elliotxx/mdfmt.svg)](https://pkg.go.dev/github.com/elliotxx/mdfmt)
[![Coverage Status](https://coveralls.io/repos/github/elliotxx/mdfmt/badge.svg)](https://coveralls.io/github/elliotxx/mdfmt)

> 💡 一款 Markdown 格式化工具，和 gofmt 比较类似，不过格式化的对象是 Markdown 文本

<p align="center">
  <img src="assets/demo.svg">
</p>

## 语言

[English](https://github.com/elliotxx/mdfmt/blob/master/README.md) | [简体中文](https://github.com/elliotxx/mdfmt/blob/master/README-zh.md)

## 安装

### Homebrew

`elliotxx/tap` 有 MacOS 和 GNU/Linux 的预编译二进制版本可用：

```
brew install elliotxx/tap/mdfmt
```

### 从源码构建

使用 Go 1.17+ 版本，你可以通过 `go install` 直接从源码安装 `mdfmt`：

```
go install github.com/elliotxx/mdfmt/cmd/mdfmt@latest
```

*注意*: 你将基于代码仓库最新的可用版本安装 `mdfmt`，尽管主分支的最新提交应该始终是一个稳定和可用的版本，但这不是安装和使用 `mdfmt` 的推荐方式。通过 `go install` 安装的 `mdfmt` 版本输出将显示默认版本号（default-version）。

## 使用

```
$ mdfmt -h
A Markdown formatter that follow the CommonMark. Like gofmt, but for Markdown.

Usage:
  mdfmt [flags] [path ...]

Examples:
  # Format specified Markdown file, and write to stdout
  mdfmt README.md
  
  # Format and rewrite for specified Markdown file
  mdfmt -w README.md
  
  # Display diffs instead of rewriting Markdown files
  mdfmt -d README.md
  
  # Format, rewrite, and display diffs for specified Markdown file
  mdfmt -d -w README.md
  
  # Format and rewrite all Markdown file in current directory
  mdfmt -w *.md
  
  # Recursive format and rewrite all Markdown file in current directory
  mdfmt -w .
  
  # Format and rewrite the specified Markdown file and directory
  mdfmt -w README.md testdir/
  
  # Format stdin to stdout
  cat README.md | mdfmt
  
  # Show version info
  mdfmt -V

Flags:
  -d, --diff      display diffs instead of rewriting files
  -h, --help      help for mdfmt
  -V, --version   show version info
  -w, --write     write result to (source) file instead of stdout
```

## 感谢

* CLI Template from [elliotxx/go-cli-prototype](https://github.com/elliotxx/go-cli-prototype)
* The specification follows [GFM](https://github.github.com/gfm/)/[CommonMark](https://commonmark.org/)
* The Markdown engine uses [lute](https://github.com/88250/lute), cool!
