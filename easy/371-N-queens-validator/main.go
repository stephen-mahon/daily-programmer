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
	fmt.Println(qcheck(args))

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
