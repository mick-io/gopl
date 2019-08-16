package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/*
Under the covers bufio.Scanner, ioutil.ReadFile, and ioutil.WriteFile use the Read and Write methods of *os.File, but
it's rare that most programmers need to access those lower-level routines directly. The higher-level functions like
those from bufio and io/ioutil are easier to use.
*/

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		// ReadFile returns a byte slice that must be converted into a string so it can be split by strings.Split.
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}
}
