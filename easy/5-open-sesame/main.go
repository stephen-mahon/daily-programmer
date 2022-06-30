package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	pass, err := os.ReadFile("password.txt")
	if err != nil {
		log.Panicf("error opening file %s: %d\n", "password.txt", err)
	}

	input := ""

	for string(pass) != input {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("enter password: ")
		input, _ = reader.ReadString('\n')
		input = strings.Trim(input, "\n")
	}
	fmt.Println("welcome to your program")
}
