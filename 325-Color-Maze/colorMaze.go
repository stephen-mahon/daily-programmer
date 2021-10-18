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
	onePass(maze)

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

func createMaze(input [][]string) [][]string {
	var maze [][]string

	key := strings.ReplaceAll(input[0][0], " ", "")

	for i := 1; i < len(input); i++ {
		for j := range input[i] {
			letters := strings.ReplaceAll(input[i][j], " ", "")
			var line []string
			for k := range letters {
				if strings.ContainsAny(key, string(letters[k])) {
					line = append(line, " ")
				} else {
					line = append(line, "1")
				}

			}
			maze = append(maze, line)
		}
	}
	return maze
}

func onePass(maze [][]string) [][]string {
	for i := range maze {
		for j := range maze[i] {
			if i != 0 {
				fmt.Printf("%v ", maze[i][j])
				if j == len(maze[i])-1 {
					fmt.Printf("\n")
				}

			} else {
				// Enter maze on 0th i-index
				if maze[i][j] == " " {
					fmt.Print("0 ")
				} else {
					fmt.Printf("%v ", string(maze[i][j]))
				}

				if j == len(maze[i])-1 {
					fmt.Printf("\n")
				}
			}
		}
	}
	return maze
}

func printMaze(maze [][]string) {
	for i := range maze {
		for j := range maze[i] {
			fmt.Printf("%v ", maze[i][j])
			if j == len(maze[i])-1 {
				fmt.Printf("\n")
			}
		}
	}
}

/*
func onePass(maze []string) []string {

	for i := range maze {
		for j := range maze[i] {

			if i != 0 {
				if j != 0 && (maze[i][j] != '1' && maze[i][j-1] == '1') {
					fmt.Printf("0")
				} else {
					fmt.Printf("%v", string(maze[i][j]))
				}

				if j == len(maze[i])-1 {
					fmt.Printf("\n")
				}

			} else {
				// Enter maze on 0th i-index
				if maze[i][j] == 32 {
					fmt.Print("0")
				} else {
					fmt.Printf("%v", string(maze[i][j]))
				}

				if j == len(maze[i])-1 {
					fmt.Printf("\n")
				}
			}
		}
	}
	return maze
}
*/
