package main

import (
	"flag"
	"fmt"
	"strings"
)

var title = "Latin squares"

type grid struct {
	x int
	y int
}

func main() {
	fmt.Println(title)
	length := flag.Int("length", 5, "the length, n, of the array.")
	array := flag.String("array", "1 2 3 4 5 5 1 2 3 4 4 5 1 2 3 3 4 5 1 2 2 3 4 5 1", "the array matrix in the form n x n where n is the length of the array.")
	flag.Parse()

	square := arrayPrint(length, array)
	for _, v := range square {
		fmt.Println(v)
	}

	cols := getCols(square)
	fmt.Println(cols)

}

func arrayPrint(length *int, vals *string) [][]string {
	arrayStr := strings.Split(*vals, " ")
	temp := make([]string, 0)
	var array [][]string

	for j := 0; j < len(arrayStr); j++ {
		if (j%*length == 0) && (j != 0) {
			array = append(array, temp)
			temp = make([]string, 0)
		}
		temp = append(temp, arrayStr[j])
	}
	array = append(array, temp)
	return array
}

func getVal(array [][]string, x, y int) string {
	return array[x][y]
}

func getCols(array [][]string) [][]string {
	length := len(array)
	temp := make([]string, 0)
	var cols [][]string
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			temp = append(temp, array[j][i])
		}
		cols = append(cols, temp)
		temp = make([]string, 0)
	}

	return cols
}
