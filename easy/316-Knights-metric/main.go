package main

import (
	"fmt"
	"math"
)

type pos struct {
	x int
	y int
}

func main() {
	start := newPos(0, 0)
	end := newPos(3, 7)
	fmt.Println(dist(*start, *end))
}

func newPos(initX int, initY int) *pos {
	position := pos{initX, initY}
	return &position
}

func dist(pos1 pos, pos2 pos) (int, int) {
	return int(math.Abs(float64(pos1.x - pos2.x))),
		int(math.Abs(float64(pos1.y - pos2.y)))
}
