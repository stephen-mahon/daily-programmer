package main

import "fmt"

func main() {
	vals := []int{}
	fmt.Println(vals, "->", simpleSubsetSum(vals))

}

func simpleSubsetSum(vals []int) bool {
	for i, v := range vals {
		if v == 0 {
			return true
		}
		for j, u := range vals {
			if i != j && v+u == 0 {
				return true
			}
		}
	}
	return false
}
