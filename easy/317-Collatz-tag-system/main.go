package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
	Production Rules
	a -> bc
	b -> a
	c -> aaa

	1. Remove the first two characters of the word
	2. Append accornding to the corresponding production rules of the first letter of the wor.
	3. if the word < 3, stop, otherwise repeat.

	init word: aaa
		remove the first two letters: a
		append based on the first letter of the work: a -> bc: abc
		Less than three? No
	new word: abc
		remove the first two letters: c
		append based on the first letter of the work: a -> bc: cbc
		Less than three? No
	new word: cbc
		remove the first two letters: c
		append based on the first letter of the work: c -> aaa: caaa
		Less than three?
	new word: caaa
		etc.

	repeat until the length of the word is one.

	to do:
		[x] check input for correct alaphabet: an error is thrown if a letter isn't part of the map
		[x] output to text file would be neat.
*/

var production = map[byte]string{
	'a': "bc",
	'b': "a",
	'c': "aaa",
}

func main() {

	init := flag.String("a", "aaa", "challenge input. production letters must contain a, b, and or c")
	fileName := flag.String("f", "output.txt", "file to write")
	flag.Parse()

	f, err := os.Create(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = checkLetters(*init)
	if err != nil {
		log.Fatalf("production letters must contain a, b, and or c: %v", err)
	}
	err = tagSystem(*init, f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Complete. Wrote to file %q\n", *fileName)
}

func checkLetters(letters string) error {
	for _, letter := range letters {
		if !strings.ContainsAny(string(letter), "abc") {
			return fmt.Errorf("unknown production letter - check input arguments:  %q in %q", letter, letters)
		}
	}
	return nil

}

func tagSystem(word string, f *os.File) error {
	dat := []byte(word + "\n")
	_, err := f.Write(dat)
	if err != nil {
		return err
	}
	if len(word) >= 2 {
		newWord := word[2:] + production[word[0]]
		tagSystem(newWord, f)
	}
	return nil
}
