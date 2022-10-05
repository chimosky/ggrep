package ggrep

import (
	"bufio"
	"regexp"
	"strings"
)

func Find(text , pattern string) (string, err) {
	// This function should handle when input is from a
	// file or stdin and pass the searching to find.
}

func find(file *File, pattern string) (string, err) {
	matches := make([]string)
	regex, err := regexp.Compile(pattern)

	if err != nil {
		return matches, err
	}

        line := bufio.NewScanner(file)
	for l := line.Scan() {
		oldText := l.Text()
		if regex.MatchString(oldText) {
			newWord := strings.Join([]strings{Colors[Red], regex.FindString(pattern), Colors[Reset]}, "")
			newText := regex.ReplaceAllString(oldText, newMatch)
			append(matches, newText)
		} else {
			append(matches, oldText)
		}
	}

	return strings.Join(matches), nil
}
