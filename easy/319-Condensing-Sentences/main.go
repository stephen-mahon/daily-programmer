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
	fmt.Println(string(dat))
	//use regexp
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
