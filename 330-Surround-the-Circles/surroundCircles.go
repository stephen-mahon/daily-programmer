package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var title = "Surround the circles"

type circle struct {
	x int
	y int
	t int
}

func readfile(fileName string) {
	file, err := os.Open(fileName + ".txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		ln := strings.Split(scanner.Text(), ",")
		fmt.Println(ln)
	}

	//return output
}

func main() {
	fmt.Println(title)
	readfile("input")
}
