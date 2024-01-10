package main

import (
	"fmt"
	"log"
	"os"
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
		matches := FindStdIn(pattern)
		fmt.Println(matches)
	} else {
		match := map[string]string{}
		for _, f := range fileNames {
			file, err := os.Open(f)
			defer file.Close()

			if err != nil {
				log.Fatal(err)
			}

			matches := Find(file, pattern)

			if len(fileNames) > 1 {
				match[file.Name()] = matches[:len(matches)-1]
			} else {
				match[""] = matches[:len(matches)-1]
			}
		}

		for k, v := range match {
			if len(k) == 0 {
				fmt.Println(v)
			} else {
				fmt.Println(k)
				fmt.Println(v)
			}
		}
	}
}
