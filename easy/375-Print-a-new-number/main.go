package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

func main() {
	arg := flag.Int("n", 0, "enter a number")
	flag.Parse()
	val, err := addOnetoEachDigit(*arg)
	if err != nil {
		log.Panicf("error with return: %v", err)
	}
	fmt.Printf("OneToNum(%v) -> %v\n", *arg, val)
}

func addOnetoEachDigit(num int) (int, error) {
	arg := strconv.Itoa(num)
	var numStr string
	for _, vStr := range arg {
		vInt, err := strconv.Atoi(string(vStr))
		if err != nil {
			return -1, err
		}
		tmp := strconv.Itoa(vInt + 1)
		numStr += tmp
	}
	return strconv.Atoi(numStr)
}
