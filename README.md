## Introduction

[![GitHub release](https://img.shields.io/github/release/elliotxx/mdfmt.svg)](https://github.com/elliotxx/mdfmt/releases)
[![Github All Releases](https://img.shields.io/github/downloads/elliotxx/mdfmt/total.svg)](https://github.com/elliotxx/mdfmt/releases)
[![license](https://img.shields.io/github/license/elliotxx/mdfmt.svg)](https://github.com/elliotxx/mdfmt/blob/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/elliotxx/mdfmt.svg)](https://pkg.go.dev/github.com/elliotxx/mdfmt)
[![Coverage Status](https://coveralls.io/repos/github/elliotxx/mdfmt/badge.svg)](https://coveralls.io/github/elliotxx/mdfmt)

> Like gofmt, but for Markdown. Based on CommonMark.

CLI Template from [elliotxx/go-cli-prototype](https://github.com/elliotxx/go-cli-prototype)

## Intallation

```
go install github.com/elliotxx/mdfmt@latest
```

## Usage

All targets:

```
$ make help
help                           This help message :)
test                           Run the tests
cover                          Generates coverage report
cover-html                     Generates coverage report and displays it in the browser
format                         Format source code
lint                           Lint, will not fix but sets exit code on error
lint-fix                       Lint, will try to fix errors and modify code
doc                            Start the documentation server with godoc
gen-version                    Update version
clean                          Clean build bundles
build-all                      Build all platforms
build-darwin                   Build for MacOS
build-linux                    Build for Linux
build-windows                  Build for Windows
```
