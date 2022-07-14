package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	arg := flag.Int("n", 0, "enter a digit")
	flag.Parse()

	var acc int
	addDigits(*arg, &acc)

	fmt.Println(*arg, "->", acc)
}

func addDigits(num int, iter *int) int {
	intString := strconv.Itoa(num)
	if len(intString) == 1 {
		return *iter
	}
	var tmp int
	for _, v := range intString {
		val, _ := strconv.Atoi(string(v))
		tmp += val
	}
	*iter += 1
	return addDigits(tmp, iter)
}
