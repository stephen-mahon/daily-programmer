package main

import (
	"fmt"
	"os"
	"strconv"
)

const grav float32 = 9.81 // m s^-2

func fSelect(function string, funcSelect []string) (string, bool) {
	var ok bool
	for _, a := range funcSelect {
		if a == function {
			ok = true
		} else {
			ok = false
		}
	}
	return function, ok
}

func physicsCalc(args, funcSelect []string) float32 {
	function := args[0]
	f, err := fSelect(function, funcSelect)
	if err != true {
		fmt.Printf("value %s in %s is not a valid function!\n", f, args)
		panic(f)
	}
	switch f {
	case funcSelect[0]:
		m, err := strconv.ParseFloat(args[1], 32)
		if err != nil {
			panic(err)
		}
		a, err := strconv.ParseFloat(args[2], 32)
		if err != nil {
			panic(err)
		}
		return float32(m * a)
	case funcSelect[1]:
		f, err := strconv.ParseFloat(args[1], 32)
		if err != nil {
			panic(err)
		}
		a, err := strconv.ParseFloat(args[2], 32)
		if err != nil {
			panic(err)
		}
		return float32(f / a)
	case funcSelect[2]:
		f, err := strconv.ParseFloat(args[1], 32)
		if err != nil {
			panic(err)
		}
		m, err := strconv.ParseFloat(args[2], 32)
		if err != nil {
			panic(err)
		}
		return float32(f / m)
	default:
		fmt.Println("Enter a valid function. Use /help for more info.")
		return -1
	}
}

func main() {
	args := os.Args[1:]
	funcSelect := make([]string, 3)
	funcSelect[0] = "force"
	funcSelect[1] = "mass"
	funcSelect[2] = "acceleration"
	fmt.Println(funcSelect)

	if len(args) == 1 && args[0] == "/help" {
		fmt.Println("Force Calculator")
		fmt.Println("Usage: fma force <mass> <acceleration>")
		fmt.Println("Usage: fma mass <force> <acceleration>")
		fmt.Println("Usage: fma acceleration <force> <mass>")
		fmt.Println("Example: ./fma force 72 9.81")
	} else {
		if len(args) != 3 {
			fmt.Println("You must enter three arguments! Type /help for help.")
		} else {
			physicsCalc(args, funcSelect)
		}
	}
}
