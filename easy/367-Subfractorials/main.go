package main

import (
	"flag"
	"fmt"
	"math"
)

func main() {
	n := flag.Int("n", 0, "subfactorial value")
	flag.Parse()
	fmt.Println(*n, subfactorial(*n))
}

func subfactorial(n int) int {
	return int(math.Floor(float64(factorial(n))/math.Exp(1) + 0.5))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}
