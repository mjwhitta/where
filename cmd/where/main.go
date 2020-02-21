package main

import (
	"gitlab.com/mjwhitta/cli"
	hl "gitlab.com/mjwhitta/hilighter"
	"gitlab.com/mjwhitta/where"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			if flags.verbose {
				panic(r.(error).Error())
			}
			errx(Exception, r.(error).Error())
		}
	}()

	validate()

	hl.Println(where.Is(cli.Arg(0)))
}
