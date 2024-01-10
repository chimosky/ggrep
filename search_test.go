package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"testing"
	_ "unsafe"
)

var tests = []struct {
	name  string
	text  string
	found bool
}{
	{"found", "test", true},
	{"not found", "notest", false},
}

func TestFind(t *testing.T) {
	file, _ := os.Open("test-text.txt")
	defer file.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Find(file, tt.text)
			var found bool

			if len(result) > 0 {
				found = true
			}

			if found != tt.found {
				t.Errorf("Search failed")
			}
		})
	}
}

//go:linkname search_find github.com/chimosky/ggrep/search.find
func search_find(string, *bufio.Scanner) (string, error)

func TestFindStdin(t *testing.T) {
	log.Println("Running TestFindStdin")
	buf := make([]byte, 1024)

	file, _ := os.Open("test-text.txt")
	defer file.Close()

	_, err := file.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			sf := strings.NewReader(string(buf[:]))
			scanner := bufio.NewScanner(sf)
			result, _ := find(tt.text, scanner)
			var found bool

			if len(result) > 0 {
				found = true
			}

			if found != tt.found {
				t.Errorf("Search failed")
			}
		})
	}
}
