## ç®ä»

[![GitHub release](https://img.shields.io/github/release/elliotxx/mdfmt.svg)](https://github.com/elliotxx/mdfmt/releases)
[![Github All Releases](https://img.shields.io/github/downloads/elliotxx/mdfmt/total.svg)](https://github.com/elliotxx/mdfmt/releases)
[![docker pulls](https://img.shields.io/docker/pulls/elliotxx/mdfmt)](https://hub.docker.com/r/elliotxx/mdfmt)
[![license](https://img.shields.io/github/license/elliotxx/mdfmt.svg)](https://github.com/elliotxx/mdfmt/blob/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/elliotxx/mdfmt.svg)](https://pkg.go.dev/github.com/elliotxx/mdfmt)
[![Coverage Status](https://coveralls.io/repos/github/elliotxx/mdfmt/badge.svg)](https://coveralls.io/github/elliotxx/mdfmt)
[![gitmoji](https://img.shields.io/badge/gitmoji-%20ð%20ð-FFDD67.svg?style=flat-square)](https://gitmoji.dev)

> ð¡ ä¸æ¬¾ Markdown æ ¼å¼åå·¥å·ï¼å gofmt æ¯è¾ç±»ä¼¼ï¼ä¸è¿æ ¼å¼åçå¯¹è±¡æ¯ Markdown ææ¬

<p align="center">
  <img src="assets/demo.svg">
</p>

æè¿å¨é¢ç¹ç¨ Markdown åææ¡£ï¼åå®¹ä¸å¤å°±ä¸å¥½ç®¡çæ ¼å¼ï¼ç»å¸¸çå°ä¸å¨å¨ç Markdown åå®¹æ¤å¨ä¸èµ·ï¼å¯¼è´æå¼ºè¿«ççæççå¤´ä¹ç»å¸¸æ¤å¨ä¸èµ·ã

å¥½å¨åç°äºå¼ºå¤§çç»æåç Markdown å¼æ [lute](https://github.com/88250/lute)ï¼å®å¯¹ä¸­æè¯­å¢çæ¯æä¹å¾å¥½ï¼äºæ¯åå© lute å¼æå¼åäºè¿æ¬¾ Markdown æ ¼å¼åå½ä»¤è¡å·¥å· [mdfmt](https://github.com/elliotxx/mdfmt)ï¼æ¬¢è¿å¤§å®¶è¯ç¨ ðð»

## ð è¯­è¨

[English](https://github.com/elliotxx/mdfmt/blob/master/README.md) | [ç®ä½ä¸­æ](https://github.com/elliotxx/mdfmt/blob/master/README-zh.md)

## â¨ ç¹æ§

* æ¯æå¤ç§è¾å¥ï¼æ åè¾å¥ãæä»¶ãç®å½ãééç¬¦ï¼å¶ä¸­æå®ç®å½ä¼éå½æ ¼å¼åç®å½ä¸ææ Markdown æä»¶
* æ¯æéåï¼å°ç»æåå¥ï¼æºï¼æä»¶èä¸æ¯æ åè¾åº
* æ¯ææ¾ç¤ºå·®å¼ï¼æ¾ç¤º Markdown æ ¼å¼åååçå·®å¼ï¼diffï¼ï¼èä¸æ¯éåæä»¶
* æ¯æååºæ ¼å¼åçæä»¶
* è·¨å¹³å°ï¼Linux, Windows, Mac
* ä¸é®å®è£ï¼æ¯æéè¿ `Homebrew`ã`go install` ç­æ¹å¼ä¸é®å®è£ `mdfmt`

## ð ï¸ å®è£

### äºè¿å¶å®è£ï¼è·¨å¹³å°: windows, linux, mac ...ï¼

ä»äºè¿å¶å®è£ï¼åªéä» `mdfmt` ç [åå¸é¡µé¢](https://github.com/elliotxx/mdfmt/releases) ä¸è½½å¯¹åºå¹³å°çäºè¿å¶æä»¶ï¼ç¶åå°äºè¿å¶æä»¶æ¾å¨å½ä»¤è¡è½è®¿é®å°çç®å½ä¸­å³å¯ã

### Homebrew

`elliotxx/tap` æ MacOS å GNU/Linux çé¢ç¼è¯äºè¿å¶çæ¬å¯ç¨ï¼

```
brew install elliotxx/tap/mdfmt
```

### ä»æºç æå»º

ä½¿ç¨ Go 1.17+ çæ¬ï¼ä½ å¯ä»¥éè¿ `go install` ç´æ¥ä»æºç å®è£ `mdfmt`ï¼

```
go install github.com/elliotxx/mdfmt/cmd/mdfmt@latest
```

*æ³¨æ*: ä½ å°åºäºä»£ç ä»åºææ°çå¯ç¨çæ¬å®è£ `mdfmt`ï¼å°½ç®¡ä¸»åæ¯çææ°æäº¤åºè¯¥å§ç»æ¯ä¸ä¸ªç¨³å®åå¯ç¨ççæ¬ï¼ä½è¿ä¸æ¯å®è£åä½¿ç¨ `mdfmt` çæ¨èæ¹å¼ãéè¿ `go install` å®è£ç `mdfmt` çæ¬è¾åºå°æ¾ç¤ºé»è®¤çæ¬å·ï¼default-versionï¼ã

### Docker

Docker ç¨æ·å¯ä»¥ç¨ä»¥ä¸å½ä»¤æå `mdfmt` çéåï¼

```
docker pull elliotxx/mdfmt
```

éªè¯:

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

## â¡ ä½¿ç¨

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

## ð æè°¢

* Markdown å¼æä½¿ç¨ [88250/lute](https://github.com/88250/lute), å¾é·!
* å½ä»¤è¡å·¥å·æ¨¡æ¿æ¥èª [elliotxx/go-cli-prototype](https://github.com/elliotxx/go-cli-prototype)
* Markdown è§èéµå¾ª [GFM](https://github.github.com/gfm/)/[CommonMark](https://commonmark.org/)
* ä½¿ç¨ [gitmoji-cli](https://github.com/carloscuesta/gitmoji-cli) æäº¤æ¼äº®ç git commit message
