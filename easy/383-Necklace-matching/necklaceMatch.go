package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "-help" {
		fmt.Println("Same Necklace.")
		fmt.Println("Example: ./sameNecklace nicole icolen")
		fmt.Println("Output: true")
	} else if args[0] == "Bonus" {
		getLinktoFile("https://raw.githubusercontent.com/dolph/dictionary/master/enable1.txt", "enable1.txt")
		bonus()
	} else {
		if len(args) != 2 {
			fmt.Println("You must enter two arguments! Type -help for help.")
		} else {
			word := args[0]
			check := args[1]
			fmt.Println(sameNecklace(word, check))
		}
	}
}

func sameNecklace(word, check string) bool {
	for i := range word {
		if check == word[i:]+word[:i] {
			return true
		}
	}
	return false
}

func repeats(word string) int {
	var num int
	if word == "" {
		return 1
	}

	for i := 0; i < len(word); i++ {
		if word[i:]+word[:i] == word {
			num++
		}
	}
	return num
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

// thanks to user/skeeto
func bonus() {
	file, err := os.Open("enable1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Printf("\nNecklace Match\n** Bonus Challenge **\n")

	list := make(map[string][]string)
	for scanner.Scan() {
		word := scanner.Text()
		ordw := canonicalize(word)
		list[ordw] = append(list[ordw], word)
	}

	for _, words := range list {
		if len(words) == 4 {
			fmt.Printf("%v\n", words)
			break
		}
	}
}

func canonicalize(s string) string {
	c := s + s[:len(s)-1]
	best := s
	for i := 1; i < len(s); i++ {
		if c[i:i+len(s)] < best {
			best = c[i : i+len(s)]
		}
	}
	return best
}
