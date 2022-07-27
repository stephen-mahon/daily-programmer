package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	args := flag.String("c", "0 0 0", "rgb values")
	flag.Parse()
	rgb, err := parseNums(args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("hexcolor(%v, %v, %v) => \"#%02x%02x%02x\"", rgb[0], rgb[1], rgb[2], rgb[0], rgb[1], rgb[2])
}

func parseNums(strInput *string) ([]int, error) {
	strSplit := strings.Split(*strInput, " ")
	if len(strSplit) != 3 {
		return []int{}, fmt.Errorf("invalid input rgb value must be 3 numbers: %s", *strInput)
	}
	var vals []int
	for i := 0; i < 3; i++ {
		val, err := strconv.Atoi(strSplit[i])
		if err != nil {
			return []int{}, fmt.Errorf("invalid input rgb value must be a numbers: %s", strSplit[i])
		}
		if val < 0 || val > 255 {
			return []int{}, fmt.Errorf("invalid input rgb value must be between 0 and 255: %v", val)
		}
		vals = append(vals, val)
	}
	return vals, nil
}
