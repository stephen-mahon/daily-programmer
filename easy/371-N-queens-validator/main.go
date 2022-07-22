package main

import (
	"flag"
	"fmt"
)

func main() {
	pos := flag.String("s", "", "")
	flag.Parse()
	fmt.Println(*pos)
}
