package main

import (
	"flag"
	"fmt"
)

var title = "Nearest prime numbers"

func main() {

	val := flag.Int("prime", 541, "the input value.")
	flag.Parse()
	fmt.Printf("%s: %v\n", title, *val)
	if isPrime(*val, 2) {
		fmt.Printf("%v is prime\n", *val)
	} else {
		prev := *val - 1
		for !isPrime(prev, 2) {
			prev--
		}
		next := *val + 1
		for !isPrime(next, 2) {
			next++
		}
		fmt.Printf("%v < %v < %v", prev, *val, next)

	}
}

func isPrime(p, n int) bool {
	if p == 0 || p == 1 {
		return false
	}

	if p == n {
		return true
	}

	if p%n == 0 {
		return false
	}
	n++
	return isPrime(p, n)
}
