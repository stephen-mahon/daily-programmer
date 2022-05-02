package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	change := flag.Int("change", 0, "making change")
	flag.Parse()
	if *change >= 0 {
		fmt.Println(makeChange(*change))
	} else {
		fmt.Println("change must be a postive value: ", *change)
		os.Exit(1)
	}

}

func makeChange(arg int) int {
	coins := [6]int{500, 100, 25, 10, 5, 1}
	var change int
	for i := range coins {
		res := arg / coins[i]
		change += res
		arg -= res * coins[i]
	}
	return change
}
