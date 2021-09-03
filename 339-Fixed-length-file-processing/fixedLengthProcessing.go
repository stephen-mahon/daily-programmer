package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "-help" {
		fmt.Println("Fixed-length file processing")
	} else {
		getLinktoFile("https://gist.githubusercontent.com/anonymous/747d5e3bbc57949d8bfe5fd82f359acb/raw/761277a2dcacafb3c06a1e6d0e405ca252098c09/Employee%2520Records.txt", "input.txt")
		highestSalary()
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

func highestSalary() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	salList := make(map[string]int)
	var name string
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "::EXT::") {
			name = strings.TrimRight(line[:19], " ")
		}

		if strings.HasPrefix(line, "::EXT::SAL") {
			salary, _ := strconv.Atoi(line[11:])
			salList[name] = salary
		}
	}
	var maxSal int

	for i := range salList {
		if salList[i] > maxSal {
			maxSal = salList[i]
			name = i
		}
	}
	fmt.Printf("%v $%v", name, moneyComma(maxSal))
}

func moneyComma(amount int) string {
	t := strconv.Itoa(amount)
	var j int
	var money string
	for i := len(t) - 1; i >= 0; i-- {
		if j%3 == 0 && j != 0 {
			money += ","
		}
		j++
		money += string(t[i])
	}
	money = reverse(money)
	return money
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
