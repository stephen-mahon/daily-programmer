package main

import "fmt"

func main() {
	pi := leibnizFormula(10000)
	fmt.Println(pi)
}

func leibnizFormula(n int) float64 {
	sign := false
	var pi float64
	for i := 1; i <= n; i += 2 {
		sign = !sign
		if sign {
			pi += 4 / float64(i)
		} else {
			pi -= 4 / float64(i)
		}
	}
	return pi
}
