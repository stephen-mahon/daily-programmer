package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

// To do
//	[ ] sort out the bug with casting large ints to int32

func main() {
	args := flag.String("a", "", "a list of integers")
	flag.Parse()
	max, min, err := concatInt(*args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(max, min)
}

/// works fine for small numbers but the casting to int64 is causing some issues with the concat
func concatInt(args string) (max, min int64, err error) {
	input := strings.Split(args, " ")
	max = math.MinInt32
	min = math.MaxInt32
	for i := range input {
		for j := range input {
			for k := range input {
				if i != j && j != k && k != i {
					// I think that here we are getting numbers larger than ints
					v := input[i] + input[j] + input[k]
					val, err := strconv.Atoi(v)
					if err != nil {
						return 0, 0, err
					}
					// I don't think that this casting is working well
					if int64(val) > max {
						// Here neither
						max = int64(val)
					}
					// And here
					if int64(val) < min {
						// finally, here too
						min = int64(val)
					}
				}
			}
		}
	}
	return max, min, nil
}
