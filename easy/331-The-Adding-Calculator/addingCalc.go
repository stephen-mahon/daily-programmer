package main

// bug div when len(a) != len(b)

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var message = "The Adding Calulator"
var help = "See README for more information."
var err = "You must enter arguments! Type -help for help."

var overflowInt9 = map[int]int{
	9: 0,
	8: 1,
	7: 2,
	6: 3,
	5: 4,
	4: 5,
	3: 6,
	2: 7,
	1: 8,
	0: 9,
}

var overflowInt10 = map[int]int{
	10: 0,
	9:  1,
	8:  2,
	7:  3,
	6:  4,
	5:  5,
	4:  6,
	3:  7,
	2:  8,
	1:  9,
}

type function struct {
	a  int
	b  int
	op string
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println(err)
	} else if len(args) == 1 && args[0] == "-help" {
		fmt.Println(message)
		fmt.Printf(help)
	} else if len(args) > 1 {
		inputs := parseInput(args)
		fmt.Println(addingCalc(inputs))

	} else {
		fmt.Println(err)
	}
}

func parseInput(args []string) function {

	a, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal("expected a number but recieved:", args[0])
	}
	b, err := strconv.Atoi(args[2])
	if err != nil {
		log.Fatal("expected a number but recieved:", args[2])
	}
	operator := string(args[1])

	return function{a, b, operator}
}

func addingCalc(input function) int {
	switch input.op {
	case "+":
		return add(input.a, input.b)

	case "*":
		return mult(input.a, input.b)

	case "**":
		return expon(input.a, input.b)

	case "-":
		return sub(input.a, input.b)
	case "/":
		return div(input.a, input.b)
	}

	return -1
}

func add(a, b int) int {
	return a + b
}

func mult(a, b int) int {
	var val int
	for i := 0; i < b; i++ {
		val += a
	}
	return val
}

func expon(a, b int) int {
	val := 1
	for i := 0; i < b; i++ {
		val *= a
	}
	return val
}

// https://en.wikipedia.org/wiki/Method_of_complements
// error when the lengths of the divisior does not equal the length of the dividend
// error for representing negative numbers (signed int problem only can't bitshift here)

func sub(a, b int) int {
	s := strconv.Itoa(b)
	var catStr string
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			val, _ := strconv.Atoi(string(s[i]))
			catStr += strconv.Itoa(overflowInt10[val])
		} else {
			val, _ := strconv.Atoi(string(s[i]))
			catStr += strconv.Itoa(overflowInt9[val])
		}
	}
	val, _ := strconv.Atoi(catStr)
	newVal := strconv.Itoa(a + val)
	val, _ = strconv.Atoi(newVal[1:])
	return val
}

func div(a, b int) int {
	val := a
	var count int
	for val > b {
		count++
		val = sub(val, b)

	}
	return count
}
