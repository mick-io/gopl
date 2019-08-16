// Exercise 1.2:  Modify the echo program to print the index and value of each of it's arguments, one per line.
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, s := range os.Args {
		fmt.Printf("line %d: %s\n", i, s)
	}
}
