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

// watch this https://www.youtube.com/watch?v=rop0W4QDOUI

// make node struct

func main() {

	filename := flag.String("f", "input", "the input text.")
	flag.Parse()
	fmt.Println(title)
	file, err := readfile(*filename)
	if err != nil {
		panic(err)
	}
	maze := createMaze(file)
	printMaze(maze)
	fmt.Println()
	nodes := onePass(maze)
	for _, v := range nodes {
		fmt.Println(v)
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

func createMaze(input [][]string) [][]bool {
	var maze [][]bool

	key := strings.ReplaceAll(input[0][0], " ", "")

	for i := 1; i < len(input); i++ {
		for j := range input[i] {
			letters := strings.ReplaceAll(input[i][j], " ", "")
			var line []bool
			for k := range letters {
				if strings.ContainsAny(key, string(letters[k])) {
					line = append(line, true)
				} else {
					line = append(line, false)
				}

			}
			maze = append(maze, line)
		}
	}
	return maze
}

func onePass(maze [][]bool) [][]int {
	var nodes [][]int

	for i := range maze {
		line := []int{}
		for j := range maze[i] {
			if i != 0 {
				if nodes[i-1][j] == 1 && maze[i][j] {
					line = append(line, 1)
				} else {
					line = append(line, 0)
				}
			} else {
				if maze[i][j] {
					line = append(line, 1)
				} else {
					line = append(line, 0)
				}
			}
		}
		nodes = append(nodes, line)
	}

	return nodes
}

func printMaze(maze [][]bool) {
	for i := range maze {
		for j := range maze[i] {
			if maze[i][j] == true {
				fmt.Print("   ")
			} else {
				fmt.Print("|X|")
			}
			if j == len(maze[i])-1 {
				fmt.Printf("\n")
			}
		}
	}
}
