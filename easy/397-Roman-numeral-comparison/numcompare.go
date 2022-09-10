package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {
	input := flag.String("o", "", "Roman Numeral Comparison. Given two roman numerals, return whether the first one is less than the second one.")
	flag.Parse()
	args := strings.Split(*input, " ")
	if len(args) != 2 {
		log.Fatalf("arguments length is not equal to two: len(%q)=%v", args, len(args))
	}
	fmt.Printf("numcompare(%q, %q) => %v\n", args[0], args[1], numeralToInt(args[0]) < numeralToInt(args[1]))

}

// numeral To Int returns the integer val of a input roman numeral. The program panics when a value is not a roman numeral.
func numeralToInt(romanNumeral string) int {
	var val int
	var sub int
	for i, v := range romanNumeral {
		_, isVal := numeralMap(v)
		if isVal != false {
			// loop for running substractive not complete
			if i != 0 && string(romanNumeral[i-1]) == "I" && string(v) == "X" {
				sub += 1
			}
			x, _ := numeralMap(v)
			val += x
		} else {
			fmt.Println("Value", string(v), "in", romanNumeral, "is not a Roman Numeral!")
			panic(string(v))
		}
	}
	return val - sub

}

// This numeral map acts like a switch but the syntax is cleaner and I can double this func as a validation for the input string
// The challenge does say to try and solve without determining the numerical value of the input so this is not the correct solution.
func numeralMap(numeral rune) (int, bool) {
	numeralVal := make(map[string]int)
	numeralVal["M"] = 1000
	numeralVal["D"] = 500
	numeralVal["C"] = 100
	numeralVal["L"] = 50
	numeralVal["X"] = 10
	numeralVal["V"] = 5
	numeralVal["I"] = 1
	v, ok := numeralVal[string(numeral)]
	return v, ok
}
