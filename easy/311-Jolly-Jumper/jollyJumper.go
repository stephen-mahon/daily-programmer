package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
)

var arg string

func main() {
	flag.StringVar(&arg, "o", "0", "arg inputs")
	flag.Parse()
	tail := flag.Args()
	v, _ := strconv.Atoi(arg)
	var vals = []int{v}
	for i := range tail {
		v, _ = strconv.Atoi(tail[i])
		vals = append(vals, v)
	}
	var state string
	if isJolly(vals) {
		state = "JOLLY"
	} else {
		state = "NOT JOLLY"
	}
	fmt.Println(vals, state)

}

// see https://www.geeksforgeeks.org/jolly-jumper-sequence/
func isJolly(args []int) bool {
	// Boolean vector to diffSet set of differences. The vector is initialized as false.
	diffSet := make([]bool, len(args))

	// Traverse all array elements
	for i := 0; i < len(args)-1; i++ {
		// Find absolute difference between current two
		d := int(math.Abs(float64(args[i] - args[i+1])))

		// If difference is out of range or repeated, return false.
		if d == 0 || d > len(args)-1 || diffSet[d] == true {
			fmt.Println(diffSet)
			return false
		}

		// Set presence of d in set.
		diffSet[d] = true
	}
	return true
}
