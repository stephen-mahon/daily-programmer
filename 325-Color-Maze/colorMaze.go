package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

var title = "Color maze"

func main() {

	filename := flag.String("f", "input", "the input text.")
	flag.Parse()
	fmt.Println(title)
	file, err := readfile(*filename)
	if err != nil {
		panic(err)
	}
	key, maze := keyandMaze(file)

	for _, line := range maze {
		for _, v := range line {
			fmt.Println(v)
		}
	}

	fmt.Printf("\n%v\n", key)
	for _, v := range maze {
		for i := range v {
			line := v[i]
			for j := range line {
				if strings.ContainsAny(key, string(line[j])) {
					fmt.Printf("%v", string(line[j]))
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}

}

func readfile(filename string) ([][]string, error) {
	file, err := os.Open(filename + ".txt")
	if err != nil {
		err := errors.New("Cannot open maze text file:" + filename + ".txt\n")
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var output [][]string
	for scanner.Scan() {
		ln := strings.Split(scanner.Text(), "\n")
		output = append(output, ln)

	}

	return output, nil
}

func keyandMaze(input [][]string) (string, [][]string) {

	return strings.ReplaceAll(input[0][0], " ", ""), input[1:]
}
