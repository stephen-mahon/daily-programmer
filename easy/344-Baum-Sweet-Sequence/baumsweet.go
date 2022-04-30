package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "\\help" {
		fmt.Printf("Baum-Sweet Sequence.\nhttps://en.wikipedia.org/wiki/Baum-Sweet_sequence\n")
	} else {
		if len(args) != 1 {
			fmt.Printf("You must enter one argument! Type \"/help\" for help.\n")
		} else {
			sequence(args[0])
		}
	}
}

func sequence(arg string) {
	iter, err := strconv.Atoi(string(arg))
	if err != nil {
		fmt.Printf("an error occured. expecting an int\nplease try again\n")
		log.Panic(err)
	}
	for i := 0; i <= iter; i++ {
		val := baumsweet(i)
		if i != iter {
			fmt.Printf("%v, ", val)
		} else {
			fmt.Printf("%v\n", val)
		}
	}
}

func baumsweet(val int) int {
	b := fmt.Sprintf("%b", val)
	var count int
	var bsArray []int
	for _, v := range b {
		if v == '0' {
			count++
		} else {
			bsArray = append(bsArray, count)
			count = 0
		}

	}
	if count != 0 {
		bsArray = append(bsArray, count)
	}
	for _, v := range bsArray {
		if v%2 != 0 && len(bsArray) > 1 {
			return 0
		}
	}
	return 1
}
