package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type chessBoard struct {
	x int
	y int
}

func main() {
	pos := flag.String("s", "", "")
	flag.Parse()
	args, err := inputVerify(pos)
	if err != nil {
		log.Fatalln(err)
	}
	vals := qcheck(args)
	printBoard(vals)

}

func inputVerify(args *string) ([]int, error) {
	input := strings.SplitAfter(*args, " ")
	var out []int
	for _, in := range input {
		val, err := strconv.Atoi(strings.Trim(in, " "))
		if err != nil {
			return []int{}, err
		}
		out = append(out, val)
	}
	return out, nil
}

func qcheck(qpos []int) []chessBoard {
	var pos2D []chessBoard
	for i, v := range qpos {
		pos2D = append(pos2D, chessBoard{x: i, y: v})
	}
	return pos2D
}

// to do: Transpose this table and reverse order
func printBoard(pos []chessBoard) {
	for i, v := range pos {
		fmt.Print(string(97 + v.x))
		for j := range pos {
			if j == pos[i].y {
				fmt.Print(" Q ")
			} else {
				fmt.Print(" . ")
			}
		}
		fmt.Println()
	}

	// this is correct wrt to the "To do"
	fmt.Print(" ")
	for i := range pos {
		fmt.Printf(" %s ", string(97+pos[i].x))
	}
	fmt.Println()
}
