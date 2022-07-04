package main

import (
	"math/big"
)

//big math

// 200 bits of percision
const prec = 200

func main() {
	// initial value of pi set to zero
	pi := new(big.Float).SetPrec(prec).SetInt64(0)

	// Leibniz converges extremely slowly. Calculating pi to ten places requires about 5 billion terms.
	// fmt.Printf("pi = %.30f\n", lSeries(pi))

	// next step is convergence acceleration techniques.
	glSeries(pi)
}

// Liebniz Series
func lSeries(pi *big.Float) *big.Float {
	// Leibniz requires and alternating sign summation
	sign := false

	// initialize values for the compution. Liebniz gives pi/4, so four is needed
	four := new(big.Float).SetPrec(prec).SetInt64(4)

	// use t as a temporary value
	t := new(big.Float)

	// iterate 1 - 1/3 + 1/5 - 1/7 + ... = pi/4
	//    pi = 4 - 4/3 + 4/5 - 4/7 + ...
	for i := 1; i < 1000000; i += 2 {
		sign = !sign
		n := new(big.Float).SetPrec(prec).SetInt64(int64(i))
		t.Quo(four, n)
		if sign {
			pi.Add(pi, t)
		} else {
			pi.Sub(pi, t)
		}
	}

	return pi
}

// https://crypto.stanford.edu/pbc/notes/pi/glseries.html
// Gregory-Leibniz Series
// tan-1(x) = x - x^3/3 + x^5/5 - ...
// let x = 1/sqrt(3)
// tan-1(1/(sqrt(3))) = 1/sqrt(3)(1 - 1/(3*3) + 1/(5*3^2) - 1/(7*3^3) + ...) = pi/6
// Better to get rid of the sqrt(3) term using
// tan-1(1) = tan-1(1/2) + tan-1(1/3)
//			= 1/2(1 - 1/(3*2^2) + 1/(5*2^4) - ...) +
//			  1/3(1 - 1/(3*3^2) + 1/(5*3^4) - ...)
func glSeries(pi *big.Float) *big.Float {
	sign := false
	t1 := new(big.Float)
	t2 := new(big.Float)

	one := new(big.Float).SetPrec(prec).SetFloat64(1.0)
	oneHalf := new(big.Float).SetPrec(prec).SetFloat64(0.5)
	oneThird := new(big.Float).SetPrec(prec).SetFloat64(1.0 / 3.0)

	for i := 1; i < 1000000; i += 2 {
		sign = !sign
		n1 := new(big.Float).SetPrec(prec).SetInt64(int64(i))
		n2 := new(big.Float).SetPrec(prec).SetInt64(int64(2 ^ (i - 1)))
		n3 := new(big.Float).SetPrec(prec).SetInt64(int64(3 ^ (i - 1)))
		t1.Mul(n1, n2)
		t1.Quo(one, t1)
		t1.Mul(oneHalf, t1)
		t2.Mul(n1, n3)
		t2.Quo(one, t2)
		t2.Mul(oneThird, t2)
		if sign {
			t1.Add(t1, t1)
			t2.Add(t2, t2)
		} else {
			t1.Sub(t1, t1)
			t2.Sub(t2, t2)
		}
		pi.Add(t1, t2)
	}
	four := new(big.Float).SetPrec(prec).SetFloat64(4.0)
	return pi.Mul(four, pi)
}
