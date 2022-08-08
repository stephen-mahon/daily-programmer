package main

import (
	"flag"
	"fmt"
	"strings"
)

var numbers = map[string]int{
	" _ | ||_|": 0,
	"     |  |": 1,
	" _  _||_ ": 2,
	" _  _| _|": 3,
	"   |_|  |": 4,
	" _ |_  _|": 5,
	" _ |_ |_|": 6,
	" _   |  |": 7,
	" _ |_||_|": 8,
	" _ |_| _|": 9,
}

func main() {
	input := flag.String("d", "-1", "numbers contained in the display")
	flag.Parse()
	output := sevenSegDisplay(*input)
	for _, v := range output {
		fmt.Print(v)
	}
	fmt.Println()
}

func sevenSegDisplay(input string) []int {
	lines := strings.Split(input, "\n")
	var out []int
	for i := 0; i < len(lines[0]); i += 3 {
		num := ""
		for j := range lines {
			num += lines[j][i : i+3]
		}
		out = append(out, numbers[num])
	}
	return out
}
