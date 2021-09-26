package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var title = "Lucky numbers"
var err = "See readme"
var help = ""

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println(err)
	} else if len(args) == 1 && args[0] == "-help" {
		fmt.Println(title)
		fmt.Printf(help)
	} else if len(args) != 1 {
		fmt.Println(err)
	} else {
		vals := numSet(args)
		luckyNums(vals)
	}
}

func numSet(args []string) []int {
	limit, err := strconv.Atoi(args[0])
	var setOfNum []int
	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i <= limit; i++ {
		setOfNum = append(setOfNum, i)

	}
	return setOfNum
}

func luckyNums(set []int) []int {

	fmt.Println(set, set[1])

	if set[1] > len(set) {
		return set
	}

	var vals []int
	for i, v := range set {
		if i%(set[1]) != 0 {
			vals = append(vals, v)
		}

	}

	return luckyNums(vals)
}
