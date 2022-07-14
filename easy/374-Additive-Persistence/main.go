package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	arg := flag.Int("n", 0, "enter a digit")
	flag.Parse()
	val := addNums(*arg, 0)
	fmt.Println(*arg, "->", val)
}

func addNums(num int, iter int) int {
	intString := strconv.Itoa(num)
	if len(intString) == 1 {
		return iter
	}
	var tmp int
	for _, v := range intString {
		val, _ := strconv.Atoi(string(v))
		tmp += val
	}
	return addNums(tmp, iter+1)
}
