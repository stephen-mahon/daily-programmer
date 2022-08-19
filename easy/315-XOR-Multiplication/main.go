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
	vals, err := parseInt(*input)
	if err != nil {
		log.Fatalf("unexpected input! - %v", err)
	}
	if len(vals) != 2 {
		log.Fatalf("too many arguments! - %v", *input)
	}
	fmt.Printf("%v@%v = %v\n", vals[0], vals[1], xorProduct(vals))

}

func parseInt(args string) (vals []int, err error) {
	input := strings.Split(args, " ")
	for _, v := range input {
		val, err := strconv.Atoi(v)
		if err != nil {
			return []int{}, err
		}
		vals = append(vals, val)
	}
	return vals, nil
}

func xorProduct(vals []int) int {
	var bin []string
	for _, v := range vals {
		bin = append(bin, strconv.FormatInt(int64(v), 2))
	}
	return -1
}
