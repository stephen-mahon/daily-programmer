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
	key, image := keyandMaze(file)
	fmt.Printf("%v\n", key)
	for _, line := range image {
		for _, v := range line {
			fmt.Println(v)
		}
	}
	fmt.Printf("\n")
	maze := createMaze(image, key)
	for _, v := range maze {
		fmt.Println(v)
	}
	fmt.Printf("\nOne Pass\n")
	vals := onePass(maze)

	fmt.Printf("\n\n###\n%v\n", vals)

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

func createMaze(input [][]string, key string) []string {
	var maze []string
	for _, v := range input {
		tmp := ""
		for i := range v {
			line := v[i]
			for j := range line {
				if strings.ContainsAny(key, string(line[j])) {
					tmp += " "

				} else if line[j] != 32 {
					tmp += "1"
				}
			}
		}
		maze = append(maze, tmp)
	}

	return maze
}

func onePass(maze []string) []string {
	nodes := make([]string, len(maze))

	for i := range maze {
		for j := range maze[i] {
			var tmp = ""
			// Enter maze on 0th i-index
			if i != 0 {
				if j != 0 && (maze[i][j-1] == '1' && maze[i-1][j] == '1') {
					fmt.Printf("0")
					tmp += "0"
				} else if (j != len(maze[i])-1 && maze[i][j] != '1') && (maze[i][j+1] == '1' && maze[i-1][j] == '1') {
					fmt.Printf("0")
					tmp += "0"
				} else {
					fmt.Printf("%v", string(maze[i][j]))
					tmp += string(maze[i][j])
				}

				if j == len(maze[i])-1 {
					fmt.Printf("\n")
					nodes[i] = tmp
				}

			} else {

				if maze[i][j] == 32 {
					fmt.Print("0")
					tmp += "0"
				} else {
					fmt.Printf("%v", string(maze[i][j]))
					tmp += string(maze[i][j])
				}

				if j == len(maze[i])-1 {
					fmt.Printf("\n")
					nodes[i] = tmp
				}
			}
		}

		//fmt.Printf("\n")

	}
	return nodes
}
