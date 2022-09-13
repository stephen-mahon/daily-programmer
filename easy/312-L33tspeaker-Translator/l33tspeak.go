package main

import (
	"flag"
	"fmt"
	"strings"
)

var l33tspeak = map[string]string{
	"A": "4",
	"B": "6",
	"E": "3",
	"I": "1",
	"L": "1",
	"M": "(V)",
	"N": "(\\)",
	"O": "0",
	"S": "5",
	"T": "7",
	"V": "\\/",
	"W": "`//",
}

func main() {
	l33t := flag.String("l33t", "", "l33tspeak")
	flag.Parse()
	var translated string
	for _, v := range *l33t {
		val, ok := l33tspeak[strings.ToUpper(string(v))]
		if ok {
			translated += val
		} else {
			translated += string(v)
		}
	}
	fmt.Println(*l33t, "->", translated)
}
