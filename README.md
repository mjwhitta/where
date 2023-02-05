# where

[![Yum](https://img.shields.io/badge/-Buy%20me%20a%20cookie-blue?labelColor=grey&logo=cookiecutter&style=for-the-badge)](https://www.buymeacoffee.com/mjwhitta)

[![Go Report Card](https://goreportcard.com/badge/github.com/mjwhitta/where?style=for-the-badge)](https://goreportcard.com/report/github.com/mjwhitta/where)
![License](https://img.shields.io/github/license/mjwhitta/where?style=for-the-badge)

## What is this?

A `which` or `command -v` like tool for Go.

## How to install

Open a terminal and run the following:

```
$ go get --ldflags "-s -w" --trimpath -u github.com/mjwhitta/where
$ go install --ldflags "-s -w" --trimpath \
    github.com/mjwhitta/where/cmd/where@latest
```

Or compile from source:

```
$ git clone https://github.com/mjwhitta/where.git
$ cd where
$ git submodule update --init
$ make
```

## Usage

```
package main

import (
    "fmt"

    "github.com/mjwhitta/where"
)

func main() {
    fmt.Println(where.Is("bash"))
}
```

## Links

- [Source](https://github.com/mjwhitta/where)
