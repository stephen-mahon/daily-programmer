package main

import (
	"fmt"
	"os"
)

func letterSum(s string) int {
	var total int
	for _, v := range s {
		total += int(v - 96)
	}
	return total
}

func main() {
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "/help" {
		fmt.Println("Letter value sum. Given a string of lowercase letters, finds the sum of the values of the letters in the string. 1 for t0 26 for z.")
		fmt.Println("Usage: lettersum <string>")
		fmt.Println("Example: ./lettersum golang")
	} else {
		if len(args) != 1 {
			fmt.Println("You must enter one argument! Type /help for help.")
		} else {
			fmt.Println(letterSum(args[0]))
		}
	}
}
