package main

import (
	"flag"
	"fmt"
)

func main() {
	arg := flag.String("n", "", "value input")
	flag.Parse()

	fmt.Printf("balanced(%v) => %v\n", *arg, balanced(*arg))
}

func balanced(val string) bool {
	if val[0] == val[len(val)-1] {
		balanced(val[1 : len(val)-1])
	} else {
		return false
	}
	return true
}
