package main

// To do
// [X] fit2
// [ ] fit3
// [ ] fitn

import (
	"fmt"
	"os"
	"strconv"
)

type rectangle struct {
	x, y int
}

type axes struct{
	x,y,z int
}

type cuboid struct {
	x, y, z int
}

func main() {
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "/help" {
		//update help with fit2
		//fmt.Println("Given a 2-dimensional rectangular crate of size X by Y, and boxes each of size x by y, returns boxes can fit into a single crate if they have to be placed so that the x-axis of the boxes is aligned with the x-axis of the crate, and the y-axis of the boxes is aligned with the y-axis of the crate")
		fmt.Println("Usage: fit1 <X Y x y>")
		fmt.Println("Example: ./fit1 25 18 6 5")
	} else {
		if len(args) != 4 {
			fmt.Println("You must enter four arguments! Type /help for help.")
		} else {
			args := stringToInt(args)
			crate, box := crateAndBox(args)
			fmt.Println("fit1:", fit1(crate, box))
			fmt.Println("fit2:", fit2(crate, box))
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

func rotateBox(box rectangle) rectangle {
	box.x = box.y
	box.y = box.x
	return box
}

func rotateCuboid(box cuboid, index string, dir int) cuboid {
	return box
}

func crateAndBox(args []int) (rectangle, rectangle) {
	crate := rectangle{args[0], args[1]}
	box := rectangle{args[2], args[3]}
	return crate, box
}

func fit1(crate, box rectangle) int {
	maxX := crate.x / box.x
	maxY := crate.y / box.y
	return maxX * maxY
}

func fit2(crate, box rectangle) int {
	var fitArray []int
	fitArray = append(fitArray, fit1(crate, box))
	fitArray = append(fitArray, fit1(crate, rotateBox(box)))
	return maxVal(fitArray)
}

func fit3(crate, box cuboid) int {
	var fitArray []int

	return maxVal(fitArray)

}
