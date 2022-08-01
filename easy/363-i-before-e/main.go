package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	word := flag.String("w", "", "word to be checked")
	flag.Parse()
	fmt.Printf("check(\"%s\") => %v\n", *word, check(*word))
}

func check(arg string) bool {
	return strings.Contains(arg, "cie") || (strings.Contains(arg, "ei") && !strings.Contains(arg, "cei"))
}
