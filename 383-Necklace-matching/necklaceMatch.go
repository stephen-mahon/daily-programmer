package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "-help" {
		fmt.Println("Same Necklace.")
		fmt.Println("Example: ./sameNecklace nicole icolen")
		fmt.Println("Output: true")
	} else {
		if len(args) != 2 {
			fmt.Println("You must enter two arguments! Type -help for help.")
		} else {
			fmt.Println(sameNecklace(args))
		}
	}
}

func sameNecklace(args []string) bool {
	word := args[0]
	check := args[1]
	for i := range word {
		if check == word[i:]+word[:i] {
			return true
		}
	}
	return false
}
