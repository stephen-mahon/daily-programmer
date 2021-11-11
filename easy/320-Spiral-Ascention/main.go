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
	val := 1
	for i := 0; i < *square; i++ {
		for j := 0; j < *square; j++ {
			fmt.Printf("%v ", val)
			val++
		}
		fmt.Println()
	}
}
