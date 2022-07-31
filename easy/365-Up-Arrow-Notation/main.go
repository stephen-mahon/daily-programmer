package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	args := flag.String("v", "", "up arrow equation")
	flag.Parse()
	ups := strings.Count(*args, "↑")
	fmt.Println(*args, ups)
}
