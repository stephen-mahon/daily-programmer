package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

func main() {
	num := flag.String("upc", "042100005264", "UPS number")
	flag.Parse()
	if len(*num) > 11 {
		log.Fatal("invalid upc value entered")
	}
	fmt.Printf("ups(%v) => %v\n", *num, upc(*num))
}

func upc(val string) int {
	input := val
	for i := 0; i < 11-len(val); i++ {
		input = "0" + input
	}

	var sumEven, sumOdd int

	for i := 0; i < len(input); i++ {
		v, err := strconv.Atoi(string(input[i]))
		if err != nil {
			log.Fatal("error with parse: ", err)
		}
		if i%2 == 0 {
			sumEven += v
		} else {
			sumOdd += v
		}
	}

	lastDigit := ((3 * sumEven) + sumOdd) % 10
	if lastDigit != 0 {
		return 10 - lastDigit

	}
	return 0
}
