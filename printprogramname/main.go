package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]

	lastSlash := -1
	for i, r := range name {
		if r == '/' {
			lastSlash = i
		}
	}
	fileName := name[lastSlash+1:]
	for _, r := range fileName {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
