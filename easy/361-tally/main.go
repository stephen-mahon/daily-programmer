package main

import (
	"flag"
	"fmt"
	"strings"
	"unicode"
)

func main() {
	args := flag.String("s", "abcde", "score: lower case +1, upper case -1")
	flag.Parse()
	fmt.Println(parseScore(*args))
}

func parseScore(score string) map[string]int {
	m := make(map[string]int)
	for i := range score {
		_, ok := m[strings.ToLower(string(score[i]))]
		if ok {
			if unicode.IsUpper(rune(score[i])) {
				m[strings.ToLower(string(score[i]))]--
			} else {
				m[string(score[i])]++
			}
		} else {
			m[strings.ToLower(string(score[i]))] = 0
			if unicode.IsUpper(rune(score[i])) {
				m[strings.ToLower(string(score[i]))]--
			} else {
				m[string(score[i])]++
			}
		}
	}
	return m
}
