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

	var oddLetterSum int
	var letterSums []int
	scanner := bufio.NewScanner(file)
	fmt.Printf("\nLetter Value Sum\n** Bonus Challenge **\n")

	fmt.Printf("\n1. Find the only word with a letter sum of 319\n")
	for scanner.Scan() {
		letterSums = append(letterSums, letterSum(scanner.Text()))
		if letterSum(scanner.Text()) >= 319 {
			fmt.Println(string(scanner.Text()))
		} else if letterSum(scanner.Text())%2 != 0 {
			oddLetterSum++
		}
	}
	fmt.Printf("\n2. Number of words with an odd letter sum\n%v\n", oddLetterSum)

	mostCommonLetterSums := countDuplicates(letterSums)
	count, maxLetterSum := mostCommonDuplicate(mostCommonLetterSums)

	fmt.Printf("\n3. What letter sum is most common, and how many words have it?\n%v, %v\n", maxLetterSum, count)

	//fmt.Printf("\n4. Find the pairs of words with the same letter sum whose lengths differ by 11 letters.\n")
}

func countDuplicates(listArray []int) map[int]int {
	duplicateArray := make(map[int]int)
	for _, i := range listArray {
		_, val := duplicateArray[i]
		if val {
			duplicateArray[i]++
		} else {
			duplicateArray[i] = 1
		}
	}
	return duplicateArray
}

func mostCommonDuplicate(listMap map[int]int) (int, int) {
	var count, item int
	for v, i := range listMap {
		if i > count {
			count = i
			item = v
		}
	}
	return item, count
}
