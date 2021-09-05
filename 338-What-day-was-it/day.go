package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "-help" {
		fmt.Println("What day was it again?")
		fmt.Println("Usage: ./day <YEAR> <MONTH> <DATE>")
		fmt.Println("Example: ./ 2021 09 04")
		fmt.Println("Output: Saturday")
	} else {
		if len(args) != 3 {
			fmt.Println("You must enter three arguments! Type -help for help.")
		} else {
			fmt.Println(day(args))
		}

	}
}

func day(args []string) string {
	var date []int
	for i, v := range args {
		val, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("expected an integer but recieved %q\n", v)
			log.Fatal()
		}

		if i == 0 {
			if val > 0 && val < 8000 {
				date = append(date, val)
			} else {
				fmt.Printf("Year must be between 0 and 8000. Recieved %q\n", v)
				log.Fatal()
			}
		} else if i == 1 {
			if val > 0 && val < 13 {
				date = append(date, val)
			} else {
				fmt.Printf("Month must be between 0 and 13. Recieved %q\n", v)
				log.Fatal()
			}
		} else {
			if val > 0 && val < 32 {
				date = append(date, val)
			} else {
				fmt.Printf("Date must be between 0 and 32. Recieved %q\n", v)
				log.Fatal()
			}
		}

	}

	return dayOfTheWeek(date)
}

func dayOfTheWeek(date []int) string {
	if len(date) != 3 {
		fmt.Printf("Date format is incorrect: recieved %q\n", date)
		log.Fatal()
	}
	/*year, month, day := date[0], date[1], date[2]
	var leapYear int
	if year%4 == 0 {
		leapYear++
	}
	*/
	return ""
}
