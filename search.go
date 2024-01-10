package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
)

var (
	Colors = map[string]string{
		"Reset": "\033[0m",
		"Red":   "\033[1;31m",
	}
)

func FindStdIn(pattern string) string {
	scanner := bufio.NewScanner(os.Stdin)
	matches, err := find(pattern, scanner)

	if err != nil {
		log.Fatal(err)
	}

	return matches
}

func Find(f *os.File, pattern string) string {
	scanner := bufio.NewScanner(f)
	matches, err := find(pattern, scanner)

	if err != nil {
		log.Fatal(err)
	}

	return matches
}

func find(pattern string, scanner *bufio.Scanner) (string, error) {
	match := make(chan string, 1024)
	var matches []string
	var wg sync.WaitGroup

	for scanner.Scan() {
		wg.Add(1)
		go searchLoop(scanner.Text(), pattern, match, &wg)
	}
	wg.Wait()
	close(match)

	for v := range match {
		matches = append(matches, v)
		matches = append(matches, "\n")
	}

	m := strings.Join(matches, "")
	return m, nil
}

func searchLoop(searchPattern, pattern string, match chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	regex, err := regexp.Compile(pattern)

	if err != nil {
		log.Fatal(err)
	}

	oldText := searchPattern
	if regex.MatchString(oldText) {
		newWord := []string{Colors["Red"], regex.FindString(pattern), Colors["Reset"]}
		replacement := strings.Join(newWord, "")
		newText := regex.ReplaceAllString(oldText, replacement)
		match <- newText
	}
}
