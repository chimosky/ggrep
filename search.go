package main

import (
	"os"
	"log"
	"fmt"
	"sync"
	"bufio"
	"regexp"
	"strings"
)

var (
        Colors = map[string]string {
                "Reset" : "\033[0m",
                "Red"   : "\033[1;31m",
        }
	match = make(chan string, 1024)
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
	scanner := bufio.NewScanner(file)
	var wg sync.WaitGroup

	for scanner.Scan() {
		wg.Add(1)
		go searchLoop(scanner.Text(), pattern, &wg)
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

func findStdin(pattern string) (string, error) {
	var matches []string
	scanner := bufio.NewScanner(os.Stdin)
	var wg sync.WaitGroup

	for scanner.Scan() {
		wg.Add(1)
		go searchLoop(scanner.Text(), pattern, &wg)
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

func searchLoop(searchPattern, pattern string, wg *sync.WaitGroup) {
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
