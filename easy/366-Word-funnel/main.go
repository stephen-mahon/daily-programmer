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
	return false
}
