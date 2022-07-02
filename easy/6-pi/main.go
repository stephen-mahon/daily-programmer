package main

import (
	"fmt"
	"math/big"
)

//big math

func main() {
	// 200 bits of percision
	const prec = 200

	// initial value of pi set to zero
	pi := new(big.Float).SetPrec(prec).SetInt64(0)

	// Leibniz requires and alternating sign summation
	sign := false

	// initialize values for the compution. Liebniz gives pi/4, so four is needed
	four := new(big.Float).SetPrec(prec).SetInt64(4)

	// use t as a temporary value
	t := new(big.Float)

	// iterate 1 - 1/3 + 1/5 - 1/7 + ... = pi/4
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

	// Leibniz converges extremely slowly. Calculating pi to ten places requires about 5 billion terms.
	fmt.Printf("pi = %.30f\n", pi)

	// next step is convergence acceleration techniques.
}
