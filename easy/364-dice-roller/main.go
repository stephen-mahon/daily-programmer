package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var s1 = rand.NewSource(time.Now().UnixNano())
var r1 = rand.New(s1)

func main() {

	arg := flag.String("d", "1d2", "dice roll in the format NdM")
	flag.Parse()
	numDie, sides, err := parseDieRoll(*arg)
	if err != nil {
		log.Fatal(err)
	}
	roller(numDie, sides)
}

// is there something more elegant that this?
func roller(numDie, sides int) {

	reader := bufio.NewReader(os.Stdin)
	var quit string
	for quit != `q` {
		rolls := diceThrowTotal(numDie, sides)
		fmt.Println(rolls, sumSlice(rolls))
		fmt.Print("\t\"q\" to exit\t")
		quit, _ = reader.ReadString('\n')
		quit = strings.Trim(quit, "\n")
	}
}

func sumSlice(slice []int) int {
	var sum int
	for _, v := range slice {
		sum += v
	}
	return sum
}

func diceThrowTotal(dice int, sides int) []int {
	var vals []int
	for i := dice; i > 0; i-- {
		val := r1.Intn(sides) + 1
		vals = append(vals, val)
	}
	return vals
}

func parseDieRoll(arg string) (numDie int, sides int, err error) {
	if strings.Contains(arg, "d") {
		dice := strings.Split(arg, "d")
		numDie, err = strconv.Atoi(dice[0])
		if err != nil || numDie < 1 || numDie > 100 {
			return 0, 0, fmt.Errorf("input error - dice number must a number between 1 and 100: %s", dice[0])
		}
		sides, err = strconv.Atoi(dice[1])
		if err != nil || sides < 2 || sides > 100 {
			return 0, 0, fmt.Errorf("input error - the sides on the die must a number between 2 and 100: %s", dice[1])
		}
	} else {
		return 0, 0, fmt.Errorf("input error - dice roll must be of format NdM, %s", arg)
	}
	return numDie, sides, nil
}
