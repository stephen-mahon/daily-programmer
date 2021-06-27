package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "/help" {
		fmt.Println("Given an hour (0-23) followed by a colon followed by the minute (0-59), returns the time in words using 12-hour format followed by am or pm.")
		fmt.Println("Usage: talkingclock <time>")
		fmt.Println("Example: ./talkingclock 12:05")
	} else {
		if len(args) != 1 {
			fmt.Println("You must enter one argument! Type /help for help.")
		} else {
			time := strings.Split(args[0], ":")
			fmt.Println(talkingClock(time))
		}
	}
}

func talkingClock(time []string) string {
	var meridies, hourTime, minuteTime string
	hour, err := strconv.Atoi(time[0])
	if err != nil {
		fmt.Println(err)
	}
	minute, err := strconv.Atoi(time[1])
	if err != nil {
		fmt.Println(err)
	}
	if hour < 12 {
		meridies = "am"
		hourTime = switchNum(hour)
	} else {
		meridies = "pm"
		hourTime = switchNum(hour - 12)
	}

	if minute > 0 && minute < 10 {
		minuteTime = "oh " + switchNum(minute)
	} else if minute >= 10 && minute < 20 {
		minuteTime = switchNum(minute)
	} else if minute >= 20 && minute < 30 {
		minuteTime = "twenty " + switchNum(minute-20)
	} else if minute >= 30 && minute < 40 {
		minuteTime = "thirty " + switchNum(minute-30)
	} else if minute >= 40 && minute < 50 {
		minuteTime = "forty " + switchNum(minute-40)
	} else if minute >= 50 && minute < 60 {
		minuteTime = "fifty " + switchNum(minute-50)
	} else {
		return "It's " + hourTime + " " + meridies
	}

	talkingTime := "It's " + hourTime + " " + minuteTime + " " + meridies
	return talkingTime
}

func switchNum(num int) string {
	switch num {
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	case 4:
		return "four"
	case 5:
		return "five"
	case 6:
		return "six"
	case 7:
		return "seven"
	case 8:
		return "eight"
	case 9:
		return "nine"
	case 10:
		return "ten"
	case 11:
		return "eleven"
	case 12:
		return "twelve"
	case 13:
		return "thirteen"
	case 14:
		return "fourteen"
	case 15:
		return "fiveteen"
	case 16:
		return "sixteen"
	case 17:
		return "seventeen"
	case 18:
		return "eighteen"
	case 19:
		return "nineteen"
	default:
		return "zero"
	}
}
