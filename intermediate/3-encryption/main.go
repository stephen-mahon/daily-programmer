package main

import (
	"fmt"
)

var title = "the quick brown fox jumps over thirteen lazy dogs"

var m = map[string]string{
	"a": "z",
	"b": "y",
	"c": "x",
	"d": "w",
	"e": "v",
	"f": "u",
	"g": "t",
	"h": "s",
	"i": "r",
	"j": "q",
	"k": "p",
	"l": "o",
	"m": "m",
	"n": "m",
	"o": "l",
	"p": "k",
	"q": "j",
	"r": "i",
	"s": "h",
	"t": "g",
	"u": "f",
	"v": "e",
	"w": "d",
	"x": "c",
	"y": "b",
	"z": "a",
	" ": " ",
}

func main() {
	fmt.Println(title)
	text := ciphertext(title)
	fmt.Println(text)
}

func ciphertext(plaintext string) string {
	var text string
	for i := range plaintext {
		text += m[string(plaintext[i])]
	}

	return text
}
