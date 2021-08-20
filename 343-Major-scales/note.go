package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "-help" {
		fmt.Println("Major scales.")
		fmt.Println("This programs takes the name of a major scale and the solfège name of a note, and returns the corresponding note in that scale.")
		fmt.Println("The movable do solfège system are referred to by the names Do, Re, Mi, Fa, So, La, and Ti, respectively.")
		fmt.Println("The 12 notes in the chromatic scale are: C  C#  D  D#  E  F  F#  G  G#  A  A#  B")
		fmt.Println("Usage: ./note <string1, string2>")
		fmt.Println("Example: ./note D Mi")
		fmt.Println("Output: F#")
	} else {
		if len(args) != 2 {
			fmt.Println("You must enter two arguments! Type -help for help.")
		} else {
			fmt.Println(note(args))
		}
	}
}

func reorder(note string, notes []string) []string {

	for i := range notes {
		if notes[i] == note {
			return append(notes[i:], notes[:i]...)
		}
	}

	fmt.Printf("error! you must use the chromatic scale you entered %q. see /help for more info.\n", note)
	log.Fatal()
	return notes
}

func note(args []string) string {
	note := strings.ToUpper(args[0])
	solfage := strings.ToLower(args[1])

	scale := map[string]int{
		"do": 0,
		"re": 2,
		"mi": 4,
		"fa": 5,
		"so": 7,
		"la": 9,
		"ti": 11,
	}

	notes := []string{
		"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B",
	}

	notes = reorder(note, notes)
	v, ok := scale[solfage]
	if ok != true {
		fmt.Printf("error! you must use the movable do solfège system. you entered %q. see /help for more info.\n", solfage)
		log.Fatal()
	}

	return notes[v]
}
