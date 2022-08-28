package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

var arg string

func main() {
	flag.StringVar(&arg, "o", "0", "arg inputs")
	flag.Parse()
	tail := flag.Args()
	input, err := parseInput(arg, tail)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(input, "->", simpleSubsetSum(input))

}

func parseInput(head string, tail []string) ([]int, error) {
	var vals []int
	val, err := strconv.Atoi(head)
	if err != nil {
		return []int{}, err
	}
	vals = append(vals, val)
	for _, v := range tail {
		val, err := strconv.Atoi(v)
		if err != nil {
			return []int{}, err
		}
		vals = append(vals, val)
	}
	return vals, nil
}

func simpleSubsetSum(vals []int) bool {
	for i, v := range vals {
		if v == 0 {
			return true
		}
		for j, u := range vals {
			if i != j && v+u == 0 {
				return true
			}
		}
	}
	return false
}
