package main

// To do - error check on number fails when input is greater than int

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "-help" {
		fmt.Println("Repeating numbers.")
	} else {
		if len(args) != 1 {
			fmt.Println("You must enter one arguments! Type -help for help.")
		} else {
			repeatingNumbers(args[0])
		}
	}
}

func repeatingNumbers(arg string) {
	_, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Printf("error, expecting a number but recieved %q\n", arg)
		log.Fatal()
	}

	m := make(map[string]int)

	for i := 0; i < len(arg)-1; i++ {
		for j := 0; j < len(arg)-(i+1); j++ {
			m[arg[i:len(arg)-j]]++
		}
	}

	for index, val := range m {
		if val == 1 {
			delete(m, index)
		}
	}
	if len(m) == 0 {
		fmt.Println(0)
	} else {
		for i, v := range m {
			fmt.Printf("%v:%v ", i, v)
		}
		fmt.Println()
	}

}
