package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var funcList = []string{"force", "mass", "acceleration"}

func main() {

	input := flag.String("f", "force", "Find the force, acceleration, or mass.")
	args := flag.String("args", "10 9.81", "input arguments")
	flag.Parse()

	function, err := fSelect(*input, funcList)
	if err != nil {
		log.Fatal(err)
	}

	vals, err := strToSlice(*args)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(vals)
	fmt.Println(physicsCalc(function, vals))

}

func strToSlice(args string) ([]string, error) {
	slice := strings.Split(args, " ")

	if len(slice) != 2 {
		return []string{}, errors.New("You must enter two arguments.")
	}

	return strings.Split(args, " "), nil
}

func fSelect(function string, funcSelect []string) (string, error) {
	for _, a := range funcSelect {
		if a == function {
			return function, nil
		}
	}
	return "", errors.New("function not found.")
}

func physicsCalc(f string, args []string) float32 {
	switch f {
	case funcList[0]:
		m, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			panic(err)
		}
		a, err := strconv.ParseFloat(args[1], 32)
		if err != nil {
			panic(err)
		}
		return float32(m * a)
	case funcList[1]:
		f, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			panic(err)
		}
		a, err := strconv.ParseFloat(args[1], 32)
		if err != nil {
			panic(err)
		}
		return float32(f / a)
	case funcList[2]:
		f, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			panic(err)
		}
		m, err := strconv.ParseFloat(args[1], 32)
		if err != nil {
			panic(err)
		}
		return float32(f / m)
	default:
		fmt.Println("Enter a valid function. Use /help for more info.")
		return -1
	}
}
