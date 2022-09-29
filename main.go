package main

import (
	"os"
	"fmt"
)

func main() {
	if len(os.Args) <= 2 {
		usage()
		os.exit(1)
	}

	fileNames := os.Args[2:]
	pattern := os.Args[1]

	for _, file := range fileNames {
		// call to search function
	}
}

func usage() {
	fmt.Println("usage: ggrep <pattern> <filname>...")
}
