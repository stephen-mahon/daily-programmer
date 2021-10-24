package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
	To Do
		[ ] user input time check
*/

var title = "Events organizable by hour"

type event struct {
	hour int
	desc string
}

var schedule = []event{}

func main() {
	fmt.Println(title)
	var morning event
	morning.hour = 8
	morning.desc = "Coffee"

	var run event
	run.hour = 6
	run.desc = "5k run"

	morning.Add()
	run.Add()

	input()

	print(schedule)

}

func (op *event) Add() {
	schedule = append(schedule, *op)
}

func input() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Description: ")
	text, _ := reader.ReadString('\n')
	fmt.Print("Time: ")
	t, _ := reader.ReadString('\n')
	t = strings.TrimSuffix(t, "\r\n")
	time, _ := strconv.Atoi(t)
	userEvent := event{time, text}
	userEvent.Add()
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
