package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var title = "Surround the circles"

type circle struct {
	x float64
	y float64
	r float64
}

type pos struct {
	x float64
	y float64
}

func readfile(fileName string) [][]string {
	file, err := os.Open(fileName + ".txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var output [][]string
	for scanner.Scan() {
		ln := strings.Split(scanner.Text(), ",")
		output = append(output, ln)

	}

	return output
}

func findCircles(args [][]string) []circle {
	var val circle
	var set []circle
	for i := range args {
		val.x, _ = strconv.ParseFloat(args[i][0], 64)
		val.y, _ = strconv.ParseFloat(args[i][1], 64)
		val.r, _ = strconv.ParseFloat(args[i][2], 64)
		set = append(set, val)
	}

	return set
}

func findCorners(c []circle) []pos {
	bottomLeft := pos{x: c[0].x, y: c[0].y}
	var topRight pos
	for i := 1; i < len(c); i++ {
		if (bottomLeft.x > c[i].x) || (bottomLeft.y > c[i].y) {
			bottomLeft.x = c[i].x - c[i].r
			bottomLeft.y = c[i].y - c[i].r
		}
		if (c[i].x > topRight.x+c[i].r) || (c[i].y > topRight.y+c[i].r) {
			topRight.x = c[i].x + c[i].r
			topRight.y = c[i].y + c[i].r
		}
	}
	return []pos{bottomLeft, topRight}
}

func checkBounds(b []pos, c []circle) []pos {
	for i := range c {
		if (c[i].x - c[i].r) < b[0].x {
			b[0].x = c[i].x - c[i].r
		}
		if (c[i].y - c[i].r) < b[0].y {
			b[0].y = c[i].y - c[i].r
		}
		if (c[i].x + c[i].r) > b[1].x {
			b[1].x = c[i].x + c[i].r
		}
		if (c[i].y + c[i].r) > b[1].y {
			b[1].y = c[i].y + c[i].r
		}
	}
	return b
}

func main() {
	fmt.Println(title)
	circles := findCircles(readfile("input"))
	corners := findCorners(circles)
	corners = checkBounds(corners, circles)
	fmt.Println(corners)

}
