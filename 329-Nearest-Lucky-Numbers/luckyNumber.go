package main

import (
	"fmt"
	"os"
)

var title = "Lucky numbers"
var err = "See readme"
var help = ""

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println(err)
	} else if len(args) == 1 && args[0] == "-help" {
		fmt.Println(title)
		fmt.Printf(help)
	} else if len(args) > 1 {
		luckyNums(args)
	} else {
		fmt.Println(err)
	}
}

func luckyNums(args []string) {
	fmt.Println(args)

}
