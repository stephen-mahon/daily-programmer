package main

// To do
// [X] fit2
// [ ] fit3
// [ ] fitn

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
		fmt.Println("Given the X, Y, and Z dimensions of the crate, and the x, y, and z dimensions of the boxes, returns the number of boxes that can fit into a single crate.")
		fmt.Println("Usage: fit3 <X Y Z x y z>")
		fmt.Println("Example: ./fit1 12 34 56 7 8 9")
	} else {
		if len(args) != 6 {
			fmt.Println("You must enter six arguments! Type /help for help.")
		} else {
			args := stringToInt(args)
			crate, box := crateAndBox3D(args)
			fmt.Println("fit3:", fit3(crate, box))
		}
	}
}

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

func maxVal(vals []int) int {
	max := 0
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max
}

func checkAxes(axes []string, check string) bool {
	for _, v := range axes {
		if v == check {
			return true
		}
	}
	return false
}

func rotateBox(box cuboid, axes string) cuboid {
	axes = strings.ToLower(axes)
	switch axes {
	case "x":
		box.y = box.z
		box.z = box.y
		return box
	case "y":
		box.x = box.z
		box.z = box.x
		return box
	default:
		box.x = box.y
		box.y = box.x
		return box
	}
}

func crateAndBox3D(args []int) (cuboid, cuboid) {
	crate := cuboid{args[0], args[1], args[2]}
	box := cuboid{args[3], args[4], args[5]}
	return crate, box
}

func printVal(box cuboid) {
	fmt.Printf("(%d,%d,%d)\n", box.x, box.y, box.z)
}

func fit(crate, box cuboid) int {
	maxX := crate.x / box.x
	maxY := crate.y / box.y
	maxZ := crate.z / box.z
	return maxX * maxY * maxZ
}

func fit3(crate, box cuboid) int {
	var fitArray []int
	fmt.Println("Initial")
	printVal(box)
	fitArray = append(fitArray, fit(crate, box))
	box = rotateBox(box, "x")
	fmt.Println("Rotate x")
	printVal(box)
	fitArray = append(fitArray, fit(crate, box))
	box = rotateBox(box, "y")
	fmt.Println("Rotate y")
	printVal(box)
	fitArray = append(fitArray, fit(crate, box))
	box = rotateBox(box, "z")
	fmt.Println("Rotate z")
	printVal(box)

	fitArray = append(fitArray, fit(crate, box))

	return maxVal(fitArray)

}
