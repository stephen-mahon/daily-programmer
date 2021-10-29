package main

import (
	"fmt"
)

var title = "3SUM"
var inputExamples = []int{9, -6, -5, 9, 8, 3, -4, 8, 1, 7, -4, 9, -9, 1, 9, -9, 9, 4, -6, -8}

func main() {
	fmt.Println(title)
	//sort.Ints(inputExamples)
	combinations(inputExamples)
}

func combinations(vals []int) [][]int {

	//var triples [][3]int

	for i := range vals {
		for j := range vals {
			for k := range vals {
				if vals[i]+vals[j]+vals[k] == 0 {
					fmt.Println(vals[i], vals[j], vals[k])
				}
			}
		}
	}

	return [][]int{}
}
