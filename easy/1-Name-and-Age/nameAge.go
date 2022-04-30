package main

import (
	"flag"
	"fmt"
)

var title = "Name and age"

func main() {
	name := flag.String("name", "Stephen", "Your name.")
	age := flag.Int("age", 1, "Your age.")
	twitter := flag.String("twitter", "stiofanomathuna", "Your Twitter handle")
	flag.Parse()

	fmt.Printf("your name is %s, you are %v years old, and your twitter name is %s.\n", *name, *age, twitterHandle(twitter))
}

func twitterHandle(ID *string) string {
	return "@" + *ID
}
