package main

import (
	"fmt"
	"time"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	x := uint64(100000)

	start := time.Now()
	PopCountLoop(x)
	fmt.Printf("Loop time: %f\n", time.Since(start).Seconds())

	start = time.Now()
	PopCountExpression(x)
	fmt.Printf("Expression time: %f\n", time.Since(start).Seconds())
}

// PopCountLoop returns the population count (number of set bits) of x.
func PopCountLoop(x uint64) int {
	var out byte
	for i := uint(0); i < 8; i++ {
		out += pc[byte(x>>(i*8))]
	}
	return int(out)
}

// PopCountExpression returns the population count (number of set bits) of x.
func PopCountExpression(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
