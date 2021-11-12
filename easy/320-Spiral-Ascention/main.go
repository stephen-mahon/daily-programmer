package main

import (
	"flag"
	"fmt"
)

var title = "Spiral Ascention"

func main() {
	fmt.Println(title)
	square := flag.Int("square", 3, "The length of the square for the spiral")
	flag.Parse()

	array := makeArray(square)
	traverseArray(array)

}

func makeArray(square *int) [][]string {
	var array [][]string
	for i := 0; i < *square; i++ {
		tmp := []string{}
		for j := 0; j < *square; j++ {
			tmp = append(tmp, "X")
		}
		array = append(array, tmp)
	}
	return array
}

func traverseArray(array [][]string) {
	for i := range array {
		for j := range array[i] {
			fmt.Printf("(%v,%v) ", i, j)
		}
		fmt.Println()
	}
}
