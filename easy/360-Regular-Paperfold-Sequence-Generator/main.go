package main

import (
	"flag"
	"fmt"
)

const seed = "11"

func main() {
	cycles := flag.Int("c", 1, "number of cycles")
	flag.Parse()
	fmt.Println(dragon(seed, *cycles))
}

func dragon(seed string, num int) string {
	return ""
}
