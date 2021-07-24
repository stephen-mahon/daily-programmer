package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func letterSum(s string) int {
	var total int
	for _, v := range s {
		total += int(v - 96)
	}
	return total
}

func main() {
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "/help" {
		fmt.Println("Letter value sum.")
	} else if args[0] == "Bonus" {
		getLinktoFile("https://raw.githubusercontent.com/dolph/dictionary/master/enable1.txt", "enable1.txt")
		bonus()
	} else {
		if len(args) != 1 {
			fmt.Println("You must enter one argument! Type /help for help.")
		} else {
			fmt.Println(letterSum(args[0]))
		}
	}
}

func getLinktoFile(webpage, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	page, err := http.Get(webpage)
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(page.Body)
	if err != nil {
		panic(err)
	}
	f.WriteString(string(content))
}

func bonus() {
	file, err := os.Open("enable1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if letterSum(scanner.Text()) >= 319 {
			fmt.Println(string(scanner.Text()))
		}
	}
}
