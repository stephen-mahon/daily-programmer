package main

import (
	"fmt"
	"sort"
)

var title = "Events organizable by hour"

type event struct {
	hour int
	desc string
}

var events = []event{}

func main() {
	fmt.Println(title)
	var morning event
	morning.hour = 8
	morning.desc = "Coffee"

	var shower event
	shower.hour = 7
	shower.desc = "Shower"

	var run event
	run.hour = 6
	run.desc = "5k run"

	addEvent(morning)
	addEvent(run)
	addEvent(shower)

	printEvents(events)
}

func addEvent(item event) {
	events = append(events, item)
}

func printEvents(all []event) {

	sort.Slice(all, func(i, j int) bool {
		return all[i].hour < all[j].hour
	})

	fmt.Print("hour, event\n")
	for _, v := range all {
		fmt.Printf("%v, %v\n", v.hour, v.desc)
	}
}
