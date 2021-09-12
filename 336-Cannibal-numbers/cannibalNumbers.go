package main

/* To do
[ ] consume func
*/

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

var message = "Cannibal numbers."
var maxNum = 15

func main() {
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "-help" {
		fmt.Println(message)
	} else {
		if len(args) != 2 {
			fmt.Println("You must enter two arguments! Type -help for help.")
		} else {
			args := input(args)
			vals, queries := genArray(args)
			fmt.Printf("%v, %v\n", vals, queries)
			fmt.Println(cannibalNums(vals, queries))
		}
	}
}

func input(args []string) []int {
	var outputs []int
	for i := range args {
		val, err := strconv.Atoi(args[i])
		if err != nil {
			fmt.Printf("error, expecting a number but recieved %q\n", args)
			log.Fatal()
		}
		outputs = append(outputs, val)
	}
	return outputs
}

func genArray(args []int) ([]int, []int) {
	var vals, queries []int
	for i := 0; i < args[0]; i++ {
		vals = append(vals, rand.Intn(maxNum))
	}
	for i := 0; i < args[1]; i++ {
		queries = append(queries, rand.Intn(maxNum))
	}
	return vals, queries
}

func cannibalNums(vals, queries []int) []int {
	var output []int
	for i := range queries {
		var cannibals, prey []int
		for j := range vals {
			if vals[j] > queries[i] {
				cannibals = append(cannibals, vals[j])
			} else {
				prey = append(prey, vals[j])
			}
		}
		var max int
		max, prey = maxAndArray(prey)
		cannibals, _ = consume(cannibals, prey, queries[i], max)

		output = append(output, len(cannibals))
	}
	return output
}

func maxAndArray(array []int) (int, []int) {
	if len(array) == 0 {
		return 0, array
	}
	var max, index int
	for i, v := range array {
		if v > max {
			max = v
			index = i
		}
	}
	return max, append(array[:index], array[index+1:]...)
}

func removeSmallest(array []int) []int {
	var min, index int
	for i, v := range array {
		if v < min {
			min = v
			index = i
		}
	}
	return append(array[:index], array[index+1:]...)
}

// This is clearly not working
func consume(cannibals, prey []int, query, max int) ([]int, []int) {
	if len(prey) == 0 {
		return cannibals, nil
	}
	if max == query {
		cannibals = append(cannibals, max)
		_, prey = maxAndArray(prey)
		return cannibals, prey
	} else {
		max++
		prey = removeSmallest(prey)

	}
	return consume(cannibals, prey, query, max)
}
