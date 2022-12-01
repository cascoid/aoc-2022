package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var allGroups = make([][]int64, 0)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := string(b)

	splitInputs := strings.Split(input, "\n")

	var currentGroup []int64
	for i, splitInput := range splitInputs {
		if splitInput == "" || i == len(splitInputs)-1 {
			allGroups = append(allGroups, currentGroup)
			currentGroup = []int64{}
			continue
		}
		conv, err := strconv.ParseInt(splitInput, 10, 64)

		if err != nil {
			panic(err)
		}

		currentGroup = append(currentGroup, conv)
	}

	partOne()

	partTwo()
}

func partOne() {
	var highestElfTotal int64
	for _, group := range allGroups {
		var total int64
		for _, num := range group {
			total += num
		}
		if total > highestElfTotal {
			highestElfTotal = total
		}
	}

	fmt.Println(highestElfTotal)
}

func partTwo() {
	var threeHighestElvesTotal = []int64{0, 0, 0}

	for _, group := range allGroups {
		var total int64
		for _, num := range group {
			total += num
		}

		for i, num := range threeHighestElvesTotal {
			if total > num {
				threeHighestElvesTotal[i] = total
				break
			}
		}
	}

	var total int64
	for _, num := range threeHighestElvesTotal {
		total += num
	}

	fmt.Println(total)
}
