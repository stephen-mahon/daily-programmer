package main

/*
	Some errors but the test-input works as expected.
	Have an error thrown up with the challenge-input of an empty slice as part of the uniqueVals func.
	Easy to fix
*/

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var message = "Packet Assembler"
var help = "Load file name with -f. See README for more information."
var err = "You must enter arguments! Type -help for help."

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println(err)
	} else if len(args) == 1 && args[0] == "-help" {
		fmt.Println(message)
		fmt.Printf(help)
	} else if len(args) == 2 && args[0] == "-f" {
		dat := readfile(args[1])
		packetAssembler(dat)
	} else {
		fmt.Println(err)
	}
}

func readfile(fileName string) []string {
	file, err := os.Open(fileName + ".txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var output []string

	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	return output
}

func packetAssembler(dat []string) {
	var messageIDs, packetNum, packetLen []int
	var someText []string
	for i := range dat {
		ln := strings.Split(dat[i], "\t")
		messageIDs = appendArray(ln[0], messageIDs)
		packetNum = appendArray(ln[1], packetNum)
		packetLen = appendArray(ln[2], packetLen)

		if len(ln) > 3 {
			someText = append(someText, ln[3])
		} else {
			someText = append(someText, "")
		}
	}

	uniqueIDs, err := uniqueVal(messageIDs)
	if err != nil {
		log.Fatal(err)
	}

	sort.Ints(uniqueIDs)

	for i := range uniqueIDs {
		message := make(map[int]string)
		var messageLen int

		for j := range messageIDs {
			if uniqueIDs[i] == messageIDs[j] {
				message[packetNum[j]] = someText[j]
				messageLen = packetLen[j]
			}

		}

		messageKeys := make([]int, 0)

		for k := range message {
			messageKeys = append(messageKeys, k)
		}

		sort.Ints(messageKeys)

		for j := range messageKeys {
			fmt.Printf("%v\t%v\t%v\t%v\n", uniqueIDs[i], j, messageLen, message[j])
		}
	}
}

func appendArray(text string, vals []int) []int {
	val, err := strconv.Atoi(text)
	if err != nil {
		log.Fatal("error. expected an int but recieved:", text)
	}

	vals = append(vals, val)

	return vals
}

func uniqueVal(vals []int) ([]int, error) {
	if len(vals) == 0 {
		return nil, errors.New("passed empty slice")
	}

	u := make([]int, 0, len(vals))
	m := make(map[int]bool)

	for _, val := range vals {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u, nil
}
