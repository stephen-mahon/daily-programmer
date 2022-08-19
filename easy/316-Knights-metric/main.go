package main

import (
	"fmt"
	"math"
)

type Pos struct {
	x int
	y int
}

func main() {
	v1 := newPos(3, 7)
	fmt.Println(v1.mag())

}

func newPos(initX int, initY int) *Pos {
	position := Pos{initX, initY}
	return &position
}

func (pos Pos) mag() float64 {
	return math.Sqrt(float64(pos.x*pos.x) + float64(pos.y*pos.y))
}
