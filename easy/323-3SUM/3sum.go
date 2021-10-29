package main

import (
	"fmt"
	"sort"
)

var title = "3SUM"
var inputExamples = []int{9, -6, -5, 9, 8, 3, -4, 8, 1, 7, -4, 9, -9, 1, 9, -9, 9, 4, -6, -8}

func main() {
	fmt.Println(title)
	sort.Ints(inputExamples)
	for _, v := range inputExamples {
		fmt.Println(v)
	}
}
