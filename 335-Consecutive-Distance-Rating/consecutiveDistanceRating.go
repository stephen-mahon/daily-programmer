package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var message = "Consecutive Distance Rating"
var help = "Load file name with -f. See README for more information."
var err = "You must enter arguments! Type -help for help."

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println(err)
	} else if len(args) == 1 && args[0] == "-help" {
		fmt.Println(message)
		fmt.Printf(help)
	} else if len(args) == 2 && args[0] == "-f" {
		readfile(args[1])
	} else {
		fmt.Println(consecDistance(args))
	}
}

func inputConvert(args []string) ([]int, error) {
	var vals []int
	for i := range args {
		val, err := strconv.Atoi(args[i])
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}

	return vals, nil
}

func consecDistance(args []string) int {
	vals, err := inputConvert(args)
	if err != nil {
		log.Printf("error in inputs. expecting a number but recieved %q\n", args)
		return -1
	}

	elementMap := make(map[int]int)
	for i := range vals {
		elementMap[vals[i]] = i
	}

	keys := make([]int, 0)
	for k := range elementMap {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	var sum int
	for i, k := range keys {
		if i != 0 && k-1 == keys[i-1] {
			sum += int(math.Abs(float64(elementMap[k] - elementMap[k-1])))
		}
	}

	return sum
}

func findMin(inputs []int) int {
	min := inputs[0]
	for _, v := range inputs {
		if v < min {
			min = v
		}
	}
	return min
}

func readfile(fileName string) {
	file, err := os.Open(fileName + ".txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var vals []string
	for scanner.Scan() {
		vals = append(vals, scanner.Text())
	}
	key := strings.Split(vals[0], " ")
	rows, _ := strconv.Atoi(key[0])
	for i := 1; i <= rows; i++ {
		ln := strings.Split(vals[i], " ")
		fmt.Println(consecDistance(ln))
	}

}
