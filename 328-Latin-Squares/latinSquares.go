package main

import (
	"flag"
	"fmt"
	"sort"
	"strings"
)

/*
	**bug 1**
	second test fail
		> -length=2 -array="1 3 3 4"
		> [1 3]
		> [3 4]
		>
		> true

	I suspect its due to the rows and cols being identical. The test must be for a list of number in order. Or check non matching rows and columns for instance R0 && C1 == 0.
*/

var title = "Latin squares"

type grid struct {
	x int
	y int
}

type sortRunes []rune

func main() {
	fmt.Println(title)
	length := flag.Int("length", 5, "the length, n, of the array.")
	array := flag.String("array", "1 2 3 4 5 5 1 2 3 4 4 5 1 2 3 3 4 5 1 2 2 3 4 5 1", "the array matrix in the form n x n where n is the length of the array.")
	flag.Parse()

	square := parseArray(length, array)
	printArray(square)

	cols := getCols(square)

	square = reduceArray(square)
	cols = reduceArray(cols)

	sortedArray := joinArray(square)
	sortedCols := joinArray(cols)

	fmt.Println(latinTest(sortedArray, sortedCols))

}

func parseArray(length *int, vals *string) [][]string {
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

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func canonicalize(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func printArray(array [][]string) {
	for _, v := range array {
		fmt.Println(v)
	}
	fmt.Println()
}

func reduceArray(array [][]string) [][]string {

	temp := make([]string, 0)
	var reduced [][]string

	for i := range array {
		tmp := ""
		for _, v := range array[i] {
			tmp += v
		}
		tmp = canonicalize(tmp)
		for j := range tmp {
			temp = append(temp, string(tmp[j]))
		}
		reduced = append(reduced, temp)
		temp = make([]string, 0)
	}

	return reduced
}

func joinArray(array [][]string) []string {
	var vals []string
	for _, v := range array {
		vals = append(vals, strings.Join(v, ""))
	}
	return vals
}

func latinTest(array1, array2 []string) bool {
	for i := range array1 {
		if array1[i] != array2[i] {
			return false
		}
	}
	return true
}
