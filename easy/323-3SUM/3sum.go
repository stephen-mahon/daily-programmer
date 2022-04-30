package main

import (
	"fmt"
	"math/bits"
)

// still need to extract unique combinations in set

var title = "3SUM"
var inputExamples = []int{9, -6, -5, 9, 8, 3, -4, 8, 1, 7, -4, 9, -9, 1, 9, -9, 9, 4, -6, -8}

func main() {
	fmt.Println(title)
	vals := combinations(inputExamples, 3)
	for i := range vals {
		if vals[i][0]+vals[i][1]+vals[i][2] == 0 {
			fmt.Println(vals[i])
		}
	}
}

// Combinations returns combinations of n elements for a given string array.
// For n < 1, it equals to All and returns all combinations.
func combinations(set []int, n int) (subsets [][]int) {
	length := uint(len(set))

	if n > len(set) {
		n = len(set)
	}

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		if n > 0 && bits.OnesCount(uint(subsetBits)) != n {
			continue
		}

		var subset []int

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, set[object])
			}
		}
		// add subset to subsets
		subsets = append(subsets, subset)
	}
	return subsets
}
