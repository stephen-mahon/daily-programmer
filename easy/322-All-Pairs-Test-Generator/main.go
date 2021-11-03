package main

import "fmt"

var title = "All pairs test generator"

var testInput = [][]string{{"0", "1"}, {"A", "B", "C"}, {"D", "E", "F", "G"}}

func main() {
	fmt.Println(title)
	printSet(testInput)
	fmt.Printf("\nTotal Combinations: %v", totalCombinations(testInput))
}

func printSet(input [][]string) {
	for i := range input {
		fmt.Printf("%v ", input[i])
	}
}

func totalCombinations(input [][]string) int {
	totalCom := 1
	for i := range input {
		totalCom *= len(input[i])
	}
	return totalCom
}
