package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	input := flag.String("v", "1 2", "XOR-product")
	flag.Parse()
	vals, err := parseUint(*input)
	if err != nil {
		log.Fatalf("unexpected input! - %v", err)
	}
	if len(vals) != 2 {
		log.Fatalf("too many arguments! - %v", *input)
	}
	fmt.Printf("%v@%v = %v\n", vals[0], vals[1], xorProduct(vals[0], vals[1]))
}

func parseUint(args string) (vals []uint, err error) {
	input := strings.Split(args, " ")
	for _, v := range input {
		val, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return []uint{}, err
		}
		vals = append(vals, uint(val))
	}
	return vals, nil
}

func xorProduct(a, b uint) uint {
	// xor truth table
	// A B | XOR
	// 0 0 |  0
	// 0 1 |  1
	// 1 0 |  1
	// 1 1 |  0
	// A(!B) + (!A)B
	var product uint

	// bitshift operation
	// While b is not zero, left shifted a and right shift b
	// b 1 | AND
	// 0 1 | 0
	// 1 1 | 1Ad
	for ; b != 0; a, b = a<<1, b>>1 {
		product ^= (a * (b & 1))
	}
	return product
}
