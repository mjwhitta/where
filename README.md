# where

A `which` or `command -v` like tool for Golang.

## How to install

Open a terminal and run the following:

```
$ go get -u gitlab.com/mjwhitta/where/cmd/where
```

## Usage

```
package main

import (
    "fmt"

    "gitlab.com/mjwhitta/where"
)

func main() {
    fmt.Println(where.Is("bash"))
}
```

## Links

- [Source](https://gitlab.com/mjwhitta/where)
