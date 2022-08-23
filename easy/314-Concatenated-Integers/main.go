package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

// To do
//	[ ] concat loop

func main() {
	args := flag.String("a", "", "a list of integers")
	flag.Parse()
	vals := strings.Split(*args, " ")
	heapsAlgorithm(vals, len(vals))

}

func heapsAlgorithm(s []string, l int) {
	if l == 1 {
		tmp := ""
		for i := range s {
			tmp += s[i]
		}
		val, _ := strconv.Atoi(tmp)
		fmt.Println(val)
	}

	for i := 0; i < l; i++ {
		heapsAlgorithm(s, l-1)
		if l%2 == 1 {
			s[0], s[l-1] = s[l-1], s[0]
		} else {
			s[i], s[l-1] = s[l-1], s[i]
		}
	}
}
