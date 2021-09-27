package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

var title = "Lucky numbers"
var err = "See readme"
var help = ""
var count int

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println(err)
	} else if len(args) == 1 && args[0] == "-help" {
		fmt.Println(title)
		fmt.Printf(help)
	} else if len(args) != 1 {
		fmt.Println(err)
	} else {
		val, err := input(args[0])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(isLucky(val))
	}
}

func input(arg string) (int, error) {
	lim, err := strconv.Atoi(arg)
	if err != nil {
		err = errors.New("expecting an number but recieved: \"" + arg + "\"")
	}
	return lim, err
}

func genArray(val int) []int {
	var vals []int
	for i := 1; i <= val; i++ {
		vals = append(vals, i)
	}

	return vals
}

func luckyNums(s []int, count int) ([]int, int) {
	if s[count] >= len(s) {
		return s, count
	}
	var vals []int
	if count == 0 {
		for i := 0; i < len(s); i += 2 {
			vals = append(vals, s[i])
		}
	} else {
		for j := 1; j < len(s)+1; j++ {
			i := j - 1
			if j%s[count] != 0 {
				vals = append(vals, s[i])
			}
		}
	}

	count++
	return luckyNums(vals, count)
}

func isLucky(val int) string {
	vals := genArray(val)
	count := 0
	vals, _ = luckyNums(vals, count)
	luckyNum := vals[len(vals)-1]

	if val == luckyNum {
		return fmt.Sprintf("%v is a lucky number", val)

	} else {
		var nextLuckNum int
		return fmt.Sprintf("%v < %v < %v", luckyNum, val, nextLuckNum)

	}

}
