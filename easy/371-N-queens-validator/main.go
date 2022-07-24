package main

import (
	"flag"
	"fmt"
	"log"
	"math"
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
	vals := qPos(args)
	printBoard(vals)
	fmt.Print("qcheck({")
	for i, v := range vals {
		if i != len(vals)-1 {
			fmt.Printf("%v,", v.y)
		} else {
			fmt.Printf("%v}) => %v\n", v.y, qcheck(vals))
		}
	}
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

func qcheck(vals []chessBoard) bool {
	for i := range vals {
		for j := range vals {
			if i != j && qLogic(vals[i], vals[j]) {
				return false
			}
		}
	}
	return true
}

// if two queens are on the same diagonal they form a square.
func qLogic(i, j chessBoard) bool {
	return i.y == j.y || math.Abs(float64(i.x-j.x)) == math.Abs(float64(i.y-j.y))
}

func qPos(qpos []int) []chessBoard {
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
