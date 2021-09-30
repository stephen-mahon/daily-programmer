package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var title = "Latin squares"

type grid struct {
	x int
	y int
}

func main() {
	fmt.Println(title)
	length := flag.Int("length", 0, "the length, n, of the array.")
	arrayCheck := flag.String("array", "", "the array matrix in the form n x n where n is the length of the array.")
	flag.Parse()

	arrayPrint(length, arrayCheck)
}

func arrayPrint(length *int, vals *string) {
	arrayStr := strings.Split(*vals, " ")
	var array [][]int
	i := 0
	for j := 0; j < len(arrayStr); j++ {
		if j%*length == 0 && j != 0 {
			fmt.Printf("\n")
			i++

		}
		val, err := strconv.Atoi(arrayStr[j])
		if err != nil {
			log.Fatal("expecting an int but recieved: ", arrayStr[i])
		}
		fmt.Printf("%v ", val)
		array[j] = append(array[j], val)
	}

	fmt.Printf("\n%v, %v\n", array, i)

}
