package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	name := getName()
	age := getAge()

	fmt.Println("Hello,", name, ".")
	fmt.Println("You are", age, "years old.")
}

func getName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What's your name? ")
	name, _ := reader.ReadString('\n')
	return name[:len(name)-1]
}

func getAge() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What year where you born? ")
	year, _ := reader.ReadString('\n')
	year = year[:len(year)-1]
	yearOfBirth, err := strconv.Atoi(year)
	if err != nil {
		panic(err)
	}
	t := time.Now()
	currentYear := t.Year()
	return currentYear - yearOfBirth
}

func getTwitterName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your Twitter name? ")
	twitterName, _ := reader.ReadString('\n')
	return twitterName[:len(twitterName)-1]
}

func getTweet(twitterName string) string {
	twitterLink := "http://www.twitter.com/" + twitterName
	page, err := http.Get(twitterLink)
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(page.Body)
	if err != nil {
		panic(err)
	}
	return string(content)
}
