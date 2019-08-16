// Exercise 1.4: Modify dup2 to print the names of all files in which each duplicated line occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fpaths := os.Args[1:]
	hasDuplicate := func(f *os.File) bool {
		counts, input := make(map[string]int), bufio.NewScanner(f)
		for input.Scan() {
			text := input.Text()
			counts[text]++
			if counts[text] > 1 {
				return true
			}
		}
		return false
		// NOTE: ignore potential errors from input.Err()
	}

	if len(fpaths) == 0 {
		if hasDuplicate(os.Stdin) {
			fmt.Println("stdin")
		}
		return
	}

	for _, fp := range fpaths {
		f, err := os.Open(fp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dip2: %v\n", err)
			continue
		}
		if hasDuplicate(f) {
			fmt.Println(filepath.Base(fp))
		}
		f.Close()
	}
}
