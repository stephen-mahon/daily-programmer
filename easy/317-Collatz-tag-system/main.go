package main

import (
	"flag"
	"fmt"
)

/*
	a -> bc
	b -> a
	c -> aaa

	start with aaa.
	1. Remove the first two characters
	2. Append accornding to the corresponding production rules.
	3. if the word >= 2, stop, otherwise continue.

	init word: aaa
	remove the first two letters: a
	append bc to a -> abc

	repeat until the length of the word is one.
*/

var production = map[byte]string{
	'a': "bc",
	'b': "a",
	'c': "aaa",
}

func main() {
	init := flag.String("a", "aaa", "challenge input")
	flag.Parse()
	tagSystem(*init)
}

func tagSystem(word string) {
	fmt.Println(word)
	if len(word) >= 2 {
		newWord := word[2:] + production[word[0]]
		tagSystem(newWord)
	}
}
