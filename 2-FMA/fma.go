package main

import (
	"fmt"
)

func main() {
	fmt.Printf(helloWorld())
}

func helloWorld() string {
	return fmt.Sprintf("Hello World!\n")
}
