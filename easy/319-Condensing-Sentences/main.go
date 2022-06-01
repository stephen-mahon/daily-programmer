package main

import (
	"flag"
	"fmt"
	"os"
)

var title = "Condensing Sentences"

type coor struct {
	x int
	y int
}

func main() {
	fmt.Println(title)
	file := flag.String("file", "input.txt", "The input file")
	flag.Parse()

	dat, err := os.ReadFile(*file)
	check(err)
	_ = parseInput(dat)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseInput(dat []byte) []string {
	var out []string
	var word string
	for i := range dat {
		if string(dat[i]) == " " {
			out = append(out, word)
			word = ""
		}
		word += string(dat[i])
	}
	out = append(out, word)
	return out
}
