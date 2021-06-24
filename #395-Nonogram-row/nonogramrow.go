package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "/help" {
		fmt.Println("Given a binary array of any length, return an array of positive integers that represent the lengths of the sets of consecutive 1's in the input array, in order from left to right.")
		fmt.Println("Usage: nonogramrow <binary array>")
		fmt.Println("Example: ./nonogramrow 1 0 1 1")
	} else {
		fmt.Println(nonogramrow(args))

	}
}

func nonogramrow(binaryArray []string) []int {
	var nongramArray []int
	var count int
	for _, v := range binaryArray {
		if v == "1" {
			count++
		} else {
			if count != 0 {
				nongramArray = append(nongramArray, count)
				count = 0
			}
			count = 0
		}
	}
	if count != 0 {
		nongramArray = append(nongramArray, count)
	}
	return nongramArray
}
