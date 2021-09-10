package main

import (
	"fmt"
	"math"
)

var message = "Minimize and Maximize"

// golden ratio for golden-section search algorithm
var gr = (math.Sqrt(5) + 1) / 2

// tolerence for golden-section search algorithm
var tol = 1e-1

func main() {
	fmt.Printf("%v\n\n", message)
	fmt.Printf("Challenge 1\nA piece of wire 100 cm in length is bent into the shape of a sector of a circle. Find the angle of the sector that maximizes the area A of the sector.\n")
	fmt.Printf("%.1f\n\n", gss(f1, "max", 0, 180, tol))
	fmt.Printf("Challenge 2\nA and B are towns 20km and 30km from a straight stretch of river 100km long. Water is pumped from a point P on the river by pipelines to both towns. Where should P be located to minimize the total length of pipe AP + PB?\n")
	fmt.Printf("%.1f\n\n", gss(f2, "min", 0, 100, tol))
}

// Golden-section search
/*
	to find the minimum or maximum of f on [a, b]
	f: a strictly unimodal function on [a,b]
	condition: minimize or maximize f
	a: the left bound of the function
	b: the right bound of the function
	tol: the precision of the output
*/
func gss(f func(float64) float64, condition string, a, b, tol float64) float64 {
	c := b - (b-a)/gr
	d := a + (b-a)/gr

	switch condition {
	case "max":
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
	case "min":
		count := 0
		for math.Abs(b-a) > tol {
			count++
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
	return -1
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
func f1(theta float64) float64 {
	return ((theta / 360) * math.Pi * 10000.0) / (4.0 * (1 + 2.0*math.Pi*(theta/360) + math.Pi*(theta/360)*math.Pi*(theta/360)))
}

/*
	A is 20 km from xy at point x and B is 30 km from xy at point y. The point P is somewhere along xy. So Two triangles: AxP and ByP.
		|Ax|^2 + |xP|^2 = |AP|^2, and |By|^2 + |yP|^2 = |BP|^2
		20^2 + (xP)^2 = AP^2
		30^2 + (yP)^2 = BP^2
	Want to minimize AP + BP
	->	[AP + BP]_min = [sqrt(20^2 + (xP)^2) + sqrt(30^2 + (yP)^2)]_min
	|xy| = 100 km.
		|xP| + |Py| = 100.
	yP = Py. In terms of yP
		yP = 100 - xP
	Substitute into the minimize function
	->	[AP + BP]_min = [sqrt(20^2 + (xP)^2) + sqrt(30^2 + (100 - xP)^2)]_min
*/
func f2(xP float64) float64 {
	return math.Sqrt((20*20)+xP*xP) + math.Sqrt((30*30)+(100-xP)*(100-xP))
}
