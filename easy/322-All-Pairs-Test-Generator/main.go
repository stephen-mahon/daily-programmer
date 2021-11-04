package main

import (
	"fmt"
	"math/bits"
)

var title = "All pairs test generator"

var testInput = [][]string{{"0", "1"}, {"A", "B", "C"}, {"D", "E", "F", "G"}}

func main() {
	fmt.Println(title)
	printSet(testInput)
	fmt.Printf("\nTotal Combinations: %v", totalCombinations(testInput))
}

func printSet(input [][]string) {
	for i := range input {
		fmt.Printf("%v ", input[i])
	}
}

func totalCombinations(input [][]string) int {
	totalCom := 1
	for i := range input {
		totalCom *= len(input[i])
	}
	return totalCom
}

// Combinations returns combinations of n elements for a given string array.
// For n < 1, it equals to All and returns all combinations.
func combinations(set []string, n int) (subsets [][]string) {
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

		var subset []string

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
