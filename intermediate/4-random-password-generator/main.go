package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func generate(length int) string {
	rand.Seed(time.Now().UnixNano())
	keys := []rune("QWERTYUIOPASDFGHJKLZXCVBNM" +
		"qwertyuiopasdfghjklzxcvbnm" +
		"1234567890" +
		"`~!@£$%^&*()-=_+[]{};'\\:\"|,./<>?")
	var password strings.Builder
	for i := 0; i < length; i++ {
		password.WriteRune(keys[rand.Intn(len(keys))])
	}

	return password.String()
}

func main() {
	length := flag.Int("l", 8, "the length of the password")
	n := flag.Int("n", 1, "the number of passwords required")
	flag.Parse()

	for i := 0; i < *n; i++ {
		fmt.Println(generate(*length))
	}
}
