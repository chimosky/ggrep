package ggrep

import (
	"os"
	"bufio"
	"regexp"
	"strings"
)

func Find(fileNames string, pattern string) (text string, err error) {
	// This function should handle when input is from a
	// file or stdin and pass the searching to find.
	for _, f := range fileNames {
		file := os.Open(f)
		found, err := find(file, pattern)
	}

	return found, err
}

func find(file *os.File, pattern string) (string, error) {
	matches := []string{}
	regex, err := regexp.Compile(pattern)

	if err != nil {
		return matches, err
	}

        line := bufio.NewScanner(file)
	for line.Scan() {
		oldText := line.Text()
		if regex.MatchString(oldText) {
			newWord := []strings{Colors[Red], regex.FindString(pattern), Colors[Reset]}
			replacement := strings.Join(newWord, "")
			newText := regex.ReplaceAllString(oldText, replacement)
			append(matches, newText)
		} else {
			append(matches, oldText)
		}
	}

	return strings.Join(matches), nil
}
