package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {
	args := flag.String("w", "", "word")
	flag.Parse()
	words := strings.Split(*args, " ")
	if len(words) != 2 {
		log.Fatalf("input error. more than two word: %q", *args)
	}
	fmt.Printf("funnel(%q, %q) => %v\n", words[0], words[1], funnel(words))
}

func funnel(words []string) bool {
	input, test := words[0], words[1]
	if len(test) != len(input)-1 {
		return false
	}
	for i := 1; i < len(input); i++ {
		check := input[0:i-1] + input[i:]
		if test == check {
			return true
		}
	}
	return false
}
