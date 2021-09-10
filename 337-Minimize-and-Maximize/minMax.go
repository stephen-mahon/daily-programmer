package main

import (
	"fmt"
	"math"
)

var message = "Minimize and Maximize"
var gr = (math.Sqrt(5) + 1) / 2
var tol = 1e-1

func main() {
	fmt.Printf("%v\n\n", message)
	fmt.Printf("Challenge 1\nA piece of wire 100 cm in length is bent into the shape of a sector of a circle. Find the angle of the sector that maximizes the area A of the sector.\n")
	fmt.Printf("%.1f\n", gss(0, 180, tol))
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
		if f(c) > f(d) {
			b = d
		} else {
			a = c
		}
		c = b - (b-a)/gr
		d = a + (b-a)/gr
	}
	return (b + a) / 2
}

/*
	Perimeter of Sector = 2r + 2 pi r (theta/360) = 2r(1 + pi(theta/360))
	Perimeter = 100 cm
	->	2r (1 + pi (theta/360)) = 100
		r = 100 / 2(1 + pi (theta/360))

	Area = (theta/360) pi r^2
	In terms of only theta
	Area = (theta/360) pi (100 / 2(1 + pi (theta/360)))^2
	Area = (theta/360) pi 10000 / 4(1 + 2 pi (theta/360) + (theta/360)^2)
*/
func f(theta float64) float64 {

	return ((theta / 360) * math.Pi * 10000.0) / (4.0 * (1 + 2.0*math.Pi*(theta/360) + math.Pi*(theta/360)*math.Pi*(theta/360)))

}
