// Exercise 1.3: Experiment to measure the difference in running time between our potentially inefficient versions and
// the one that uses strings.Join.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	echo1 := func() {
		var s, sep string
		for i := 1; i < len(os.Args); i++ {
			s += sep + os.Args[i]
			sep = " "
		}
		fmt.Println(s)
	}

	echo2 := func() {
		s, sep := "", ""
		for _, arg := range os.Args[1:] {
			s += sep + arg
			sep = " "
		}
		fmt.Println(s)
	}

	echo3 := func() {
		fmt.Println(strings.Join(os.Args[1:], " "))
	}

	start := time.Now()
	echo1()
	echo1Time := time.Since(start).Seconds()

	start = time.Now()
	echo2()
	echo2Time := time.Since(start).Seconds()

	start = time.Now()
	echo3()
	echo3Time := time.Since(start).Seconds()

	msg := "echo1 Time: %e seconds\necho2 Time: %e seconds\necho3 Time: %e seconds\n"
	fmt.Printf(msg, echo1Time, echo2Time, echo3Time)
}
