package ggrep

import (
	"os"
	"log"
	"fmt"
	"bufio"
	"regexp"
	"strings"
)

func FindStdIn(file *os.File, pattern string) {
	matches, err := find(file, pattern)

	if err != nil {
		log.Fatal(err)
	}

	for _, v := range matches {
		fmt.Println(v)
	}
}

func Find(fileNames []string, pattern string) {
	// This function should handle when input is from a
	// file or stdin and pass the searching to find.
	matches := map[string]string{}
	for _, f := range fileNames {
		file, err := os.Open(f)
		if err != nil {
			log.Fatal(err)
		}

		found, err := find(file, pattern)
		if err != nil {
			log.Fatal(err)
		}
		matches[file.Name()] = found
	}

	for k, v := range matches {
		fmt.Println(k)
		fmt.Println(v)
	}
}

func find(file *os.File, pattern string) (string, error) {
	var matches []string
	regex, err := regexp.Compile(pattern)

	if err != nil {
		return "", err
	}

        line := bufio.NewScanner(file)
	for line.Scan() {
		oldText := line.Text()
		if regex.MatchString(oldText) {
			newWord := []string{Colors["Red"], regex.FindString(pattern), Colors["Reset"]}
			replacement := strings.Join(newWord, "")
			newText := regex.ReplaceAllString(oldText, replacement)
			matches = append(matches, newText)
			matches = append(matches, "\n")
		} else {
			continue
		}
	}

	m := strings.Join(matches, "")
	return m, nil
}
