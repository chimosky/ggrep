package ggrep

import (
	"os"
	"fmt"
)

func main() {

	fileNames := os.Args[2:]
	pattern := os.Args[1]

	if len(os.Args) <= 2 {
		if os.Stdin != nil {
			FindStdIn(os.Stdin, pattern)
		} else {
			usage()
			os.Exit(1)
		}
	}


	Find(fileNames, pattern)
}

func usage() {
	fmt.Println("usage: ggrep <pattern> <filname>...")
}
