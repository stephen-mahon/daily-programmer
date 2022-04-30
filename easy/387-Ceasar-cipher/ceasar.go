package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "/help" {
		fmt.Println("Ceasar Cipher.")
	} else {
		switch args[0] {
		case "warmup":
			warmup(args[1:])
		}
	}
}

func warmup(args []string) {
	shiftReg, err := strconv.Atoi(args[len(args)-1])
	if err != nil {
		log.Fatal("you must enter a number shift.")
	}
	fmt.Printf("warmup('")

	for i := range args {
		if i != len(args)-2 && i != len(args)-1 {
			fmt.Printf("%v, ", args[i])
		} else if i != len(args)-1 {
			fmt.Printf("%v'", args[i])
		}
	}

	fmt.Printf(", %v) => \"", shiftReg)

	for i, v := range args {
		if i != len(args)-1 {
			for _, u := range v {
				//needs some work here on the truncation
				if shiftReg < 26 {
					fmt.Printf("%v", string(u+rune(shiftReg)))
				} else {
					fmt.Printf("%v", string(u+rune(shiftReg)%26))
				}
			}
			fmt.Printf("\"\n")
		}
	}
}
