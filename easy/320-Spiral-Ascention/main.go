package main

import (
	"flag"
	"fmt"
)

var title = "Spiral Ascention"

type coor struct {
	x int
	y int
}

func main() {
	fmt.Println(title)
	square := flag.Int("square", 3, "The length of the square for the spiral")
	flag.Parse()

	array := makeArray(square)
	printArray(square, array)
	traverseArray(array)

}

func makeArray(square *int) []coor {
	var array []coor

	for i := 1; i <= *square; i++ {
		for j := 1; j <= *square; j++ {
			array = append(array, coor{i, j})
		}
	}

	return array
}

func printArray(square *int, array []coor) {
	for i := 1; i <= len(array); i++ {
		if i != 0 && i%*square == 0 {
			fmt.Printf("(%v,%v)\n", array[i-1].x, array[i-1].y)
		} else {
			fmt.Printf("(%v,%v) ", array[i-1].x, array[i-1].y)
		}
	}
}

func traverseArray(array []coor) {
	visit := make([]bool, len(array))
	fmt.Println(visit)
	for i := range array {
		if !visit[i] {
			fmt.Println(array[i].x, array[i].y)
			visit[i] = true
		}

	}
	fmt.Println(visit)
}
