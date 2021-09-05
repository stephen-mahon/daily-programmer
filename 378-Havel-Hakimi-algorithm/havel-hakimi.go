package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "/help" {
		fmt.Println("Havel-Hakimi Algorithm. See readme for more information.")
		fmt.Println("usage: ./havel-hakimi <array of numbers>")
		fmt.Println("example: ./havel-hakimi 5 3 0 2 6 2 0 7 2 5")
		fmt.Println("output: [5 3 0 2 6 2 0 7 2 5] => false")
	} else {
		vals := parseInput(args)
		fmt.Printf("%v => %v", vals, havelHakimi(vals))
	}

}

func parseInput(args []string) []int {
	var vals []int
	for i := range args {
		val, err := strconv.Atoi(string(args[i]))
		if err != nil {
			fmt.Printf("error. you must enter a number. you entered %q.\n", args[i])
			log.Fatal()
		}
		vals = append(vals, val)
	}

	return vals
}

func havelHakimi(vals []int) bool {
	vals = elemZeros(vals)
	if len(vals) == 0 {
		return true
	}
	vals = descOrder(vals)
	maxVal := vals[0]
	vals = vals[1:]
	if lengthCheck(maxVal, vals) {
		return false
	}
	vals = frontElim(maxVal, vals)
	return havelHakimi(vals)
}

func elemZeros(ans []int) []int {

	var vals []int
	for i := range ans {
		if ans[i] != 0 {
			vals = append(vals, ans[i])
		}
	}

	return vals
}

func descOrder(vals []int) []int {
	sort.Slice(vals, func(i, j int) bool {
		return vals[i] > vals[j]
	})
	return vals
}

func lengthCheck(l int, vals []int) bool {

	return l > len(vals)
}

func frontElim(n int, vals []int) []int {
	if n > len(vals) {
		fmt.Println("front elemination problem. n is longer than the sequence.")
		log.Fatal()
	}
	for i := 0; i < n; i++ {
		vals[i] = vals[i] - 1
	}
	return vals

}
