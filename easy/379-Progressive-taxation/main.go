package main

import (
	"flag"
	"fmt"
)

func main() {
	income := flag.Int("i", 0, "Gives the tax amount owed on a whole-number income amount up to ¤100,000,000.")
	flag.Parse()
	fmt.Printf("tax(%v) => %v\n", *income, tax(*income))
}

func tax(income int) int {
	fmt.Println("income", income)
	if income <= 10000 {
		return 0
	} else if income <= 30000 {
		return case1(income)
	} else if income <= 100000 {
		return case2(income)
	} else {
		return case3(income)
	}
}

func case1(val int) int {
	return int(float32(val-10000) * 0.1)
}

func case2(val int) int {
	return case1(30000) + int(float32(val-30000)*0.25)
}

func case3(val int) int {
	return case2(100000) + int(float32(val-100000)*0.4)
}
