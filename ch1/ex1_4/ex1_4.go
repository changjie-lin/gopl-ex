// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.

// Exercise 1.4: Modify dup2 to print the names of all files in which
//   each duplicate line occurs
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	countsByFiles := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, countsByFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, countsByFiles)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("---%v---\n", countsByFiles[line])
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, countsByFiles map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if !arrayContains(countsByFiles[line], f.Name()) {
			countsByFiles[line] = append(countsByFiles[line], f.Name())
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

func arrayContains(array []string, value string) bool {
	for _, item := range array {
		if item == value {
			return true
		}
	}
	return false
}
