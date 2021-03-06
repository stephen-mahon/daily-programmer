package main

// To do
// [ ] fitN
// [ ] add check for negative integer input

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type axes struct {
	x, y, z int
}

type cuboid struct {
	x, y, z int
}

func main() {
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "/help" {
		fmt.Println("Given the X, Y, and Z dimensions of a crate, and the x, y, and z dimensions of a box, returns the max number of boxes that can fit into a single crate.")
		fmt.Println("Usage: ./fit3 <X Y Z x y z>")
		fmt.Println("Example: ./fit3 12 34 56 7 8 9")
	} else {
		if len(args) != 6 {
			fmt.Println("You must enter six arguments! Type /help for help.")
		} else {
			args := stringToInt(args)
			crate, box := crateAndBox3D(args)
			fmt.Println(fit3D(crate, box))
		}
	}
}

// takes in the input tags and converts to ints. Will panic exit with error if an input cannot be converted to int.
// N.B. no check for negative numbers
func stringToInt(args []string) []int {
	var argInt []int
	for _, v := range args {
		val, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		argInt = append(argInt, val)
	}
	return argInt
}

// returns the max int in a slice of ints
func maxVal(vals []int) int {
	max := 0
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max
}

// rotates the cuboid by switching the other axes values
func (box *cuboid) rotate(axis string) {
	axis = strings.ToLower(axis)
	var a, b int
	switch axis {
	case "x":
		// There is clever way to solve this using pointers and addresses but vars a and b do the job for now.
		a = box.y
		b = box.z
		box.y = b
		box.z = a
	case "y":
		a = box.x
		b = box.z
		box.x = b
		box.z = a
	default:
		a = box.x
		b = box.y
		box.x = b
		box.y = a
	}
}

// creates a crate and box as a cuboid type from the arg input
func crateAndBox3D(args []int) (cuboid, cuboid) {
	crate := cuboid{args[0], args[1], args[2]}
	box := cuboid{args[3], args[4], args[5]}
	return crate, box
}

// for debugging the rotate method
func (box cuboid) printVal() {
	fmt.Printf("(%d,%d,%d)\n", box.x, box.y, box.z)
}

// basic fit function. Only works with int due to the floating point round off.
func fit(crate, box cuboid) int {
	maxX := crate.x / box.x
	maxY := crate.y / box.y
	maxZ := crate.z / box.z
	return maxX * maxY * maxZ
}

// looks at all the orientations of the box in the crate and returns the max number.
func fit3D(crate, box cuboid) int {
	var fitArray []int
	// This slice cycles through the six different possible orientations of the boxes
	axes := []string{"x", "y", "z", "x", "y", "z"}
	for _, v := range axes {
		box.rotate(v)
		fitArray = append(fitArray, fit(crate, box))
	}
	return maxVal(fitArray)
}
