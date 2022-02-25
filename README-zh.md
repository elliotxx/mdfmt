## 简介

[![GitHub release](https://img.shields.io/github/release/elliotxx/mdfmt.svg)](https://github.com/elliotxx/mdfmt/releases)
[![Github All Releases](https://img.shields.io/github/downloads/elliotxx/mdfmt/total.svg)](https://github.com/elliotxx/mdfmt/releases)
[![docker pulls](https://img.shields.io/docker/pulls/elliotxx/mdfmt)](https://hub.docker.com/r/elliotxx/mdfmt)
[![license](https://img.shields.io/github/license/elliotxx/mdfmt.svg)](https://github.com/elliotxx/mdfmt/blob/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/elliotxx/mdfmt.svg)](https://pkg.go.dev/github.com/elliotxx/mdfmt)
[![Coverage Status](https://coveralls.io/repos/github/elliotxx/mdfmt/badge.svg)](https://coveralls.io/github/elliotxx/mdfmt)
[![gitmoji](https://img.shields.io/badge/gitmoji-%20😜%20😍-FFDD67.svg?style=flat-square)](https://gitmoji.dev)

> 💡 一款 Markdown 格式化工具，和 gofmt 比较类似，不过格式化的对象是 Markdown 文本

<p align="center">
  <img src="assets/demo.svg">
</p>

最近在频繁用 Markdown 写文档，内容一多就不好管理格式，经常看到一坨坨的 Markdown 内容挤在一起，导致有强迫症的我的眉头也经常挤在一起。

好在发现了强大的结构化的 Markdown 引擎 [lute](https://github.com/88250/lute)，它对中文语境的支持也很好，于是借助 lute 引擎开发了这款 Markdown 格式化命令行工具 [mdfmt](https://github.com/elliotxx/mdfmt)，欢迎大家试用 👏🏻

## 📜 语言

[English](https://github.com/elliotxx/mdfmt/blob/master/README.md) | [简体中文](https://github.com/elliotxx/mdfmt/blob/master/README-zh.md)

## ✨ 特性

* 支持多种输入：标准输入、文件、目录、通配符，其中指定目录会递归格式化目录下所有 Markdown 文件
* 支持重写：将结果写入（源）文件而不是标准输出
* 支持显示差异：显示 Markdown 格式化前后的差异（diff），而不是重写文件
* 支持列出格式化的文件
* 跨平台：Linux, Windows, Mac
* 一键安装：支持通过 `Homebrew`、`go install` 等方式一键安装 `mdfmt`

## 🛠️ 安装

### 二进制安装（跨平台: windows, linux, mac ...）

从二进制安装，只需从 `mdfmt` 的 [发布页面](https://github.com/elliotxxx/mdfmt/releases) 下载对应平台的二进制文件，然后将二进制文件放在命令行能访问到的目录中即可。

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

### Docker

Docker 用户可以用以下命令拉取 `mdfmt` 的镜像：

```
docker pull elliotxx/mdfmt
```

验证:

```bash
$ docker run --rm elliotxx/mdfmt:latest mdfmt -h
...
$ docker run --rm elliotxx/mdfmt:latest mdfmt -V
...
$ docker run -v $PWD:$PWD --rm elliotxx/mdfmt:latest mdfmt -d /Users/yym/workspace/mdfmt/pkg/md/testdata/hello-more.md
diff -u /Users/yym/workspace/mdfmt/pkg/md/testdata/hello-more.md.orig /Users/yym/workspace/mdfmt/pkg/md/testdata/hello-more.md
--- /Users/yym/workspace/mdfmt/pkg/md/testdata/hello-more.md.orig
+++ /Users/yym/workspace/mdfmt/pkg/md/testdata/hello-more.md
@@ -1,6 +1,7 @@
 # hello
+
 > hello

-|name|age|
-|--|--|
-|Mike|18|
+| name | age |
+| ---- | --- |
+| Mike | 18  |
```

## ⚡ 使用

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
  
  # List files whose formatting differs from mdfmt's
  mdfmt -l .
  
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
  -l, --list      list files whose formatting differs from mdfmt's
  -V, --version   show version info
  -w, --write     write result to (source) file instead of stdout
```

## 🙏 感谢

* Markdown 引擎使用 [88250/lute](https://github.com/88250/lute), 很酷!
* 命令行工具模板来自 [elliotxx/go-cli-prototype](https://github.com/elliotxx/go-cli-prototype)
* Markdown 规范遵循 [GFM](https://github.github.com/gfm/)/[CommonMark](https://commonmark.org/)
