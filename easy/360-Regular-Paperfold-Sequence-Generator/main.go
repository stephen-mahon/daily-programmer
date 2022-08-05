package main

import (
	"flag"
	"fmt"
)

const seed = "1"

func main() {
	cycles := flag.Int("c", 2, "number of cycles")
	flag.Parse()
	fmt.Println(dragon(seed, *cycles))
}

func dragon(seed string, num int) string {
	for num > 0 {
		fmt.Printf("cycle #%v: %v -> ", num, seed)
		seed = seed + seed + middleComplemented(seed)
		fmt.Printf("%v\n", seed)
		dragon(seed, num-1)
	}
	return seed
}

func middleComplemented(digits string) string {
	middle := digits[(len(digits)-1)/2]
	if digits[middle] == '1' {
		digits = digits[0:middle] + "0" + digits[middle:]
	} else {
		digits = digits[0:middle] + "1" + digits[middle:]
	}
	return digits
}
