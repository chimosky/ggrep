package ggrep

import (
	"os"
	"log"
	"fmt"
	"bufio"
	"regexp"
	"strings"
)

var (
        Colors = map[string]string {
                "Reset" : "\033[0m",
                "Red"   : "\033[1;31m",
        }
)

func FindStdIn(pattern string) {
	matches, err := findStdin(pattern)

	if err != nil {
		log.Fatal(err)
	}


	fmt.Println(matches[:len(matches)-1])
}

func Find(f *os.File, pattern string) {
	matches := map[string]string{}
	found, err := find(f, pattern)

	if err != nil {
		log.Fatal(err)
	}
	matches[f.Name()] = found[:len(found)-1]

	if len(matches[f.Name()]) == 0 {
		fmt.Println("No matches found")
		os.Exit(0)
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
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		oldText := scanner.Text()
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

func findStdin(pattern string) (string, error) {
	var matches []string
	regex, err := regexp.Compile(pattern)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		oldText := scanner.Text()
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
