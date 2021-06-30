package main

// To do
// [ ] fit2
// [ ] fit3
// [ ] fitn

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "/help" {
		fmt.Println("Given a 2-dimensional rectangular crate of size X by Y, and boxes each of size x by y, returns boxes can fit into a single crate if they have to be placed so that the x-axis of the boxes is aligned with the x-axis of the crate, and the y-axis of the boxes is aligned with the y-axis of the crate")
		fmt.Println("Usage: fit1 <X Y x y>")
		fmt.Println("Example: ./fit1 25 18 6 5")
	} else {
		if len(args) != 4 {
			fmt.Println("You must enter four arguments! Type /help for help.")
		} else {
			fmt.Println(fit1(args))
		}
	}
}

func fit1(args []string) int {
	var fitArray []int
	for _, v := range args {
		val, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		fitArray = append(fitArray, val)
	}

	maxX := fitArray[0] / fitArray[2]
	maxY := fitArray[1] / fitArray[3]

	return maxX * maxY
}

func fit2(args []string) int {
	var fitArray []int
	// loop to append a slice of fitArray values with the rotated boxes
	fitArray = append(fitArray, fit1(args))
	max := 0
	for _, v := range fitArray {
		if v > max {
			v = max
		}
	}
	return max
}
