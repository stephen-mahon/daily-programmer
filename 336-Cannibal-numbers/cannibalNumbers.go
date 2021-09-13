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
	"time"
)

var message = "Cannibal numbers."
var help = "See README.\nusage:\tcannibalNumbers <int> <int>\nexample:cannibalNumbers 7 2\n\t[21 9 5 8 10 1 3], [10 15]\n\t[4 2]"
var err = "You must enter two arguments! Type -help for help."
var maxNum = 20

func main() {
	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "-help" {
		fmt.Println(message)
		fmt.Printf(help)
	} else {
		if len(args) != 2 {
			fmt.Println(err)
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
			if vals[j] >= queries[i] {
				cannibals = append(cannibals, vals[j])
			} else {
				prey = append(prey, vals[j])
			}

		}

		for prey != nil {
			cannibals, prey = consume(cannibals, prey, queries[i])
		}

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
	var index int
	min := array[0]
	for i, v := range array {
		if v < min {
			min = v
			index = i
		}
	}
	return append(array[:index], array[index+1:]...)
}

func consume(cannibals, prey []int, query int) ([]int, []int) {
	max, prey := maxAndArray(prey)

	if max == query {
		cannibals = append(cannibals, max)
		return cannibals, prey
	}

	if len(prey) == 0 {
		return cannibals, nil
	}

	max++
	prey = removeSmallest(prey)

	return consume(cannibals, append(prey, max), query)
}
