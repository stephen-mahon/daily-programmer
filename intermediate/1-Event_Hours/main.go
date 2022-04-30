package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
	To Do
		[x] user input time check
		[ ] add edit
*/

var title = "Events organizable by hour"

type event struct {
	hour int
	desc string
}

var schedule = []event{}

func newEvent() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Description: ")
	text, _ := reader.ReadString('\n')
	fmt.Print("Time: ")
	t, _ := reader.ReadString('\n')

	t = strings.TrimSuffix(t, "\r\n")
	time, _ := strconv.Atoi(t)
	_, err := timeCheck(time)

	if err != nil {
		fmt.Println(err)
		newEvent()
	} else {
		userEvent := event{time, text}
		userEvent.Add()
	}
}

func (op *event) Add() {
	schedule = append(schedule, *op)
}

func print(all []event) {
	sort.Slice(all, func(i, j int) bool {
		return all[i].hour < all[j].hour
	})
	fmt.Print("hour, event\n")
	for _, v := range all {
		fmt.Printf("%v, %v\n", v.hour, v.desc)
	}
}

func timeCheck(hour int) (int, error) {
	if hour < 0 || hour > 23 {
		err := fmt.Sprintf("incorrect time: \"%v\". use 24 hour clock.", hour)
		return -1, errors.New(err)
	}
	return hour, nil
}

func trimReadString(in string) string {
	return strings.TrimSuffix(in, "\r\n")
}

func userInterface(exit bool) bool {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("New Event (N) | View Schedule (V) | Quit (Q)")
	input, _ := reader.ReadString('\n')

	switch strings.ToUpper(trimReadString(input)) {
	case "N":
		newEvent()
		return false
	case "V":
		print(schedule)
		return false
	case "Q":
		return true
	default:
		return false
	}
}

func main() {
	fmt.Println(title)
	exit := false

	for !exit {
		exit = userInterface(exit)
		fmt.Println()
	}

}
