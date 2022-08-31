package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var webpage = "https://raw.githubusercontent.com/dolph/dictionary/master/enable1.txt"
var enable = "enable1.txt"

func letterSum(s string) int {
	var total int
	for _, v := range s {
		total += int(v - 96)
	}
	return total
}

func main() {
	args := flag.String("o", "", "Letter value sum")
	flag.Parse()
	if strings.ToLower(*args) == "bonus" {
		getLinktoFile(webpage, enable)
		bonus()
	} else {
		fmt.Println(letterSum(*args))
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
	var maxWord string
	var letterSums []int
	var words []string
	scanner := bufio.NewScanner(file)
	fmt.Printf("\nLetter Value Sum\n** Bonus Challenge **\n")

	// Question 1 & 2 use the same loop
	fmt.Printf("\n1. Find the only word with a letter sum of 319\n")
	for scanner.Scan() {
		letterSums = append(letterSums, letterSum(scanner.Text()))
		words = append(words, scanner.Text())
		if letterSum(scanner.Text()) >= 319 {
			maxWord = string(scanner.Text())
		}
		if letterSum(scanner.Text())%2 != 0 {
			oddLetterSum++
		}
	}
	fmt.Printf("%s\n", maxWord)
	fmt.Printf("\n2. Number of words with an odd letter sum\n%v\n", oddLetterSum)

	// Question 3
	fmt.Printf("\n3. What letter sum is most common, and how many words have it?\n")
	mostCommonLetterSums := countDuplicates(letterSums)
	count, maxLetterSum := mostCommonDuplicate(mostCommonLetterSums)
	fmt.Printf("%v, %v\n", maxLetterSum, count)

	// Question 4
	fmt.Printf("\n4. Find the pairs of words with the same letter sum whose lengths differ by 11 letters.\n")

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
			item, count = v, i
		}
	}
	return item, count
}
