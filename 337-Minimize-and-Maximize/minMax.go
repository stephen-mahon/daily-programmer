package main

import (
	"fmt"
	"math"
)

var message = "Minimize and Maximize"
var pi = 3.14
var length = 100.0
var gr = (math.Sqrt(5) + 1) / 2
var tol = 1e-5

func main() {
	fmt.Println(message)
}

// Golden-section search
/*
	to find the minimum of f on [a, b]
	f: a strictly unimodal function on [a,b]
*/
func gss(a, b, tol float64) float64 {
	c := b - (b-a)/gr
	d := a + (b-a)/gr

	for math.Abs(b-a) > tol {
		if f(c) < f(d) {
			b = d
		} else {
			a = c
		}
		c = b - (b-a)/gr
		d = a + (b-a)/gr
	}
	return (b + a) / 2
}

func f(val float64) float64 {
	return -1
}
