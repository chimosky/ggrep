package ggrep

import (
	"os"
	"fmt"
	"log"
)

const usage = "usage: ggrep <pattern> <filname>..."

func main() {

	if len(os.Args) == 1 {
		fmt.Println(usage)
		os.Exit(1)
	}

	pattern := os.Args[1]
	fileNames := os.Args[2:]

	if len(fileNames) == 0 {
		FindStdIn(pattern)
	}

	for _, f := range fileNames {
		file, err := os.Open(f)
		if err != nil {
			log.Fatal(err)
		}
		Find(file, pattern)
	}
}
