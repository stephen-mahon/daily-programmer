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
	//fmt.Printf("%v\n", seed)
	if num == 0 {
		return seed
	} else {
		seed = seed + "1" + middleComplemented(seed)
		seed = dragon(seed, num-1)
	}
	return seed
}

func middleComplemented(digits string) string {
	middle := (len(digits) - 1) / 2
	if len(digits) == 1 {
		if digits[0] == '1' {
			return "0"
		} else {
			return "1"
		}
	}
	if digits[middle] == '1' {
		digits = digits[0:middle] + "0" + digits[middle+1:]
	} else {
		digits = digits[0:middle] + "1" + digits[middle:]
	}
	return digits
}
