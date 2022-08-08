package main

import (
	"flag"
	"fmt"
)

func main() {
	input := flag.String("i", "", "A string of alphabetical characters.")
	flag.Parse()
	fmt.Println(recurring(*input))
}

func recurring(arg string) (string, int) {
	letters := make(map[string]int)
	for i := range arg {
		_, ok := letters[string(arg[i])]
		if ok {
			return string(arg[i]), letters[string(arg[i])]
		} else {
			letters[string(arg[i])] = i
		}
	}
	return "", -1
}
