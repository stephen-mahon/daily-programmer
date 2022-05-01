package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "-help" || args[0] == "-h" {
		fmt.Println("Making change")
		fmt.Printf("usage:   ./making-change <int>\n")
		fmt.Printf("example: ./making-change 123456\n")
		fmt.Printf("output:  254\n")
	} else {
		if len(args) != 1 {
			fmt.Printf("You must enter one argument! Type \"-help\" for help.\n")
		} else {
			res, _ := makeChange(args[0])
			fmt.Println(res)
		}
	}
}

func makeChange(arg string) (int, error) {
	val, err := strconv.Atoi(arg)
	if err != nil {
		return -1, err
	}
	coins := [6]int{500, 100, 25, 10, 5, 1}
	var change int
	for i := range coins {
		res := val / coins[i]
		change += res
		val -= res * coins[i]
	}
	return change, nil
}
